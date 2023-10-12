package rapi

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/parse"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *Presenter) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUlid(id); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	notif, err := p.notificationUsecase.UpdateStatus(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.NotificationNotFound) {
			err = _error.HttpErrString("notifikasi tidak ditemukan", response.CM01)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	notifResp := &schema.ResponseNotification{
		ID:           notif.ID,
		ProfileID:    notif.ProfileID,
		UserConfigID: notif.UserConfigID,
		Message:      notif.Message,
		Title:        notif.Title,
		Icon:         notif.Title,
		Status:       notif.Status,
		CreatedAt:    notif.CreatedAt,
	}
	parse.SuccessResponseEncode(w, notifResp, "update notification successfully")
}

func (p *Presenter) DeleteByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUlid(id); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	err := p.notificationUsecase.DeleteByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	parse.SuccessResponseEncode(w, nil, "deleted notifikasi successfully")
}

func (p *Presenter) DeleteAllProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	err := p.notificationUsecase.DeleteAllByProfileID(r.Context(), profileID)
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	parse.SuccessResponseEncode(w, nil, "deleted semua notifikasi successfully")
}

func (p *Presenter) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	notifications, cursorResp, err := p.notificationUsecase.GetAllByProfileID(r.Context(), &usecase.RequestGetAllNotification{
		ID:        cursor,
		ProfileID: profileID,
		Order:     order,
	})
	if err != nil {
		parse.ErrorResponseEncode(w, err)
		return
	}

	if notifications == nil {
		parse.SuccessResponseEncode(w, map[string]any{
			"cursor":       "",
			"notification": nil,
		}, "data notification")
	}

	var notifResps []schema.ResponseNotification
	var notifResp *schema.ResponseNotification

	for _, notification := range *notifications {
		notifResp = &schema.ResponseNotification{
			ID:           notification.ID,
			ProfileID:    notification.ProfileID,
			UserConfigID: notification.UserConfigID,
			Message:      notification.Message,
			Title:        notification.Title,
			Icon:         notification.Title,
			Status:       notification.Status,
			CreatedAt:    notification.CreatedAt,
		}

		notifResps = append(notifResps, *notifResp)
	}

	resp := map[string]any{
		"cursor":       cursorResp,
		"notification": &notifResps,
	}

	parse.SuccessResponseEncode(w, resp, "data notification")
}

func (p *Presenter) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUlid(id); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("notifikasi tidak ditemukan", response.CM01))
		return
	}
	if err := util.ParseUUID(profileID); err != nil {
		parse.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	notif, err := p.notificationUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.NotificationNotFound) {
			err = _error.HttpErrString("notifikasi tidak ditemukan", response.CM01)
		}
		parse.ErrorResponseEncode(w, err)
		return
	}

	notifResp := &schema.ResponseNotification{
		ID:           notif.ID,
		ProfileID:    notif.ProfileID,
		UserConfigID: notif.UserConfigID,
		Message:      notif.Message,
		Title:        notif.Title,
		Icon:         notif.Title,
		Status:       notif.Status,
		CreatedAt:    notif.CreatedAt,
	}

	parse.SuccessResponseEncode(w, notifResp, "data notifikasi")
}
