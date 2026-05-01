-- name: CreateEvent :one
INSERT INTO events(name, date, venue, capacity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1;

-- name: ListEvents :many
SELECT * FROM events;
