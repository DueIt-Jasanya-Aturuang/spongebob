package profileConfig_usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (p *ProfileConfigUsecaseImpl) GetByNameAndID(ctx context.Context, req *usecase.RequestGetProfileConfig) (*usecase.ResponseProfileConfig, error) {
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

	profileConfig, err := p.profileCfgRepo.GetByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ProfileConfigNotFound
		}

		return nil, err
	}

	formatConfigValue := map[string]any{}
	err = json.Unmarshal([]byte(profileConfig.ConfigValue), &formatConfigValue)
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

	resp := &usecase.ResponseProfileConfig{
		ID:          profileConfig.ID,
		ProfileID:   profileConfig.ProfileID,
		ConfigName:  profileConfig.ConfigName,
		ConfigValue: configValue,
		Status:      profileConfig.Status,
		Days:        days,
		Token:       formatConfigValue["token"].(string),
	}
	return resp, nil
}
