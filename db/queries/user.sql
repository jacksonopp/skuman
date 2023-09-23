-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

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

-- name: UpdateVerificationCode :exec
UPDATE users
SET verification_code = $1
WHERE id = $2;

-- name: FindUserByIdAndVerification :one
SELECT * FROM users
WHERE id = $1
AND verification_code = $2;

-- name: SetVerification :one
UPDATE users
SET verified = true
WHERE id = $1
AND verification_code = $2
RETURNING *;