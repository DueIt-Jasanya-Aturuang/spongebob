package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(
	accountUsecase usecase.AccountUsecase,
) *AccountHandler {
	return &AccountHandler{
		accountUsecase: accountUsecase,
	}
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	req := new(dto.UpdateAccountReq)

	err := response.ParserMultipartForm(r, req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	req.UserID = r.Header.Get("User-ID")
	req.ProfileID = chi.URLParam(r, "profile-id")

	err = validation.UpdateAccountValidate(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	user, profile, err := h.accountUsecase.UpdateAccount(r.Context(), req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: map[string]any{
			"user":    user,
			"profile": profile,
		},
	}

	response.NewSucc(w, r, resp, 200)
}
