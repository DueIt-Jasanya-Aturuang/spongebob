package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

type NotificationHandlerImpl struct {
	notifUsecase domain.NotificationUsecase
}

func NewNotificationHandlerImpl(
	notifUsecase domain.NotificationUsecase,
) *NotificationHandlerImpl {
	return &NotificationHandlerImpl{
		notifUsecase: notifUsecase,
	}
}

func (n *NotificationHandlerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	if _, err := ulid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}

	notif, err := n.notifUsecase.UpdateStatus(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.NotificationNotFound) {
			err = _error.HttpErrString("notifikasi tidak ditemukan", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, notif, "update notification successfully")
}

func (n *NotificationHandlerImpl) DeleteByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	if _, err := ulid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}

	err := n.notifUsecase.DeleteByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted notifikasi successfully")
}

func (n *NotificationHandlerImpl) DeleteAllProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	err := n.notifUsecase.DeleteAllByProfileID(r.Context(), profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted semua notifikasi successfully")
}

func (n *NotificationHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")
	order, operation := helper.GetOrder(order)

	req := &domain.RequestGetAllByPaginate{
		ProfileID: profileID,
		ID:        cursor,
		Operation: operation,
		Order:     order,
	}

	notification, cursorResp, err := n.notifUsecase.GetAllByProfileID(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := map[string]any{
		"cursor":       cursorResp,
		"notification": notification,
	}

	helper.SuccessResponseEncode(w, resp, "data notification")
}

func (n *NotificationHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	if _, err := ulid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}

	notif, err := n.notifUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.NotificationNotFound) {
			err = _error.HttpErrString("notifikasi tidak ditemukan", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, notif, "data notifikasi")
}
