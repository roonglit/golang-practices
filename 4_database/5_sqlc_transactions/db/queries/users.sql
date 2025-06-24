-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id ASC;

-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    password = $2,
    email = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;