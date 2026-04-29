package handlers

import (
	"net/http"

	"github.com/bamItsCam/nft-tickets/internal/contract"
	"github.com/danielgtaylor/huma/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Instance *bind.BoundContract
	Ticket   *contract.Ticket
	Opts     *bind.TransactOpts
	Pool     *pgxpool.Pool
	BaseUrl  string
}

func RegisterEventRoutes(api huma.API, app *App) {
	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/events",
		Summary:       "Create an event",
		DefaultStatus: http.StatusCreated,
	}, app.HandleEventCreate)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/events",
		Summary: "List all events",
	}, app.HandleEventList)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/events/{id}",
		Summary: "Get an event",
	}, app.HandleEventGet)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/events/{id}/tickets",
		Summary: "List tickets for an event",
	}, app.HandleEventListTickets)
}

func RegisterUserRoutes(api huma.API, app *App) {
	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/users",
		Summary:       "Create a user",
		DefaultStatus: http.StatusCreated,
	}, app.HandleUserCreate)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/users/{id}",
		Summary: "Get a user",
	}, app.HandleUserGet)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/users/{id}/tickets",
		Summary: "List tickets for a user",
	}, app.HandleUserListTickets)
}

func RegisterTicketRoutes(api huma.API, app *App) {
	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/tickets",
		Summary:       "Create a ticket",
		DefaultStatus: http.StatusCreated,
	}, app.HandleTicketCreate)

	huma.Register(api, huma.Operation{
		Method:        http.MethodGet,
		Path:          "/tickets/{id}",
		Summary:       "Get a ticket",
		DefaultStatus: http.StatusOK,
	}, app.HandleTicketGet)

	huma.Register(api, huma.Operation{
		Method:        http.MethodDelete,
		Path:          "/tickets/{id}",
		Summary:       "Redeem a ticket",
		DefaultStatus: http.StatusNoContent,
	}, app.HandleTicketBurn)

	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/tickets/{id}/transfer",
		Summary:       "Transfer a ticket",
		DefaultStatus: http.StatusOK,
	}, app.HandleTicketTransfer)
}

func sliceCast[I, O any](slice []I, f func(I) O) []O {
	result := make([]O, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
