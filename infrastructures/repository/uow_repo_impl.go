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
		log.Err(err).Msg(exception.LogErrDBOpenConn)
		return nil, err
	}
	return conn, nil
}

func (repo *UnitOfWorkImpl) StartTx(ctx context.Context, opts *sql.TxOptions) error {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBOpenConn)
		return err
	}

	tx, err := conn.BeginTx(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBTxStart)
		return err
	}

	repo.tx = tx
	repo.conn = conn
	return nil
}

func (repo *UnitOfWorkImpl) EndTx(err error) error {
	if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, driver.ErrBadConn) {
		log.Info().Msg(exception.LogInfoDBTxRollback)
		if errRollback := repo.tx.Rollback(); errRollback != nil {
			log.Err(errRollback).Msg(exception.LogErrDBTxRollback)
			return errRollback
		}
	} else {
		log.Info().Msg(exception.LogInfoDBTxCommit)
		errCommit := repo.tx.Commit()
		if errCommit != nil && !errors.Is(errCommit, sql.ErrTxDone) {
			log.Err(errCommit).Msg(exception.LogErrDBTxCommit)
			log.Info().Msg(exception.LogInfoDBTxRollback)
			if errRollback := repo.tx.Rollback(); errRollback != nil && !errors.Is(err, sql.ErrTxDone) {
				log.Err(errRollback).Msg(exception.LogErrDBTxRollback)
				return errRollback
			}
		}
	}

	errConn := repo.conn.Close()
	if errConn != nil {
		log.Err(errConn).Msg(exception.LogErrDBCloseConn)
		// return errConn
	}

	return nil
}

func (repo *UnitOfWorkImpl) GetTx() (*sql.Tx, error) {
	if repo.tx != nil {
		return repo.tx, nil
	}
	log.Err(exception.Err500TxNil).Msg(exception.LogErrDBTxNil)
	return nil, exception.Err500TxNil
}

func (repo *UnitOfWorkImpl) CallTx(tx *sql.Tx) error {
	if tx != nil {
		repo.tx = tx
		return nil
	}
	log.Err(exception.Err500TxNil).Msg(exception.LogErrDBTxNil)
	return exception.Err500TxNil
}
