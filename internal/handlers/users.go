package handlers

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/danielgtaylor/huma/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v5"
)

type UserPublic struct {
	ID            int32     `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	Email         string    `json:"email"`
	WalletAddress string    `json:"walletAddress"`
}

type UserCreateInput struct {
	Body struct {
		Email         string `json:"email"`
		WalletAddress string `json:"walletAddress,omitempty"`
	}
}

type UserCreateOutput struct {
	Body UserPublic
}

type UserGetInput struct {
	Id int `path:"id"`
}

type UserGetByWalletInput struct {
	Wallet string `path:"wallet,omitempty"`
}

type UserOutput struct {
	Body UserPublic
}

type UserListTicketsOutput struct {
	Body []TicketPublic
}

func (a *App) HandleUserCreate(ctx context.Context, input *UserCreateInput) (*UserCreateOutput, error) {
	walletAddress := input.Body.WalletAddress

	// if wallet is provided and valid, proceed with self-custodial
	// otherwise, fall back to generating the wallet here
	if !common.IsHexAddress(walletAddress) {
		userPrivateKey, err := crypto.GenerateKey()
		if err != nil {
			slog.Error("failed to generate user key", "error", err)
			return nil, huma.Error500InternalServerError("failed to create user")
		}
		walletAddress = crypto.PubkeyToAddress(userPrivateKey.PublicKey).Hex()
	}

	args := db.CreateUserParams{
		Email:         input.Body.Email,
		WalletAddress: walletAddress,
	}
	user, err := a.Querier.CreateUser(ctx, args)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return nil, huma.Error500InternalServerError("failed to create user")
	}
	resp := &UserCreateOutput{}
	resp.Body = UserPublic(user)
	return resp, nil
}

func (a *App) HandleUserGet(ctx context.Context, input *UserGetInput) (*UserOutput, error) {
	user, err := a.Querier.GetUser(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("user not found")
	} else if err != nil {
		slog.Error("failed to get user", "error", err)
		return nil, huma.Error500InternalServerError("failed to get user")
	}
	resp := &UserOutput{
		Body: UserPublic(user),
	}
	return resp, nil
}

func (a *App) HandleUserListTickets(ctx context.Context, input *UserGetInput) (*UserListTicketsOutput, error) {
	user, err := a.Querier.GetUser(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("user not found")
	} else if err != nil {
		slog.Error("failed to get user", "error", err)
		return nil, huma.Error500InternalServerError("failed to get user")
	}

	tickets, err := a.Querier.ListTicketsByOwner(ctx, user.ID)
	if err != nil {
		slog.Error("failed to list user's tickets", "user_id", user.ID, "error", err)
		return nil, huma.Error500InternalServerError("failed to list user's tickets")
	}
	resp := &UserListTicketsOutput{
		Body: sliceCast(tickets, func(e db.Ticket) TicketPublic {
			return TicketPublic(e)
		}),
	}
	return resp, nil
}
