package rest

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/validation"
)

type ProfileCfgHandler struct {
	profileCfgUsecase usecase.ProfileCfgUsecase
}

func NewProfileCfgHandler(
	profileCfgUsecase usecase.ProfileCfgUsecase,
) *ProfileCfgHandler {
	return &ProfileCfgHandler{
		profileCfgUsecase: profileCfgUsecase,
	}
}

func (h *ProfileCfgHandler) CreateProfileCfg(w http.ResponseWriter, r *http.Request) {
	req := new(dto.CreateProfileCfgReq)

	err := response.DecodeReq(r, req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	req.UserID = r.Header.Get("User-ID")
	req.ProfileID = chi.URLParam(r, "profile-id")

	err = validation.CreateProfileCfg(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		req.Value = strings.Split(req.ConfigValue, " ")[0]
		req.IanaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		req.Value = req.ConfigValue
	}

	profileCfg, err := h.profileCfgUsecase.CreateProfileCfg(r.Context(), *req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: profileCfg,
	}

	response.NewSucc(w, r, resp, 201)
}

func (h *ProfileCfgHandler) GetProfileCfgByNameAndID(w http.ResponseWriter, r *http.Request) {
	req := new(dto.GetProfileCfgReq)

	req.ConfigName = chi.URLParam(r, "config-name")
	req.ProfileID = chi.URLParam(r, "profile-id")
	req.UserID = r.Header.Get("User-ID")

	err := validation.GetProfileCfgValidation(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	profileCfg, err := h.profileCfgUsecase.GetProfileCfgByNameAndID(r.Context(), *req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: profileCfg,
	}

	response.NewSucc(w, r, resp, 200)
}

func (h *ProfileCfgHandler) UpdateProfileCfg(w http.ResponseWriter, r *http.Request) {
	req := new(dto.UpdateProfileCfgReq)

	err := response.DecodeReq(r, req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	req.ConfigName = chi.URLParam(r, "config-name")
	req.ProfileID = chi.URLParam(r, "profile-id")
	req.UserID = r.Header.Get("User-ID")

	err = validation.UpdateProfileCfgValidate(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		req.Value = strings.Split(req.ConfigValue, " ")[0]
		req.IanaTimezone = strings.Split(req.ConfigValue, " ")[1]
	} else {
		req.Value = req.ConfigValue
	}

	profileCfg, err := h.profileCfgUsecase.UpdateProfileCfg(r.Context(), *req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: profileCfg,
	}

	response.NewSucc(w, r, resp, 200)
}
