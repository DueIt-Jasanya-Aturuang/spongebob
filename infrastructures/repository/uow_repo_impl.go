package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type UnitOfWorkImpl struct {
	db   *sql.DB
	tx   *sql.Tx
	conn *sql.Conn
}

func NewUnitOfWorkImpl(db *sql.DB) repository.UnitOfWork {
	return &UnitOfWorkImpl{
		db: db,
	}
}

func (repo *UnitOfWorkImpl) OpenConn(ctx context.Context) (*sql.Conn, error) {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		log.Err(err).Msg(exception.LogErrOpenConnDB)
		return nil, err
	}
	return conn, nil
}

func (repo *UnitOfWorkImpl) StartTx(ctx context.Context, opts *sql.TxOptions) error {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		log.Err(err).Msg(exception.LogErrOpenConnDB)
		return err
	}

	tx, err := conn.BeginTx(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exception.LogErrTxStart)
		return err
	}

	repo.tx = tx
	repo.conn = conn
	return nil
}

func (repo *UnitOfWorkImpl) EndTx(err error) error {
	if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, driver.ErrBadConn) {
		log.Info().Msg(exception.LogInfoTxRollback)
		if errRollback := repo.tx.Rollback(); errRollback != nil {
			log.Err(errRollback).Msg(exception.LogErrTxRollback)
			return errRollback
		}
	} else {
		log.Info().Msg(exception.LogInfoTxCommit)
		errCommit := repo.tx.Commit()
		if errCommit != nil && !errors.Is(errCommit, sql.ErrTxDone) {
			log.Err(errCommit).Msg(exception.LogErrTxCommit)
			log.Info().Msg(exception.LogInfoTxRollback)
			if errRollback := repo.tx.Rollback(); errRollback != nil && !errors.Is(err, sql.ErrTxDone) {
				log.Err(errRollback).Msg(exception.LogErrTxRollback)
				return errRollback
			}
		}
	}

	errConn := repo.conn.Close()
	if errConn != nil {
		log.Err(errConn).Msg(exception.LogErrCloseConn)
		// return errConn
	}

	return nil
}

func (repo *UnitOfWorkImpl) GetTx() (*sql.Tx, error) {
	if repo.tx != nil {
		return repo.tx, nil
	}
	return nil, exception.Err500TxNil
}

func (repo *UnitOfWorkImpl) CallTx(tx *sql.Tx) error {
	if tx != nil {
		repo.tx = tx
	}
	return exception.Err500TxNil
}
