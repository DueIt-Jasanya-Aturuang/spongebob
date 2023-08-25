package repository

import (
	"context"
	"database/sql"
)

//counterfeiter:generate -o ./../mocks . UnitOfWork
type UnitOfWork interface {
	StartTx(c context.Context, opts *sql.TxOptions) error
	EndTx(err error) error
	GetTx() (*sql.Tx, error)
	CallTx(tx *sql.Tx) error
	OpenConn(c context.Context) (*sql.Conn, error)
}
