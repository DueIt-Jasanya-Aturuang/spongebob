package rapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/parse"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *Presenter) CreateProfileCfg(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(userID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateProfileConfig)

	err := parse.DecodeJson(r, req)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	err = req.Validate()
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	var value string
	var ianaTimezone string
	if req.ConfigName == "DAILY_NOTIFY" {
		value = strings.Split(req.ConfigValue, " ")[0]
		ianaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		value = req.ConfigValue
	}

	profileConfig, err := p.profileConfigUsecase.Create(r.Context(), &usecase.RequestCreateProfileConfig{
		ConfigValue:  req.ConfigValue,
		Days:         req.Days,
		ConfigName:   req.ConfigName,
		Status:       req.Status,
		Token:        req.Token,
		UserID:       userID,
		ProfileID:    profileID,
		Value:        value,
		IanaTimezone: ianaTimezone,
	})
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString("profile tidak ditemukan", response.CM04)
		}
		if errors.Is(err, usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		if errors.Is(err, usecase.ProfileConfigIsExist) {
			err = _error.HttpErrString("profile config sudah dibuat", response.CM06)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseProfileConfig{
		ID:          profileConfig.ID,
		ProfileID:   profileConfig.ProfileID,
		ConfigName:  profileConfig.ConfigName,
		ConfigValue: profileConfig.ConfigValue,
		Status:      profileConfig.Status,
		Days:        profileConfig.Days,
		Token:       profileConfig.Token,
	}
	parse.SuccessResponseEncode(w, resp, "created profile config successfully")
}

func (p *Presenter) GetProfileCfgByNameAndID(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	profileID := r.Header.Get("Profile-ID")
	configName := chi.URLParam(r, "config-name")

	if err := util.ParseUUID(userID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	if configName != "DAILY_NOTIFY" && configName != "MONTHLY_PERIOD" {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid config name", response.CM01))
		return
	}
	profileConfig, err := p.profileConfigUsecase.GetByNameAndID(r.Context(), &usecase.RequestGetProfileConfig{
		UserID:     userID,
		ConfigName: configName,
		ProfileID:  profileID,
	})
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString("profile tidak ditemukan", response.CM04)
		}
		if errors.Is(err, usecase.ProfileConfigNotFound) {
			err = _error.HttpErrString("profile config tidak ditemukan", response.CM01)
		}
		if errors.Is(err, usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseProfileConfig{
		ID:          profileConfig.ID,
		ProfileID:   profileConfig.ProfileID,
		ConfigName:  profileConfig.ConfigName,
		ConfigValue: profileConfig.ConfigValue,
		Status:      profileConfig.Status,
		Days:        profileConfig.Days,
		Token:       profileConfig.Token,
	}
	parse.SuccessResponseEncode(w, resp, "data profile config")
}

func (p *Presenter) UpdateProfileCfg(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(userID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequsetUpdateProfileConfig)

	err := parse.DecodeJson(r, req)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	configName := chi.URLParam(r, "config-name")

	err = req.Validate(configName)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	var value string
	var ianaTimezone string
	if configName == "DAILY_NOTIFY" {
		value = strings.Split(req.ConfigValue, " ")[0]
		ianaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		value = req.ConfigValue
	}

	profileConfig, err := p.profileConfigUsecase.Update(r.Context(), &usecase.RequsetUpdateProfileConfig{
		ConfigValue:  req.ConfigValue,
		Days:         req.Days,
		Status:       req.Status,
		Token:        req.Token,
		ProfileID:    profileID,
		UserID:       userID,
		ConfigName:   configName,
		Value:        value,
		IanaTimezone: ianaTimezone,
	})
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString("profile tidak ditemukan", response.CM04)
		}
		if errors.Is(err, usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM05], response.CM04)
		}
		if errors.Is(err, usecase.ProfileConfigNotFound) {
			err = _error.HttpErrString("profile config tidak ditemukan", response.CM01)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseProfileConfig{
		ID:          profileConfig.ID,
		ProfileID:   profileConfig.ProfileID,
		ConfigName:  profileConfig.ConfigName,
		ConfigValue: profileConfig.ConfigValue,
		Status:      profileConfig.Status,
		Days:        profileConfig.Days,
		Token:       profileConfig.Token,
	}
	parse.SuccessResponseEncode(w, resp, "update profile config successfully")
}
