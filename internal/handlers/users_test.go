package handlers_test

import (
	"context"
	"testing"

	dbmocks "github.com/bamItsCam/nft-tickets/internal/db/mocks"

	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/bamItsCam/nft-tickets/internal/handlers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

func TestHandleUserCreate_SelfCustodial(t *testing.T) {
	q := dbmocks.NewMockQuerier(t)
	q.EXPECT().
		CreateUser(mock.Anything, mock.MatchedBy(func(a db.CreateUserParams) bool {
			return a.WalletAddress == stubOwner
		})).
		Return(db.User{ID: 1, Email: "test@example.com", WalletAddress: stubOwner}, nil)

	app := &handlers.App{Querier: q}

	input := &handlers.UserCreateInput{}
	input.Body.Email = "test@example.com"
	input.Body.WalletAddress = stubOwner

	resp, err := app.HandleUserCreate(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body.WalletAddress != stubOwner {
		t.Errorf("got wallet %s, want %s", resp.Body.WalletAddress, stubOwner)
	}
}

func TestHandleUserCreate_CustodialGeneratesWallet(t *testing.T) {
	q := dbmocks.NewMockQuerier(t)
	q.EXPECT().
		CreateUser(mock.Anything, mock.MatchedBy(func(a db.CreateUserParams) bool {
			return common.IsHexAddress(a.WalletAddress)
		})).
		Return(db.User{ID: 1, Email: "test@example.com", WalletAddress: stubOwner}, nil)

	app := &handlers.App{Querier: q}

	input := &handlers.UserCreateInput{}
	input.Body.Email = "test@example.com"

	resp, err := app.HandleUserCreate(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !common.IsHexAddress(resp.Body.WalletAddress) {
		t.Errorf("expected generated wallet to be a valid hex address, got %s", resp.Body.WalletAddress)
	}
}
