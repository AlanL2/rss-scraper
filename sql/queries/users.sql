-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex')) -- create query that takes arguments as input
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users where api_key = $1;