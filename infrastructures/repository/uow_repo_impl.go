package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
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

func (repo *UnitOfWorkImpl) OpenConn(ctx context.Context) error {
	log.Info().Msg("start open con")

	conn, err := repo.db.Conn(ctx)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenConnDB)
		return err
	}

	repo.conn = conn

	return nil
}

func (repo *UnitOfWorkImpl) GetConn() (*sql.Conn, error) {
	log.Info().Msgf("GET CONNECTION")
	if repo.conn != nil {
		return repo.conn, nil
	}

	err := fmt.Errorf("%s", message.ErrConnNilDB)
	log.Err(err).Msg(message.ErrConnNilDB)

	return nil, err
}

func (repo *UnitOfWorkImpl) CloseConn() {
	log.Info().Msg("close connection")
	err := repo.conn.Close()
	if err != nil {
		log.Err(err).Msg(message.ErrCloseConnDB)
	}
}

func (repo *UnitOfWorkImpl) StartTx(ctx context.Context, opts *sql.TxOptions) error {
	log.Info().Msg("START TRANSACTION")

	if repo.conn != nil {
		tx, err := repo.conn.BeginTx(ctx, opts)
		if err != nil {
			log.Err(err).Msg(message.ErrOpenTxDB)
			return err
		}
		repo.tx = tx
		return nil
	}

	err := fmt.Errorf("%s", message.ErrConnNilDB)
	return err
}

func (repo *UnitOfWorkImpl) EndTx(err error) error {
	log.Info().Msg("END TRANSACTION")
	if repo.tx != nil {
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, driver.ErrBadConn) {
			log.Info().Msg(message.InfoTxRollbackDB)
			if errRollback := repo.tx.Rollback(); errRollback != nil {
				log.Err(errRollback).Msg(message.ErrRollbackTxDB)
				return errRollback
			}
		} else {
			log.Info().Msg(message.InfoTxCommitDB)
			if errCommit := repo.tx.Commit(); errCommit != nil && !errors.Is(err, driver.ErrBadConn) {
				log.Err(errCommit).Msg(message.ErrCommitTxDB)
				return errCommit
			}
		}
	}

	err = fmt.Errorf("%s", message.ErrTxNilDB)
	return nil
}

func (repo *UnitOfWorkImpl) GetTx() (*sql.Tx, error) {
	if repo.tx != nil {
		return repo.tx, nil
	}

	err := fmt.Errorf("%s", message.ErrTxNilDB)
	log.Err(err).Msg(message.ErrTxNilDB)
	return nil, err
}
