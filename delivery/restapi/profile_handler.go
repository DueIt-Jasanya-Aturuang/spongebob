package restapi

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type ProfileHandler struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileHandler(
	profileUsecase usecase.ProfileUsecase,
) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
	}
}

func (h *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	req := new(dto.GetProfileReq)

	userId := r.Header.Get("User-ID")

	req.UserID = userId

	err := validation.GetProfileValidation(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	profile, err := h.profileUsecase.GetProfileByID(r.Context(), req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: profile,
	}

	response.NewSucc(w, r, resp, 200)
}

func (h *ProfileHandler) StoreProfile(w http.ResponseWriter, r *http.Request) {
	req := new(dto.StoreProfileReq)

	err := response.DecodeReq(r, req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	err = validation.StoreProfileValidation(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	profile, err := h.profileUsecase.StoreProfile(r.Context(), req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}
	resp := model.ResponseSuccess{
		Data: profile,
	}
	response.NewSucc(w, r, resp, 201)
}
