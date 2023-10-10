package profileConfig_usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (p *ProfileConfigUsecaseImpl) Create(ctx context.Context, req *usecase.RequestCreateProfileConfig) (*usecase.ResponseProfileConfig, error) {
	profile, err := p.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, usecase.ProfileNotFound
		}

		return nil, err
	}

	if profile.UserID != req.UserID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, usecase.ProfileUserIDAndReqUserIDNotMatch
	}

	formatConfigValue, err := usecase.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		return nil, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return nil, err
	}

	profileConfig := req.ToModel(FormatConfigValueByte)
	err = p.profileCfgRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		exist, err := p.profileCfgRepo.Create(ctx, profileConfig)
		if err != nil {
			return err
		}
		if exist {
			log.Info().Msg("profile config sudah terdaftar")
			return usecase.ProfileConfigIsExist
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

	resp := &usecase.ResponseProfileConfig{
		ID:          profileConfig.ID,
		ProfileID:   profileConfig.ProfileID,
		ConfigName:  profileConfig.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      profileConfig.Status,
		Days:        days,
		Token:       req.Token,
	}
	return resp, nil
}
