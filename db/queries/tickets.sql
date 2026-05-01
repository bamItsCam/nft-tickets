-- name: CreateTicket :one
INSERT INTO tickets(event_id, owner_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetTicket :one
SELECT * FROM tickets
WHERE id = $1;

-- name: ListTicketsByEvent :many
SELECT * from tickets
WHERE event_id = $1;

-- name: ListTicketsByOwner :many
SELECT * from tickets
WHERE owner_id = $1;

-- name: BurnTicket :one
UPDATE tickets SET burned_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;

-- name: UpdateTicketOwner :one
UPDATE tickets set owner_id = $2
WHERE id = $1
RETURNING *;
