package repository

import (
	"context"
	"database/sql"
)

//counterfeiter:generate -o ./../mocks . UnitOfWork
type UnitOfWork interface {
	OpenConn(c context.Context) error
	GetConn() (*sql.Conn, error)
	CloseConn()
	StartTx(c context.Context, opts *sql.TxOptions) error
	EndTx(err error) error
	GetTx() (*sql.Tx, error)
}
