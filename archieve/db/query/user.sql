-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users (
  username, email
) VALUES (
  ?, ?
);

-- name: UpdateUser :exec
UPDATE users
SET username=?, email=?
WHERE id=?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
