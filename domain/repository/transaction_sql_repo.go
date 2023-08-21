package repository

import (
	"context"
	"database/sql"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . SqlTransactionRepo
type SqlTransactionRepo interface {
	StartTX(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Transaction(context.Context, *sql.TxOptions, func(*sql.Tx) error) error
}
