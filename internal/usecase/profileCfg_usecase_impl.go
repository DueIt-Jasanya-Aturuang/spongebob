package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/dtoconv"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/format"
)

type ProfileCfgUsecaseImpl struct {
	profileRepo    repository.ProfileRepo
	profileCfgRepo repository.ProfileCfgRepo
	ctxTimeout     time.Duration
}

func NewProfileCfgUsecaseImpl(
	profileRepo repository.ProfileRepo,
	profileCfgRepo repository.ProfileCfgRepo,
	ctxTimeout time.Duration,
) usecase.ProfileCfgUsecase {
	return &ProfileCfgUsecaseImpl{
		profileRepo:    profileRepo,
		profileCfgRepo: profileCfgRepo,
		ctxTimeout:     ctxTimeout,
	}
}

func (u *ProfileCfgUsecaseImpl) CreateProfileCfg(c context.Context, req dto.CreateProfileCfgReq) (profileCfgResp *dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("profile user id : %s | user id %s", profile.UserID, req.UserID)
	if profile.UserID != req.UserID {
		return nil, exception.Err401Msg
	}

	formatConfigValue, err := format.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		return nil, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return nil, err
	}

	profileCfgRepoUOW := u.profileCfgRepo.UoW()
	err = profileCfgRepoUOW.StartTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		if errEndTx := profileCfgRepoUOW.EndTx(err); errEndTx != nil {
			err = errEndTx
			profileCfgResp = nil
		}
	}()

	profileCfg := dtoconv.CreateProfileCfgToModel(req, FormatConfigValueByte)
	err = u.profileCfgRepo.StoreProfileCfg(ctx, profileCfg)
	if err != nil {
		return nil, err
	}

	profileCfgResp = &dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   profileCfg.ProfileID,
		ConfigName:  profileCfg.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      profileCfg.Status,
	}
	return profileCfgResp, nil
}

func (u *ProfileCfgUsecaseImpl) GetProfileCfgByNameAndID(c context.Context, req dto.GetProfileCfgReq) (profileCfgResp *dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return nil, err
	}
	if profile.UserID != req.UserID {
		return nil, exception.Err401Msg
	}

	profileCfg, err := u.profileCfgRepo.GetProfileCfgByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		return nil, err
	}

	formatConfigValue := map[string]any{}
	err = json.Unmarshal([]byte(profileCfg.ConfigValue), &formatConfigValue)
	if err != nil {
		return nil, err
	}

	var configValue string
	switch req.ConfigName {
	case "DAILY_NOTIFY":
		configValue = fmt.Sprintf("%s %s", formatConfigValue["config_time_user"], formatConfigValue["config_timezone_user"])
	case "MONTHLY_PERIOD":
		configValue = fmt.Sprintf("%s", formatConfigValue["config_date"])
	}

	profileCfgResp = &dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   profileCfg.ProfileID,
		ConfigName:  profileCfg.ConfigName,
		ConfigValue: configValue,
		Status:      profileCfg.Status,
	}
	return profileCfgResp, nil
}

func (u *ProfileCfgUsecaseImpl) UpdateProfileCfg(c context.Context, req dto.UpdateProfileCfgReq) (profileCfgResp *dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("data : %v", profile)
	if profile.UserID != req.UserID {
		return nil, exception.Err401Msg
	}

	profileCfg, err := u.profileCfgRepo.GetProfileCfgByNameAndID(ctx, profile.ProfileID, req.ConfigName)
	if err != nil {
		return nil, err
	}

	formatConfigValue, err := format.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		return nil, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return nil, err
	}

	profileCfgRepoUOW := u.profileCfgRepo.UoW()
	err = profileCfgRepoUOW.StartTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		errEndTx := profileCfgRepoUOW.EndTx(err)
		if errEndTx != nil {
			err = errEndTx
			profileCfgResp = nil
		}
	}()

	profileCfgConv := dtoconv.UpdateProfileCfgToModel(req, FormatConfigValueByte, req.ConfigName, profileCfg.ID)
	err = u.profileCfgRepo.UpdateProfileCfg(ctx, profileCfgConv)
	if err != nil {
		return nil, err
	}

	profileCfgResp = &dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   req.ProfileID,
		ConfigName:  req.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      req.Status,
	}
	return profileCfgResp, nil
}
