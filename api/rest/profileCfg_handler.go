package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/_usecase"
)

type ProfileCfgHandler struct {
	profileCfgUsecase domain.ProfileConfigUsecase
}

func NewProfileCfgHandler(
	profileCfgUsecase domain.ProfileConfigUsecase,
) *ProfileCfgHandler {
	return &ProfileCfgHandler{
		profileCfgUsecase: profileCfgUsecase,
	}
}

func (h *ProfileCfgHandler) CreateProfileCfg(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateProfileConfig)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	req.UserID = r.Header.Get("User-ID")
	req.ProfileID = r.Header.Get("Profile-ID")

	err = validation.CreateProfileCfg(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		req.Value = strings.Split(req.ConfigValue, " ")[0]
		req.IanaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		req.Value = req.ConfigValue
	}

	profileCfg, err := h.profileCfgUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		if errors.Is(err, _usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM05], response.CM05)
		}
		// conflict
		if errors.Is(err, _usecase.ProfileConfigIsExist) {
			err = _error.HttpErrString("profile config sudah dibuat", response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	w.Header().Del("Profile-ID")
	helper.SuccessResponseEncode(w, profileCfg, "created profile config successfully")
}

func (h *ProfileCfgHandler) GetProfileCfgByNameAndID(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestGetProfileConfig)

	req.ConfigName = chi.URLParam(r, "config-name")
	req.ProfileID = r.Header.Get("Profile-ID")
	req.UserID = r.Header.Get("User-ID")

	err := validation.GetProfileCfgValidation(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	profileCfg, err := h.profileCfgUsecase.GetByNameAndID(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		if errors.Is(err, _usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM05], response.CM05)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	w.Header().Del("Profile-ID")
	helper.SuccessResponseEncode(w, profileCfg, "data profile config")
}

func (h *ProfileCfgHandler) UpdateProfileCfg(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("head | %s", r.Header.Get("test"))
	req := new(domain.RequsetUpdateProfileConfig)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	req.ConfigName = chi.URLParam(r, "config-name")
	req.ProfileID = r.Header.Get("Profile-ID")
	req.UserID = r.Header.Get("User-ID")

	err = validation.UpdateProfileCfgValidate(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		req.Value = strings.Split(req.ConfigValue, " ")[0]
		req.IanaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		req.Value = req.ConfigValue
	}

	profileCfg, err := h.profileCfgUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		if errors.Is(err, _usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM05], response.CM05)
		}
		if errors.Is(err, _usecase.ProfileConfigNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	w.Header().Del("Profile-ID")
	helper.SuccessResponseEncode(w, profileCfg, "update profile config successfully")
}
