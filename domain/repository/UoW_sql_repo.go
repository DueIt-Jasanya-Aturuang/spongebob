package repository

import (
	"context"
	"database/sql"
)

//counterfeiter:generate -o ./../mocks . UnitOfWork
type UnitOfWork interface {
	StartTx(context.Context, *sql.TxOptions) error
	EndTx(error) error
	GetTx() (*sql.Tx, error)
	CallTx(*sql.Tx) error
	OpenConn(context.Context) (*sql.Conn, error)
}
