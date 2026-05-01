package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	dbmocks "github.com/bamItsCam/nft-tickets/internal/db/mocks"

	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/bamItsCam/nft-tickets/internal/handlers"
	"github.com/danielgtaylor/huma/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/mock"
)

const stubOwner = "0x0000000000000000000000000000000000000001"

func makeTickets(n int) []db.Ticket {
	tickets := make([]db.Ticket, n)
	for i := range tickets {
		tickets[i] = db.Ticket{ID: int32(i + 1)}
	}
	return tickets
}

func burnedTicket() db.Ticket {
	t := time.Now()
	return db.Ticket{ID: 1, BurnedAt: &t}
}

func httpStatus(err error) int {
	var se huma.StatusError
	if errors.As(err, &se) {
		return se.GetStatus()
	}
	return 0
}

// --- HandleTicketCreate ---

func TestHandleTicketCreate_Capacity(t *testing.T) {
	tests := []struct {
		name          string
		capacity      int32
		existingCount int
	}{
		{"at capacity", 5, 5},
		{"over capacity", 5, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q := dbmocks.NewMockQuerier(t)
			q.EXPECT().GetUser(mock.Anything, int32(1)).Return(db.User{ID: 1, WalletAddress: stubOwner}, nil)
			q.EXPECT().GetEvent(mock.Anything, int32(1)).Return(db.Event{ID: 1, Capacity: tc.capacity}, nil)
			q.EXPECT().ListTicketsByEvent(mock.Anything, int32(1)).Return(makeTickets(tc.existingCount), nil)

			app := &handlers.App{Querier: q}

			input := &handlers.TicketCreateInput{}
			input.Body.EventId = 1
			input.Body.OwnerId = 1

			_, err := app.HandleTicketCreate(context.Background(), input)

			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if got := httpStatus(err); got != http.StatusBadRequest {
				t.Errorf("got status %d, want %d", got, http.StatusBadRequest)
			}
		})
	}
}

// --- HandleTicketBurn ---

func TestHandleTicketBurn_AlreadyBurned(t *testing.T) {
	q := dbmocks.NewMockQuerier(t)
	q.EXPECT().GetTicket(mock.Anything, int32(1)).Return(burnedTicket(), nil)

	app := &handlers.App{Querier: q}
	_, err := app.HandleTicketBurn(context.Background(), &handlers.TicketGetInput{Id: 1})

	if got := httpStatus(err); got != http.StatusConflict {
		t.Errorf("got status %d, want %d", got, http.StatusConflict)
	}
}

// --- HandleTicketTransfer ---

func TestHandleTicketTransfer_AlreadyBurned(t *testing.T) {
	q := dbmocks.NewMockQuerier(t)
	q.EXPECT().GetTicket(mock.Anything, int32(1)).Return(burnedTicket(), nil)

	app := &handlers.App{Querier: q, Opts: &bind.TransactOpts{}}

	input := &handlers.TicketTransferInput{Id: 1}
	input.Body.NewOwnerId = 2
	_, err := app.HandleTicketTransfer(context.Background(), input)

	if got := httpStatus(err); got != http.StatusConflict {
		t.Errorf("got status %d, want %d", got, http.StatusConflict)
	}
}

// --- HandleTicketTransferPrepare ---

func TestHandleTicketTransferPrepare_InvalidAddress(t *testing.T) {
	app := &handlers.App{}
	input := &handlers.TicketTransferPrepareInput{Id: 1, To: "not-an-address"}

	_, err := app.HandleTicketTransferPrepare(context.Background(), input)

	if got := httpStatus(err); got != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", got, http.StatusBadRequest)
	}
}

func TestHandleTicketTransferPrepare_AlreadyBurned(t *testing.T) {
	q := dbmocks.NewMockQuerier(t)
	q.EXPECT().GetTicket(mock.Anything, int32(1)).Return(burnedTicket(), nil)

	app := &handlers.App{Querier: q}
	input := &handlers.TicketTransferPrepareInput{Id: 1, To: stubOwner}

	_, err := app.HandleTicketTransferPrepare(context.Background(), input)

	if got := httpStatus(err); got != http.StatusConflict {
		t.Errorf("got status %d, want %d", got, http.StatusConflict)
	}
}

// --- HandleTicketTransferBroadcast ---

func TestHandleTicketTransferBroadcast_MalformedHex(t *testing.T) {
	app := &handlers.App{}

	input := &handlers.TicketTransferBroadcastInput{Id: 1}
	input.Body.SignedTx = "not-hex"

	_, err := app.HandleTicketTransferBroadcast(context.Background(), input)

	if got := httpStatus(err); got != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", got, http.StatusBadRequest)
	}
}

func TestHandleTicketTransferBroadcast_InvalidTxBinary(t *testing.T) {
	app := &handlers.App{}

	input := &handlers.TicketTransferBroadcastInput{Id: 1}
	input.Body.SignedTx = "0xdeadbeef"

	_, err := app.HandleTicketTransferBroadcast(context.Background(), input)

	if got := httpStatus(err); got != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", got, http.StatusBadRequest)
	}
}
