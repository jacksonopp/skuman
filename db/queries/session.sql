-- name: GetSessionBySessionId :one
SELECT sessions.session_id, sessions.expires_at, users.id, users.email FROM sessions
INNER JOIN users ON sessions.user_id=users.id
WHERE session_id = $1;

-- name: CreateSession :one
INSERT INTO sessions (
  session_id, user_id, expires_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateSession :one
UPDATE sessions
SET expires_at = $1
WHERE session_id = $1
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE session_id = $1;

-- name: IsSessionExpired :one
SELECT EXISTS (
  SELECT id FROM sessions
  WHERE expires_at < now()
  AND session_id = $1
);

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at < now();