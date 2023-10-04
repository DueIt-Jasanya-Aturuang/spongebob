package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/helpers"
)

type ProfileConfigUsecaseImpl struct {
	profileRepo    domain.ProfileRepo
	profileCfgRepo domain.ProfileConfigRepo
	notifRepo      domain.NotificationRepo
}

func NewProfileConfigUsecaseImpl(
	profileRepo domain.ProfileRepo,
	profileCfgRepo domain.ProfileConfigRepo,
	notifRepo domain.NotificationRepo,
) domain.ProfileConfigUsecase {
	return &ProfileConfigUsecaseImpl{
		profileRepo:    profileRepo,
		profileCfgRepo: profileCfgRepo,
		notifRepo:      notifRepo,
	}
}

func (p *ProfileConfigUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateProfileConfig) (*domain.ResponseProfileConfig, error) {
	err := p.profileCfgRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer p.profileCfgRepo.CloseConn()

	profile, err := p.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, ProfileNotFound
		}

		return nil, err
	}

	if profile.UserID != req.UserID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, ProfileUserIDAndReqUserIDNotMatch
	}

	formatConfigValue, err := helpers.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		return nil, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return nil, err
	}

	profileConfig := converter.CreateProfileCfgToModel(req, FormatConfigValueByte)
	err = p.profileCfgRepo.StartTx(ctx, helpers.LevelReadCommitted(), func() error {
		exist, err := p.profileCfgRepo.Create(ctx, profileConfig)
		if err != nil {
			return err
		}
		if exist {
			log.Info().Msg("profile config sudah terdaftar")
			return ProfileConfigIsExist
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var days []string
	switch req.ConfigName {
	case "DAILY_NOTIFY":
		days = formatConfigValue["days"].([]string)
	}

	resp := converter.ProfileConfigModelToResponse(profileConfig, req.ConfigValue, days)
	return resp, nil
}

func (p *ProfileConfigUsecaseImpl) GetByNameAndID(ctx context.Context, req *domain.RequestGetProfileConfig) (*domain.ResponseProfileConfig, error) {
	err := p.profileCfgRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer p.profileCfgRepo.CloseConn()

	profile, err := p.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, ProfileNotFound
		}

		return nil, err
	}

	if profile.UserID != req.UserID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, ProfileUserIDAndReqUserIDNotMatch
	}

	profileCfg, err := p.profileCfgRepo.GetByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ProfileConfigNotFound
		}

		return nil, err
	}

	formatConfigValue := map[string]any{}
	err = json.Unmarshal([]byte(profileCfg.ConfigValue), &formatConfigValue)
	if err != nil {
		return nil, err
	}

	var configValue string
	var days []string
	switch req.ConfigName {
	case "DAILY_NOTIFY":
		configValue = fmt.Sprintf("%s %s", formatConfigValue["config_time_user"], formatConfigValue["config_timezone_user"])
		daysInterface := formatConfigValue["days"].([]interface{})
		days = make([]string, len(daysInterface))
		for i, d := range daysInterface {
			days[i] = d.(string)
		}
	case "MONTHLY_PERIOD":
		configValue = fmt.Sprintf("%s", formatConfigValue["config_date"])
	}

	resp := converter.ProfileConfigModelToResponse(profileCfg, configValue, days)
	return resp, nil
}

func (p *ProfileConfigUsecaseImpl) Update(ctx context.Context, req *domain.RequsetUpdateProfileConfig) (*domain.ResponseProfileConfig, error) {
	err := p.profileCfgRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer p.profileCfgRepo.CloseConn()

	profile, err := p.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, ProfileNotFound
		}

		return nil, err
	}

	if profile.UserID != req.UserID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, ProfileUserIDAndReqUserIDNotMatch
	}

	profileCfg, err := p.profileCfgRepo.GetByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ProfileConfigNotFound
		}

		return nil, err
	}

	formatConfigValue, err := helpers.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		return nil, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return nil, err
	}

	profileCfgConv := converter.UpdateProfileCfgToModel(req, FormatConfigValueByte, req.ConfigName, profileCfg.ID)

	err = p.profileCfgRepo.StartTx(ctx, helpers.LevelReadCommitted(), func() error {
		err = p.profileCfgRepo.Update(ctx, profileCfgConv)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var days []string
	switch req.ConfigName {
	case "DAILY_NOTIFY":
		days = formatConfigValue["days"].([]string)
	}

	resp := converter.ProfileConfigModelToResponse(profileCfgConv, req.ConfigValue, days)

	return resp, nil
}

func (p *ProfileConfigUsecaseImpl) GetBySchedulerDailyNotify(ctx context.Context, ProfileConfigScheduler domain.ProfileConfigScheduler) error {
	err := p.profileCfgRepo.OpenConn(ctx)
	if err != nil {
		return err
	}

	profileConfigs, err := p.profileCfgRepo.GetBySchedulerDailyNotify(ctx, ProfileConfigScheduler)
	if err != nil {
		return err
	}

	if len(*profileConfigs) < 1 {
		return nil
	}

	notify, err := p.notifRepo.GetNotifHelperByName(ctx, (*profileConfigs)[0].ConfigName)

	for _, profileConfig := range *profileConfigs {
		formatConfigValue := map[string]any{}
		err = json.Unmarshal([]byte(profileConfig.ConfigValue), &formatConfigValue)
		if err != nil {
			return err
		}

		err = p.profileCfgRepo.StartTx(ctx, helpers.LevelReadCommitted(), func() error {
			err = p.notifRepo.Create(ctx, &domain.Notification{
				ID:           ulid.Make().String(),
				ProfileID:    profileConfig.ProfileID,
				UserConfigID: profileConfig.ID,
				Message:      notify.Message,
				Status:       "unread",
				Title:        notify.Title,
				Icon:         notify.Message,
				AuditInfo: domain.AuditInfo{
					CreatedAt: time.Now().Unix(),
					CreatedBy: profileConfig.ProfileID,
					UpdatedAt: time.Now().Unix(),
				},
			})
			if err != nil {
				return err
			}

			// 	push to fmc

			return nil
		})

		if err != nil {
			return err
		}

	}

	return nil
}
