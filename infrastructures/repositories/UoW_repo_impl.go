package repositories

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type UnitOfWorkImpl struct {
	db *sql.DB
	tx *sql.Tx
}

func NewUnitOfWorkImpl(db *sql.DB) repository.UnitOfWork {
	return &UnitOfWorkImpl{
		db: db,
	}
}

func (u *UnitOfWorkImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := u.db.BeginTx(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrTxStart)
		return nil, err
	}

	u.tx = tx
	return tx, nil
}

func (u *UnitOfWorkImpl) Commit() error {
	err := u.tx.Commit()
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrTxCommit)
		return err
	}

	return nil
}

func (u *UnitOfWorkImpl) Rollback() error {
	err := u.tx.Rollback()
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrTxRollback)
		return err
	}

	return nil
}
