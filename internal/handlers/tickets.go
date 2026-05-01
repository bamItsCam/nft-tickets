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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5"
)

const UriPattern = "%s/tickets/%d/metadata"

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
		Hash string `json:"hash"`
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

type TicketTransferPrepareInput struct {
	Id int    `path:"id"`
	To string `query:"to"`
}

type TicketTransferBroadcastInput struct {
	Id   int `path:"id"`
	Body struct {
		SignedTx string `json:"signedTx"`
	}
}

type TicketTransferPrepareOutput struct {
	Body struct {
		From     string `json:"from"`
		To       string `json:"to"` // the contract, not the ticket recipient
		Data     string `json:"data"`
		Nonce    uint64 `json:"nonce"`
		GasLimit string `json:"gasLimit"`
		GasPrice string `json:"gasPrice"`
		ChainId  int    `json:"chainId"`
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

func (a *App) HandleTicketTransferPrepare(ctx context.Context, input *TicketTransferPrepareInput) (*TicketTransferPrepareOutput, error) {
	tokenId := big.NewInt(int64(input.Id))
	catchallErr := huma.Error500InternalServerError("failed to generate transfer payload")

	to := input.To
	if !common.IsHexAddress(to) {
		slog.Warn("invalid to address", "to_address", to)
		return nil, huma.Error400BadRequest("invalid to address")
	}

	ticket, err := a.Querier.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	} else if ticket.BurnedAt != nil {
		slog.Info("ticket already burned according to db")
		return nil, huma.Error409Conflict("ticket already redeemed")
	}

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
		return nil, catchallErr
	}

	toAddress := common.HexToAddress(to)

	// prevent unnecessary chain ops where we can
	if currentOwnerAddress == toAddress {
		slog.Info("attempted ticket transfer to same owner")
		return nil, huma.Error400BadRequest("ticket cannot be transferred to existing owner")
	}

	nonce, err := a.Client.PendingNonceAt(ctx, currentOwnerAddress)
	if err != nil {
		slog.Error("failed to get nonceAt", "error", err)
		return nil, catchallErr
	}
	suggestGas, err := a.Client.SuggestGasPrice(ctx)
	if err != nil {
		slog.Error("failed to get suggested gas price", "error", err)
		return nil, catchallErr
	}

	contractAddress := a.Instance.Address()

	packed = a.Ticket.PackTransferFrom(currentOwnerAddress, toAddress, tokenId)
	estimateGas, err := a.Client.EstimateGas(ctx, ethereum.CallMsg{
		From: currentOwnerAddress,
		To:   &contractAddress,
		Data: packed,
	})
	if err != nil {
		slog.Error("failed to get gas limit", "error", err)
		return nil, catchallErr
	}

	resp := &TicketTransferPrepareOutput{}
	resp.Body.To = contractAddress.Hex()
	resp.Body.From = currentOwnerAddress.Hex()
	resp.Body.Data = hexutil.Encode(packed)
	resp.Body.Nonce = nonce
	resp.Body.GasLimit = fmt.Sprintf("0x%x", estimateGas)
	resp.Body.GasPrice = fmt.Sprintf("0x%x", suggestGas)
	resp.Body.ChainId = int(a.ChainId.Int64())

	return resp, nil
}

func (a *App) HandleTicketTransferBroadcast(ctx context.Context, input *TicketTransferBroadcastInput) (*TicketOutput, error) {
	tx := new(types.Transaction)
	if b, err := hexutil.Decode((input.Body.SignedTx)); err != nil {
		slog.Warn("failed to decode signedTx from hex to bytes", "error", err)
		return nil, huma.Error400BadRequest("failed to decode signedTx")
	} else if err := tx.UnmarshalBinary(b); err != nil {
		slog.Warn("failed to decode signedTx from bytes to tx", "error", err)
		return nil, huma.Error400BadRequest("failed to decode signedTx")
	}

	err := a.Client.SendTransaction(ctx, tx)
	if err != nil {
		slog.Error("failed to send client-signed transaction", "error", err)
		return nil, huma.Error500InternalServerError("failed to send transaction")
	}

	r, err := bind.WaitMined(ctx, a.Client, tx.Hash())
	if err != nil {
		slog.Error("waitMined failed", "error", err)
		return nil, huma.Error500InternalServerError("failed to send transaction")
	}

	var newOwnerAddress common.Address
	var tokenId *big.Int

	for _, log := range r.Logs {
		transfer, err := a.Ticket.UnpackTransferEvent(log)
		if err != nil {
			continue
		}
		newOwnerAddress = transfer.To
		tokenId = transfer.TokenId
		break
	}

	if tokenId == nil {
		slog.Error("no transfer event found in logs")
		return nil, huma.Error500InternalServerError("internal error")
	}

	newUser, err := a.Querier.GetUserByWallet(ctx, newOwnerAddress.Hex())
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("new owner wallet not registered to a user")
	} else if err != nil {
		slog.Error("failed to find new user by wallet address", "address", newOwnerAddress, "error", err)
		return nil, huma.Error500InternalServerError("internal error")
	}

	args := db.UpdateTicketOwnerParams{
		ID:      int32(tokenId.Int64()),
		OwnerID: newUser.ID,
	}

	ticket, err := a.Querier.UpdateTicketOwner(ctx, args)
	if err != nil {
		slog.Error("failed to update ticket in db with new owner information", "ticket_id", ticket.ID, "error", err)
		return nil, huma.Error500InternalServerError("internal error")
	}

	resp := &TicketOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.OnChainOwner = newOwnerAddress.Hex()
	return resp, nil
}

