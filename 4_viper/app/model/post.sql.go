// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: post.sql

package model

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :execresult
INSERT INTO posts (
  title, content, author_id
) VALUES (
  ?, ?, ?
)
`

type CreatePostParams struct {
	Title    string
	Content  string
	AuthorID int64
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPost, arg.Title, arg.Content, arg.AuthorID)
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, title, content, author_id, created_at, updated_at FROM posts
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, content, author_id, created_at, updated_at FROM posts
ORDER BY created_at
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsOfUser = `-- name: ListPostsOfUser :many
SELECT posts.id, posts.title, posts.content, posts.author_id, posts.created_at, posts.updated_at
FROM posts
WHERE posts.author_id = ?
`

func (q *Queries) ListPostsOfUser(ctx context.Context, authorID int64) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsOfUser, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
