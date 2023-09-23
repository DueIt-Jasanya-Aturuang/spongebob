package _repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type ProfileRepoImpl struct {
	domain.UnitOfWorkRepository
}

func NewProfileRepoImpl(uow domain.UnitOfWorkRepository) domain.ProfileRepo {
	return &ProfileRepoImpl{
		UnitOfWorkRepository: uow,
	}
}

func (p *ProfileRepoImpl) GetByID(ctx context.Context, id string) (*domain.Profile, error) {
	query := `SELECT id, user_id, quotes, profesi, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profile domain.Profile
	if err = stmt.QueryRowContext(ctx, id).Scan(
		&profile.ProfileID,
		&profile.UserID,
		&profile.Quote,
		&profile.Profesi,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &profile, nil
}

func (p *ProfileRepoImpl) GetByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	query := `SELECT id, user_id, quotes, profesi, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profile domain.Profile
	if err = stmt.QueryRowContext(ctx, userID).Scan(
		&profile.ProfileID,
		&profile.UserID,
		&profile.Quote,
		&profile.Profesi,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &profile, nil
}

func (p *ProfileRepoImpl) Create(ctx context.Context, profile *domain.Profile) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL)`
	var exist bool

	tx, err := p.GetTx()
	if err != nil {
		return false, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errQueryStmt)
		}
	}()

	if err = querySTMT.QueryRowContext(ctx, profile.UserID).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	query = `INSERT INTO dueit.m_profiles (id, user_id, quotes, profesi, created_at, created_by, updated_at) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	if _, err = execSTMT.ExecContext(
		ctx,
		profile.ProfileID,
		profile.UserID,
		profile.Quote,
		profile.Profesi,
		profile.CreatedAt,
		profile.CreatedBy,
		profile.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return false, err
	}

	return false, nil
}

func (p *ProfileRepoImpl) Update(ctx context.Context, profile *domain.Profile) error {
	query := `UPDATE dueit.m_profiles SET quotes = $1, profesi = $2,  updated_by = $3, updated_at = $4 
              WHERE user_id = $5 AND id = $6 AND deleted_at IS NULL`

	tx, err := p.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		profile.Quote,
		profile.Profesi,
		profile.ProfileID,
		profile.UpdatedAt,
		profile.UserID,
		profile.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
