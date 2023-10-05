package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type ProfileConfigRepoImpl struct {
	domain.UnitOfWorkRepository
}

func NewProfileConfigRepoImpl(uow domain.UnitOfWorkRepository) domain.ProfileConfigRepo {
	return &ProfileConfigRepoImpl{
		UnitOfWorkRepository: uow,
	}
}

func (p *ProfileConfigRepoImpl) Create(ctx context.Context, profileCfg *domain.ProfileConfig) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2)`
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

	if err = querySTMT.QueryRowContext(ctx, profileCfg.ProfileID, profileCfg.ConfigName).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	query = `INSERT INTO dueit.m_user_config 
    					 (id, profile_id, config_name, config_value, status, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

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

	_, err = execSTMT.ExecContext(
		ctx,
		profileCfg.ID,
		profileCfg.ProfileID,
		profileCfg.ConfigName,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.CreatedAt,
		profileCfg.CreatedBy,
		profileCfg.UpdatedAt,
	)

	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
	}

	return false, err
}

func (p *ProfileConfigRepoImpl) Update(ctx context.Context, profileCfg *domain.ProfileConfig) error {
	query := `UPDATE dueit.m_user_config SET config_value = $1, status = $2, updated_at = $3, updated_by = $4
			  WHERE id = $5 and profile_id = $6`

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

	_, err = stmt.ExecContext(
		ctx,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.UpdatedAt,
		profileCfg.UpdatedBy,
		profileCfg.ID,
		profileCfg.ProfileID,
	)

	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
	}

	return err
}

func (p *ProfileConfigRepoImpl) GetByNameAndID(ctx context.Context, profileID string, configName string) (*domain.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, 
                     created_by, updated_at, updated_by, deleted_at, deleted_by
			  FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profileCfg domain.ProfileConfig
	if err = stmt.QueryRowContext(ctx, profileID, configName).Scan(
		&profileCfg.ID,
		&profileCfg.ProfileID,
		&profileCfg.ConfigName,
		&profileCfg.ConfigValue,
		&profileCfg.Status,
		&profileCfg.CreatedAt,
		&profileCfg.CreatedBy,
		&profileCfg.UpdatedAt,
		&profileCfg.UpdatedBy,
		&profileCfg.DeletedAt,
		&profileCfg.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &profileCfg, nil
}

func (p *ProfileConfigRepoImpl) GetBySchedulerDailyNotify(ctx context.Context, ProfileConfigScheduler domain.ProfileConfigScheduler) (*[]domain.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status
              FROM dueit.m_user_config AS muc
              WHERE (config_value->>'config_time_notify')::time = $1::time AND config_value->'days' ? $2 AND status='on' AND config_name='DAILY_NOTIFY'
				AND NOT EXISTS (
					SELECT 1 FROM dueit.m_notification AS mn WHERE mn.user_config_id = muc.id
				)`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	rows, err := stmt.QueryContext(ctx, ProfileConfigScheduler.Time, ProfileConfigScheduler.Day)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, errClose)
		}
	}()

	var profileCfgs []domain.ProfileConfig
	var profileCfg domain.ProfileConfig

	for rows.Next() {
		if err = rows.Scan(
			&profileCfg.ID,
			&profileCfg.ProfileID,
			&profileCfg.ConfigName,
			&profileCfg.ConfigValue,
			&profileCfg.Status,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}
		profileCfgs = append(profileCfgs, profileCfg)
	}

	return &profileCfgs, nil
}

func (p *ProfileConfigRepoImpl) GetBySchedulerMonthlyPeriode(ctx context.Context, tgl int, id string) (*[]domain.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status
              FROM dueit.m_user_config AS muc
              WHERE (config_value->>'config_date') IS NOT NULL
              AND (config_value->>'config_date')::int = $1 
              AND status='on' 
              AND config_name='MONTHLY_PERIOD'
			  AND NOT EXISTS (
					SELECT 1 FROM dueit.m_notification AS mn WHERE mn.user_config_id = muc.id
				)`
	if id != "" {
		query += ` AND id > '` + id + `'`
	}
	query += ` ORDER BY id ASC LIMIT 10`

	conn, err := p.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	rows, err := stmt.QueryContext(ctx, tgl)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, errClose)
		}
	}()

	var profileCfgs []domain.ProfileConfig
	var profileCfg domain.ProfileConfig

	for rows.Next() {
		if err = rows.Scan(
			&profileCfg.ID,
			&profileCfg.ProfileID,
			&profileCfg.ConfigName,
			&profileCfg.ConfigValue,
			&profileCfg.Status,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}
		profileCfgs = append(profileCfgs, profileCfg)
	}

	return &profileCfgs, nil
}
