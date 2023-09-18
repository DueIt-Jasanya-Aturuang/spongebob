package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/_usecase"
)

type AccountHandler struct {
	accountUsecase domain.AccountUsecase
}

func NewAccountHandler(
	accountUsecase domain.AccountUsecase,
) *AccountHandler {
	return &AccountHandler{
		accountUsecase: accountUsecase,
	}
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateAccount)

	err := helper.ParserMultipartForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	req.UserID = r.Header.Get("User-ID")
	req.ProfileID = chi.URLParam(r, "profile-id")

	err = validation.UpdateAccountValidate(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	user, profile, err := h.accountUsecase.UpdateAccount(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		if errors.Is(err, _usecase.UserNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		if errors.Is(err, _usecase.ProfileUserIDAndReqUserIDNotMatch) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM05], response.CM05)
		}
		if errors.Is(err, _usecase.PhoneNumberIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"phone_number": {
					err.Error(),
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	data := map[string]any{
		"user":    user,
		"profile": profile,
	}

	helper.SuccessResponseEncode(w, data, "successfully update akun")
}

func (h *AccountHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestGetProfile)

	userId := r.Header.Get("User-ID")

	req.UserID = userId

	err := validation.GetProfileValidation(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	profile, err := h.accountUsecase.GetProfileByUserID(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.ProfileNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, profile, "data profile")
}

func (h *AccountHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateProfile)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	err = validation.CreateProfileValidation(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	profile, err := h.accountUsecase.CreateProfile(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.UserNotFound) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, profile, "successully membuat profile")
}