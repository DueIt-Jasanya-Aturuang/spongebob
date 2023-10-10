package repository

import (
	"context"
	"database/sql"
)

type UnitOfWorkRepository interface {
	StartTx(ctx context.Context, opts *sql.TxOptions, fn func() error) error
	GetTx() (*sql.Tx, error)
	GetDB() (*sql.DB, error)
}

func LevelReadCommitted() *sql.TxOptions {
	return &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}
}
