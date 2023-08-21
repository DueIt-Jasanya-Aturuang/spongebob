package repository

import (
	"context"
	"database/sql"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . UnitOfWork
type UnitOfWork interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Commit() error
	Rollback() error
}
