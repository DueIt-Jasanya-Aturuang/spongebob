package repository

import (
	"context"
	"database/sql"
)

//counterfeiter:generate -o ./../mocks . UnitOfWork
type UnitOfWork interface {
	BeginTx(context.Context, *sql.TxOptions) error
	Commit() error
	Rollback() error
	GetTx() *sql.Tx
	CallTx(*sql.Tx) error
}
