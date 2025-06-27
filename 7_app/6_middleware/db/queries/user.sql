-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, 
  password_hash, 
  email,
  created_at,
  updated_at
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;