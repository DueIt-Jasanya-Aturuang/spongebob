package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type SqlTransactionRepoImpl struct {
	db *sql.DB
}

func NewSqlTransactionRepoImpl(db *sql.DB) repository.SqlTransactionRepo {
	return &SqlTransactionRepoImpl{
		db: db,
	}
}

func (s *SqlTransactionRepoImpl) StartTX(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, opts)
	return tx, err
}

func (s *SqlTransactionRepoImpl) Transaction(ctx context.Context, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := s.StartTX(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrStartTx)
		return err
	}

	err = fn(tx)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrTx)
		if errRB := tx.Rollback(); errRB != nil {
			log.Err(err).Msg(exceptions.LogErrTxRollback)
			return fmt.Errorf("error %v : || error rb : %v", err, errRB)
		}
		return err
	}

	return tx.Commit()
}
