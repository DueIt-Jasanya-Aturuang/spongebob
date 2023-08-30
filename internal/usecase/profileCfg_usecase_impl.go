package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
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

func (u *ProfileCfgUsecaseImpl) CreateProfileCfg(
	c context.Context, req dto.CreateProfileCfgReq,
) (profileCfgResp dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	if profile.UserID != req.UserID {
		return dto.ProfileCfgResp{}, model.ErrUnauthorization
	}

	formatConfigValue, err := helpers.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		if errors.Is(err, model.ErrBadInput) {
			return dto.ProfileCfgResp{}, response.Err400(map[string][]string{
				"config_value": {
					"invalid layout time or invalid iana timezone",
				},
			}, err)
		}
		return dto.ProfileCfgResp{}, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	err = u.profileRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}
	defer func() {
		if errEndTx := u.profileRepo.EndTx(err); errEndTx != nil {
			err = errEndTx
			profileCfgResp = dto.ProfileCfgResp{}
		}
	}()

	profileCfg := helpers.CreateProfileCfgToModel(req, FormatConfigValueByte)
	err = u.profileCfgRepo.StoreProfileCfg(ctx, profileCfg)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	profileCfgResp = dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   profileCfg.ProfileID,
		ConfigName:  profileCfg.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      profileCfg.Status,
	}

	return profileCfgResp, nil
}

func (u *ProfileCfgUsecaseImpl) GetProfileCfgByNameAndID(
	c context.Context, req dto.GetProfileCfgReq,
) (profileCfgResp dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	if profile.UserID != req.UserID {
		return dto.ProfileCfgResp{}, model.ErrUnauthorization
	}

	profileCfg, err := u.profileCfgRepo.GetProfileCfgByNameAndID(ctx, req.ProfileID, req.ConfigName)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	formatConfigValue := map[string]any{}
	err = json.Unmarshal([]byte(profileCfg.ConfigValue), &formatConfigValue)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	var configValue string
	switch req.ConfigName {
	case "DAILY_NOTIFY":
		configValue = fmt.Sprintf("%s %s", formatConfigValue["config_time_user"], formatConfigValue["config_timezone_user"])
	case "MONTHLY_PERIOD":
		configValue = fmt.Sprintf("%s", formatConfigValue["config_date"])
	}

	profileCfgResp = dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   profileCfg.ProfileID,
		ConfigName:  profileCfg.ConfigName,
		ConfigValue: configValue,
		Status:      profileCfg.Status,
	}

	return profileCfgResp, nil
}

func (u *ProfileCfgUsecaseImpl) UpdateProfileCfg(
	c context.Context, req dto.UpdateProfileCfgReq,
) (profileCfgResp dto.ProfileCfgResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	if profile.UserID != req.UserID {
		return dto.ProfileCfgResp{}, model.ErrUnauthorization
	}

	profileCfg, err := u.profileCfgRepo.GetProfileCfgByNameAndID(ctx, profile.ProfileID, req.ConfigName)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	formatConfigValue, err := helpers.ConfigValue(req.ConfigName, req.Value, req.IanaTimezone, req.Days)
	if err != nil {
		if errors.Is(err, model.ErrBadInput) {
			return dto.ProfileCfgResp{}, response.Err400(map[string][]string{
				"config_value": {
					"invalid layout time or invalid iana timezone",
				},
			}, err)
		}
		return dto.ProfileCfgResp{}, err
	}
	formatConfigValue["token"] = req.Token

	FormatConfigValueByte, err := json.Marshal(formatConfigValue)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	err = u.profileRepo.StartTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}
	defer func() {
		if errEndTx := u.profileRepo.EndTx(err); errEndTx != nil {
			err = errEndTx
			profileCfgResp = dto.ProfileCfgResp{}

		}
	}()

	profileCfgConv := helpers.UpdateProfileCfgToModel(req, FormatConfigValueByte, req.ConfigName, profileCfg.ID)
	err = u.profileCfgRepo.UpdateProfileCfg(ctx, profileCfgConv)
	if err != nil {
		return dto.ProfileCfgResp{}, err
	}

	profileCfgResp = dto.ProfileCfgResp{
		ID:          profileCfg.ID,
		ProfileID:   req.ProfileID,
		ConfigName:  req.ConfigName,
		ConfigValue: req.ConfigValue,
		Status:      req.Status,
	}

	return profileCfgResp, nil
}
