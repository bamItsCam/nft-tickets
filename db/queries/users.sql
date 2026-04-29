-- name: CreateUser :one
INSERT INTO users(email, wallet_address)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByWallet :one
SELECT * FROM users
WHERE wallet_address = $1;
