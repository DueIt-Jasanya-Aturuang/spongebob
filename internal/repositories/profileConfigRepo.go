package repositories

import (
	"context"
	"database/sql"

	domainerror "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-error"
	domainprofilecfg "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile-cfg"
	"github.com/rs/zerolog/log"
)

type ProfileCfgRepoImpl struct{}

func NewProfileCfgRepoImpl() domainprofilecfg.ProfileCfgRepo {
	return &ProfileCfgRepoImpl{}
}

func (repo *ProfileCfgRepoImpl) scanRow(row *sql.Row) (*domainprofilecfg.ProfileCfg, error) {
	var profileCfg domainprofilecfg.ProfileCfg

	if err := row.Scan(
		&profileCfg.ID,
		&profileCfg.ProfileId,
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
		log.Err(err).Msg(domainerror.LogErrScanning)
		return nil, err
	}
	return &profileCfg, nil
}

func (repo *ProfileCfgRepoImpl) scanRows(rows *sql.Rows) (*[]domainprofilecfg.ProfileCfg, error) {
	var profileCfgs []domainprofilecfg.ProfileCfg

	for rows.Next() {
		var profileCfg domainprofilecfg.ProfileCfg
		if err := rows.Scan(
			&profileCfg.ID,
			&profileCfg.ProfileId,
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
			log.Err(err).Msg(domainerror.LogErrScanning)
			return nil, err
		}
		profileCfgs = append(profileCfgs, profileCfg)
	}
	return &profileCfgs, nil
}

func (repo *ProfileCfgRepoImpl) GetProfileCfgById(ctx context.Context, db *sql.DB, id string) (*domainprofilecfg.ProfileCfg, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
              FROM dueit.m_user_config WHERE profile_id = $1 OR id = $2`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id, id)

	profileCfg, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return profileCfg, nil
}

func (repo *ProfileCfgRepoImpl) GetProfileCfgByScheduler(ctx context.Context, db *sql.DB, model domainprofilecfg.ProfileCfgScheduler) (*[]domainprofilecfg.ProfileCfg, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
              FROM dueit.m_user_config WHERE (config_value->>'config_time_notify')::time >= $1::time AND config_value->'days' ? $2;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrSTMT)
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, model.Time, model.Day)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrQuery)
	}

	profileCfgs, err := repo.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return profileCfgs, nil
}

func (repo *ProfileCfgRepoImpl) StoreProfileCfg(ctx context.Context, tx *sql.Tx, entity domainprofilecfg.ProfileCfg) error {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2)`
	var exists bool
	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrSTMT)
		return err
	}
	if err = querySTMT.QueryRowContext(ctx, entity.ProfileId, entity.ConfigName).Scan(&exists); err != nil {
		log.Err(err).Msg(domainerror.LogErrQuery)
		return err
	}
	if exists {
		return domainerror.ErrProfileConfigAlvailable
	}

	// process insert
	query = `INSERT INTO dueit.m_user_config (id, profile_id, config_name, config_value, status, created_at, created_by, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrSTMT)
		return err
	}

	_, err = execSTMT.ExecContext(
		ctx,
		entity.ID,
		entity.ProfileId,
		entity.ConfigName,
		entity.ConfigValue,
		entity.Status,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.UpdatedAt,
	)
	return err
}

func (repo *ProfileCfgRepoImpl) UpdateProfileCfg(ctx context.Context, tx *sql.Tx, entity domainprofilecfg.ProfileCfg) error {
	query := `UPDATE dueit.m_user_config SET config_value = $1, status = $2, updated_at = $3, updated_by = $4 WHERE id = $5`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domainerror.LogErrSTMT)
		return err
	}

	_, err = stmt.ExecContext(
		ctx,
		entity.ConfigValue,
		entity.Status,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	)
	return err
}
