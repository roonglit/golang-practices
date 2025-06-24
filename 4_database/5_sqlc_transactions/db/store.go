package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Store interface {
	Querier
	CreateUserWithAuditLog(ctx context.Context, arg CreateUserParams) (User, error)
}

type SQLStore struct {
	*Queries
	Conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) Store {
	return &SQLStore{
		Queries: New(conn),
		Conn:    conn,
	}
}

func (s SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.Conn.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
