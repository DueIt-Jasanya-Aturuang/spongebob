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

func (p *ProfileConfigUsecaseImpl) Update(ctx context.Context, req *usecase.RequsetUpdateProfileConfig) (*usecase.ResponseProfileConfig, error) {
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

	profileCfg, err := p.profileCfgRepo.GetByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ProfileConfigNotFound
		}

		return nil, err
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

	profileConfigConv := req.ToModel(FormatConfigValueByte, req.ConfigName, profileCfg.ID)

	err = p.profileCfgRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = p.profileCfgRepo.Update(ctx, profileConfigConv)
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

	resp := &usecase.ResponseProfileConfig{
		ID:          profileConfigConv.ID,
		ProfileID:   profileConfigConv.ProfileID,
		ConfigName:  profileConfigConv.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      profileConfigConv.Status,
		Days:        days,
		Token:       req.Token,
	}
	return resp, nil
}
