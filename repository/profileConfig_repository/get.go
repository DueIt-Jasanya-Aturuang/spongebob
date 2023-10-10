package profileConfig_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileConfigRepositoryImpl) GetByNameAndID(ctx context.Context, profileID string, configName string) (*repository.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, 
                     created_by, updated_at, updated_by, deleted_at, deleted_by
			  FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2`

	db, err := p.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profileCfg repository.ProfileConfig
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

func (p *ProfileConfigRepositoryImpl) GetBySchedulerDailyNotify(ctx context.Context, minuteSecond string, day string) (*[]repository.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status
              FROM dueit.m_user_config AS muc
              WHERE (config_value->>'config_time_notify')::time = $1::time 
              AND config_value->'days' ? $2 
              AND status='on' 
              AND config_name='DAILY_NOTIFY'
			  AND NOT EXISTS (
				SELECT 1 FROM dueit.m_notification AS mn WHERE mn.user_config_id = muc.id
				AND DATE_TRUNC('day', to_timestamp(mn.created_at)::date) = current_date
			  )`

	db, err := p.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	rows, err := stmt.QueryContext(ctx, minuteSecond, day)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, errClose)
		}
	}()

	var profileCfgs []repository.ProfileConfig
	var profileCfg repository.ProfileConfig

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

func (p *ProfileConfigRepositoryImpl) GetBySchedulerMonthlyPeriode(ctx context.Context, tgl int, id string) (*[]repository.ProfileConfig, error) {
	query := `SELECT id, profile_id, config_name, config_value, status
              FROM dueit.m_user_config AS muc
              WHERE (config_value->>'config_date') IS NOT NULL
              AND (config_value->>'config_date')::int = $1 
              AND status='on' 
              AND config_name='MONTHLY_PERIOD'
			  AND NOT EXISTS (
				SELECT 1 FROM dueit.m_notification AS mn WHERE mn.user_config_id = muc.id
				AND DATE_TRUNC('day', to_timestamp(mn.created_at)::date) = current_date
			  )`
	if id != "" {
		query += ` AND id > '` + id + `'`
	}
	query += ` ORDER BY id ASC LIMIT 10`

	db, err := p.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
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

	var profileCfgs []repository.ProfileConfig
	var profileCfg repository.ProfileConfig

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
