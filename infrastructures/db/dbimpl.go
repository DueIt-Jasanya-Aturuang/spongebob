package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBimpl struct {
	*sql.DB
}

func NewDB() *DBimpl {
	return &DBimpl{
		NewPgConn(),
	}
}

func (db *DBimpl) StartTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)
	return tx, err
}

func (db *DBimpl) Transaction(ctx context.Context, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := db.StartTX(ctx, opts)
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
