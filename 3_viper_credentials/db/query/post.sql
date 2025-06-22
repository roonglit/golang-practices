-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: CreatePost :execresult
INSERT INTO posts (
  title, content, author_id
) VALUES (
  ?, ?, ?
);

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?;

-- name: ListPostsOfUser :many
SELECT posts.*
FROM posts
WHERE posts.author_id = ?