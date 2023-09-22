-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, verified, verification_code
) VALUES (
    $1, $2, false, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;