package db

import (
	"context"
	"database/sql"
	"fmt"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./mocks . SQL
type SQL interface {
	StartTX(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Transaction(context.Context, *sql.TxOptions, func(*sql.Tx) error) error
	SqlDB() *sql.DB
}

type SQLImpl struct {
	*sql.DB
}

func NewSQLImpl() SQL {
	return &SQLImpl{
		NewPgConn(),
	}
}

func (sql *SQLImpl) SqlDB() *sql.DB {
	return sql.DB
}

func (sql *SQLImpl) StartTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := sql.BeginTx(ctx, opts)
	return tx, err
}

func (sql *SQLImpl) Transaction(ctx context.Context, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := sql.StartTX(ctx, opts)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if errRB := tx.Rollback(); errRB != nil {
			return fmt.Errorf("error %v : || error rb : %v", err, errRB)
		}
		return err
	}

	return tx.Commit()
}
