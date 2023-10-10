package rapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/parse"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *Presenter) UpdateAccount(w http.ResponseWriter, r *http.Request) {
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

	req := new(schema.RequestUpdateAccount)

	err := parse.ParserMultipartForm(r, req)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	err = req.Validate()
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	user, profile, err := p.accountUsecase.UpdateAccount(r.Context(), &usecase.RequestUpdateAccount{
		UserID:      userID,
		ProfileID:   profileID,
		FullName:    req.FullName,
		Gender:      req.Gender,
		Image:       req.Image,
		PhoneNumber: req.PhoneNumber,
		Quote:       req.Quote,
		Profesi:     req.Profesi,
	})
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		if errors.Is(err, usecase.UserNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		if errors.Is(err, usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		if errors.Is(err, usecase.PhoneNumberIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"phone_number": {
					err.Error(),
				},
			}, response.CM06)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	userResp := &schema.ResponseUser{
		ID:              user.ID,
		FullName:        user.FullName,
		Gender:          user.Gender,
		Image:           user.Image,
		Username:        user.Username,
		Email:           user.Email,
		EmailFormat:     user.EmailFormat,
		PhoneNumber:     user.PhoneNumber,
		EmailVerifiedAt: user.EmailVerifiedAt,
	}
	profileResp := &schema.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     profile.Quote,
		Profesi:   profile.Profesi,
	}

	resp := map[string]any{
		"user":    userResp,
		"profile": profileResp,
	}

	parse.SuccessResponseEncode(w, resp, "successfully update akun")
}

func (p *Presenter) GetProfileByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")

	if err := util.ParseUUID(userID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}

	profile, err := p.accountUsecase.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     profile.Quote,
		Profesi:   profile.Profesi,
	}
	parse.SuccessResponseEncode(w, resp, "data profile")
}

func (p *Presenter) Otorisasi(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")

	if err := util.ParseUUID(userID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}

	profile, err := p.accountUsecase.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ProfileNotFound) {
			err = _error.HttpErrString("invalid your profile", response.CM04)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Profile-ID", profile.ProfileID)
	w.WriteHeader(200)
	resp := response.HttpResponse{
		Status:  "success",
		Message: "ok",
		Errors:  nil,
		Data:    nil,
	}
	if errEncode := json.NewEncoder(w).Encode(resp); errEncode != nil {
		log.Err(errEncode).Msgf(util.LogErrEncode, resp, errEncode)
	}
}

func (p *Presenter) CreateProfile(w http.ResponseWriter, r *http.Request) {
	req := new(schema.RequestCreateProfile)
	err := parse.DecodeJson(r, req)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	if err = util.ParseUUID(req.UserID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid user id", response.CM04))
		return
	}

	profile, err := p.accountUsecase.CreateProfile(r.Context(), req.UserID)
	if err != nil {
		if errors.Is(err, usecase.UserNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     profile.Quote,
		Profesi:   profile.Profesi,
	}
	parse.SuccessResponseEncode(w, resp, "successully membuat profile")
}
