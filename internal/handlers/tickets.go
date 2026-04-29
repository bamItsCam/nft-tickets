package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/bamItsCam/nft-tickets/internal/contract"
	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/danielgtaylor/huma/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5"
)

var UriPattern = "%s/tickets/%d/metadata"

type TicketCreateInput struct {
	Body struct {
		EventId int32 `json:"eventId"`
		OwnerId int32 `json:"ownerId"`
	}
}

type TicketPublic struct {
	ID        int32      `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	EventID   int32      `json:"eventId"`
	OwnerID   int32      `json:"ownerId"`
	BurnedAt  *time.Time `json:"burnedAt,omitempty"`
}

type TicketCreateOutput struct {
	Body struct {
		TicketPublic
		Hash string `json:"hash"` // todo: do we even need this? or just nice to show?
	}
}

type TicketGetInput struct {
	Id int `path:"id"`
}

type TicketTransferInput struct {
	Id   int `path:"id"`
	Body struct {
		NewOwnerId int32 `json:"newOwnerId"`
	}
}

type TicketOutput struct {
	Body struct {
		TicketPublic
		OnChainOwner string `json:"onChainOwner"`
	}
}

type TicketMetadataOutput struct {
	Body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
}

func (a *App) HandleTicketTransfer(ctx context.Context, input *TicketTransferInput) (*TicketOutput, error) {
	q := db.New(a.Pool)
	tokenId := big.NewInt(int64(input.Id))

	ticket, err := q.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	} else if ticket.BurnedAt != nil {
		slog.Info("ticket already burned, don't bother checking the chain")
		return nil, huma.Error409Conflict("ticket already redeemed")
	}

	newOwner, err := q.GetUser(ctx, int32(input.Body.NewOwnerId))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("new owner not found")
	} else if err != nil {
		slog.Error("failed to get new ticket owner", "error", err)
		return nil, huma.Error500InternalServerError("failed to get new ticket owner")
	}

	// read current ticket ownership from chain, db could be out of sync
	packed := a.Ticket.PackOwnerOf(tokenId)
	currentOwnerAddress, err := bind.Call(a.Instance, &bind.CallOpts{Context: ctx}, packed, a.Ticket.UnpackOwnerOf)
	if err != nil {
		raw, isRevert := ethclient.RevertErrorData(err)
		if isRevert {
			unpacked, _ := a.Ticket.UnpackError(raw)
			if _, ok := unpacked.(*contract.TicketERC721NonexistentToken); ok {
				slog.Warn("failed to transfer on chain, either not found or already redeemed", "error", err)
				return nil, huma.Error404NotFound("ticket not found or already redeemed")
			}
		}
		slog.Error("failed to get current owner of token from chain", "error", err)
		return nil, huma.Error500InternalServerError("failed to transfer ticket")
	}

	// prevent unnecessary chain ops where we can
	if currentOwnerAddress.String() == newOwner.WalletAddress {
		slog.Info("attempted ticket transfer to same owner, breaking out")
		return nil, huma.Error400BadRequest("ticket cannot be transferred to existing owner")
	}

	txOpts := *a.Opts
	txOpts.Context = ctx

	packed = a.Ticket.PackOperatorTransfer(currentOwnerAddress, common.HexToAddress(newOwner.WalletAddress), tokenId)
	_, err = bind.Transact(a.Instance, &txOpts, packed)
	if err != nil {
		slog.Error("failed to transfer ticket on chain", "error", err)
		return nil, huma.Error500InternalServerError("failed to transfer ticket")
	}

	args := db.UpdateTicketOwnerParams{
		ID:      ticket.ID,
		OwnerID: newOwner.ID,
	}

	ticket, err = q.UpdateTicketOwner(ctx, args)
	if err != nil {
		slog.Error("failed to update ticket owner in db", "ticket_id", ticket.ID, "owner_id", args.OwnerID, "error", err)
		return nil, huma.Error500InternalServerError("failed to transfer ticket")
	}

	resp := &TicketOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.OnChainOwner = currentOwnerAddress.Hex()
	return resp, nil
}

func (a *App) HandleTicketBurn(ctx context.Context, input *TicketGetInput) (*struct{}, error) {
	q := db.New(a.Pool)
	tokenId := big.NewInt(int64(input.Id))

	ticket, err := q.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	} else if ticket.BurnedAt != nil {
		slog.Info("ticket already burned, don't bother checking the chain")
		return nil, huma.Error409Conflict("ticket already redeemed")
	}

	txOpts := *a.Opts
	txOpts.Context = ctx

	packed := a.Ticket.PackOperatorBurn(tokenId)
	_, err = bind.Transact(a.Instance, &txOpts, packed)
	if err != nil {
		raw, isRevert := ethclient.RevertErrorData(err)
		if isRevert {
			unpacked, _ := a.Ticket.UnpackError(raw)
			if _, ok := unpacked.(*contract.TicketERC721NonexistentToken); ok {
				slog.Warn("failed to burn ticket on chain, either not found or already redeemed", "error", err)
				return nil, huma.Error404NotFound("ticket not found or already redeemed")
			}
		}
		slog.Error("failed to burn ticket on chain", "error", err)
		return nil, huma.Error500InternalServerError("failed to redeem ticket")
	}

	// chain is source of truth for burning, so update the db after
	_, err = q.BurnTicket(ctx, ticket.ID)
	if err != nil {
		// clean these up somehow? bleh
		// worst case, someone tries to punch this ticket again, and our 409 check above fails, but the chain
		// catches with TicketERC721NonexistentToken, so this isn't _that_ bad
		slog.Error("failed to mark ticket as burned in db", "ticket_id", ticket.ID, "error", err)
		return nil, huma.Error500InternalServerError("failed to redeem")
	}

	return nil, nil
}

func (a *App) HandleTicketGet(ctx context.Context, input *TicketGetInput) (*TicketOutput, error) {
	q := db.New(a.Pool)
	tokenId := big.NewInt(int64(input.Id))

	ticket, err := q.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	}

	packed := a.Ticket.PackOwnerOf(tokenId)
	currentOwnerAddress, err := bind.Call(a.Instance, &bind.CallOpts{Context: ctx}, packed, a.Ticket.UnpackOwnerOf)

	resp := &TicketOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.OnChainOwner = currentOwnerAddress.Hex()
	return resp, nil
}

func (a *App) HandleTicketMetadataGet(ctx context.Context, input *TicketGetInput) (*TicketMetadataOutput, error) {
	q := db.New(a.Pool)

	ticket, err := q.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket metadata")
	}

	event, err := q.GetEvent(ctx, ticket.EventID)
	if errors.Is(err, pgx.ErrNoRows) {
		slog.Error("ticket event not found", "error", err, "ticket_id", ticket.ID, "event_id", ticket.EventID)
		return nil, huma.Error500InternalServerError("failed to get ticket metadata")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket metadata")
	}

	resp := &TicketMetadataOutput{}
	resp.Body.Name = fmt.Sprintf("Ticket #%d", ticket.ID)
	resp.Body.Description = fmt.Sprintf("Ticket for '%s' at '%s'", event.Name, event.Venue)

	return resp, nil
}

func (a *App) HandleTicketCreate(ctx context.Context, input *TicketCreateInput) (*TicketCreateOutput, error) {
	q := db.New(a.Pool)

	args := db.CreateTicketParams{
		EventID: input.Body.EventId,
		OwnerID: input.Body.OwnerId,
	}

	owner, err := q.GetUser(ctx, args.OwnerID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("provided user not found")
	} else if err != nil {
		slog.Error("failed to fetch user", "error", err)
		return nil, huma.Error500InternalServerError("failed to create ticket")
	}

	event, err := q.GetEvent(ctx, args.EventID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("provided event not found")
	} else if err != nil {
		slog.Error("failed to fetch event", "error", err)
		return nil, huma.Error500InternalServerError("failed to create ticket")
	}

	eventTickets, err := q.ListTicketsByEvent(ctx, args.EventID)
	if err != nil {
		slog.Error("failed to fetch event tickets", "error", err)
		return nil, huma.Error500InternalServerError("failed to create ticket")
	}

	if len(eventTickets) >= int(event.Capacity) {
		slog.Warn("event is at capacity", "capacity", event.Capacity)
		return nil, huma.Error400BadRequest("event is at capacity, no room remaining")
	}

	// make the ticket in db first so that we have a token id
	ticket, err := q.CreateTicket(ctx, args)
	if err != nil {
		slog.Error("failed to create ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to create ticket")
	}

	txOpts := *a.Opts
	txOpts.Context = ctx

	packed := a.Ticket.PackMint(common.HexToAddress(owner.WalletAddress), big.NewInt(int64(ticket.ID)), fmt.Sprintf(UriPattern, a.BaseUrl, ticket.ID))
	tx, err := bind.Transact(a.Instance, &txOpts, packed)
	if err != nil {
		slog.Error("failed to mint ticket", "error", err)
		// roll back ticket creation if minting failed
		err := q.DeleteTicket(ctx, ticket.ID)
		if err != nil {
			// worst case, our db thinks a ticket exists that doesn't on chain.
			// this isn't great. this will mess with capacity checks and could give
			// users the false sense that a ticket exists when it does not
			// todo: should we repair this on ticket fetch?
			slog.Error("failed to rollback ticket creation, db and chain are now out of sync", "error", err, "ghost_ticket", ticket.ID)
		}
		return nil, huma.Error500InternalServerError("failed to create ticket")
	}

	resp := &TicketCreateOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.Hash = tx.Hash().String()
	return resp, nil
}
