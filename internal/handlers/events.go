package handlers

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/bamItsCam/nft-tickets/internal/db"
	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
)

type EventPublic struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Venue     string    `json:"venue"`
	Capacity  int32     `json:"capacity"`
}

type EventCreateInput struct {
	Body struct {
		Name     string    `json:"name"`
		Date     time.Time `json:"date"`
		Venue    string    `json:"venue"`
		Capacity int       `json:"capacity" minimum:"1"`
	}
}

type EventOutput struct {
	Body EventPublic
}

type EventGetInput struct {
	Id int `path:"id"`
}

type EventListOutput struct {
	Body []EventPublic
}

type EventListTicketsOutput struct {
	Body []TicketPublic
}

func (a *App) HandleEventCreate(ctx context.Context, input *EventCreateInput) (*EventOutput, error) {
	q := db.New(a.Pool)
	args := db.CreateEventParams{
		Name:     input.Body.Name,
		Date:     input.Body.Date,
		Venue:    input.Body.Venue,
		Capacity: int32(input.Body.Capacity),
	}
	event, err := q.CreateEvent(ctx, args)
	if err != nil {
		slog.Error("failed to create event", "error", err)
		return nil, huma.Error500InternalServerError("failed to create event")
	}
	resp := &EventOutput{
		Body: EventPublic(event),
	}
	return resp, nil
}

func (a *App) HandleEventList(ctx context.Context, input *struct{}) (*EventListOutput, error) {
	q := db.New(a.Pool)
	events, err := q.ListEvents(ctx)
	if err != nil {
		slog.Error("failed to list events", "error", err)
		return nil, huma.Error500InternalServerError("failed to list events")
	}
	resp := &EventListOutput{
		Body: sliceCast(events, func(e db.Event) EventPublic {
			return EventPublic(e)
		}),
	}
	return resp, nil
}

func (a *App) HandleEventGet(ctx context.Context, input *EventGetInput) (*EventOutput, error) {
	q := db.New(a.Pool)
	event, err := q.GetEvent(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("event not found")
	} else if err != nil {
		slog.Error("failed to get event", "error", err)
		return nil, huma.Error500InternalServerError("failed to get event")
	}
	resp := &EventOutput{
		Body: EventPublic(event),
	}
	return resp, nil
}

func (a *App) HandleEventListTickets(ctx context.Context, input *EventGetInput) (*EventListTicketsOutput, error) {
	q := db.New(a.Pool)
	event, err := q.GetEvent(ctx, int32(input.Id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, huma.Error404NotFound("event not found")
	} else if err != nil {
		slog.Error("failed to get event", "error", err)
		return nil, huma.Error500InternalServerError("failed to get event")
	}

	tickets, err := q.ListTicketsByEvent(ctx, event.ID)
	if err != nil {
		slog.Error("failed to list tickets for event", "event_id", event.ID, "error", err)
		return nil, huma.Error500InternalServerError("failed to list tickets for event")
	}
	resp := &EventListTicketsOutput{
		Body: sliceCast(tickets, func(e db.Ticket) TicketPublic {
			return TicketPublic(e)
		}),
	}
	return resp, nil
}