func (a *App) HandleTicketTransfer(ctx context.Context, input *TicketTransferInput) (*TicketOutput, error) {
	tokenId := big.NewInt(int64(input.Id))
	txOpts := *a.Opts
	txOpts.Context = ctx

	ticket, err := a.Querier.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	} else if ticket.BurnedAt != nil {
		slog.Info("ticket already burned according to db")
		return nil, huma.Error409Conflict("ticket already redeemed")
	}

	newOwner, err := a.Querier.GetUser(ctx, int32(input.Body.NewOwnerId))
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

	newOwnerAddress := common.HexToAddress(newOwner.WalletAddress)

	// prevent unnecessary chain ops where we can
	if currentOwnerAddress == newOwnerAddress {
		slog.Info("attempted ticket transfer to same owner")
		return nil, huma.Error400BadRequest("ticket cannot be transferred to existing owner")
	}

	packed = a.Ticket.PackOperatorTransfer(currentOwnerAddress, newOwnerAddress, tokenId)
	_, err = bind.Transact(a.Instance, &txOpts, packed)
	if err != nil {
		slog.Error("failed to transfer ticket on chain", "error", err)
		return nil, huma.Error500InternalServerError("failed to transfer ticket")
	}

	args := db.UpdateTicketOwnerParams{
		ID:      ticket.ID,
		OwnerID: newOwner.ID,
	}

	ticket, err = a.Querier.UpdateTicketOwner(ctx, args)
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
	tokenId := big.NewInt(int64(input.Id))

	ticket, err := a.Querier.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	} else if ticket.BurnedAt != nil {
		slog.Info("ticket already burned according to db")
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
	_, err = a.Querier.BurnTicket(ctx, ticket.ID)
	if err != nil {
		// clean these up somehow?
		// worst case, someone tries to punch this ticket again, and our 409 check above fails, but the chain
		// catches with TicketERC721NonexistentToken, so this isn't _that_ bad
		slog.Error("failed to mark ticket as burned in db", "ticket_id", ticket.ID, "error", err)
		return nil, huma.Error500InternalServerError("failed to redeem ticket")
	}

	return nil, nil
}

func (a *App) HandleTicketGet(ctx context.Context, input *TicketGetInput) (*TicketOutput, error) {
	tokenId := big.NewInt(int64(input.Id))

	ticket, err := a.Querier.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	}

	packed := a.Ticket.PackOwnerOf(tokenId)
	currentOwnerAddress, err := bind.Call(a.Instance, &bind.CallOpts{Context: ctx}, packed, a.Ticket.UnpackOwnerOf)
	if err != nil {
		slog.Error("failed to determine owner of token", "error", err, "token_id", tokenId)
		return nil, huma.Error500InternalServerError("failed to get ticket")
	}

	resp := &TicketOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.OnChainOwner = currentOwnerAddress.Hex()
	return resp, nil
}

func (a *App) HandleTicketMetadataGet(ctx context.Context, input *TicketGetInput) (*TicketMetadataOutput, error) {
	ticket, err := a.Querier.GetTicket(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("ticket not found")
	} else if err != nil {
		slog.Error("failed to get ticket", "error", err)
		return nil, huma.Error500InternalServerError("failed to get ticket metadata")
	}

	event, err := a.Querier.GetEvent(ctx, ticket.EventID)
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
	catchallErr := huma.Error500InternalServerError("failed to create ticket")

	args := db.CreateTicketParams{
		EventID: input.Body.EventId,
		OwnerID: input.Body.OwnerId,
	}

	owner, err := a.Querier.GetUser(ctx, args.OwnerID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("provided user not found")
	} else if err != nil {
		slog.Error("failed to fetch user", "error", err)
		return nil, catchallErr
	}

	event, err := a.Querier.GetEvent(ctx, args.EventID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error400BadRequest("provided event not found")
	} else if err != nil {
		slog.Error("failed to fetch event", "error", err)
		return nil, catchallErr
	}

	eventTickets, err := a.Querier.ListTicketsByEvent(ctx, args.EventID)
	if err != nil {
		slog.Error("failed to fetch event tickets", "error", err)
		return nil, catchallErr
	}

	if len(eventTickets) >= int(event.Capacity) {
		slog.Warn("event is at capacity", "capacity", event.Capacity)
		return nil, huma.Error400BadRequest("event is at capacity, no room remaining")
	}

	// make the ticket in db first so that we have a token id
	ticket, err := a.Querier.CreateTicket(ctx, args)
	if err != nil {
		slog.Error("failed to create ticket", "error", err)
		return nil, catchallErr
	}

	txOpts := *a.Opts
	txOpts.Context = ctx

	packed := a.Ticket.PackMint(common.HexToAddress(owner.WalletAddress), big.NewInt(int64(ticket.ID)), fmt.Sprintf(UriPattern, a.BaseUrl, ticket.ID))
	tx, err := bind.Transact(a.Instance, &txOpts, packed)
	if err != nil {
		slog.Error("failed to mint ticket", "error", err)
		// roll back ticket creation if minting failed
		err := a.Querier.DeleteTicket(ctx, ticket.ID)
		if err != nil {
			// worst case, our db thinks a ticket exists that doesn't on chain.
			// this isn't great. this will mess with capacity checks and could give
			// users the false sense that a ticket exists when it does not
			slog.Error("failed to rollback ticket creation, db and chain are now out of sync", "error", err, "ghost_ticket", ticket.ID)
		}
		return nil, catchallErr
	}

	resp := &TicketCreateOutput{}
	resp.Body.TicketPublic = TicketPublic(ticket)
	resp.Body.Hash = tx.Hash().String()
	return resp, nil
}
