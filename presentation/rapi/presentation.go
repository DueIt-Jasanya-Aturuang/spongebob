package rapi

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	cusmiddleware "github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi/middleware"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

type Presenter struct {
	accountUsecase       usecase.AccountUsecase
	notificationUsecase  usecase.NotificationUsecase
	profileConfigUsecase usecase.ProfileConfigUsecase
}

type Dependency struct {
	AccountUsecase       usecase.AccountUsecase
	NotificationUsecase  usecase.NotificationUsecase
	ProfileConfigUsecase usecase.ProfileConfigUsecase
}

type PresenterConfig struct {
	Dependency *Dependency
}

func NewPresenter(config PresenterConfig) (*http.Server, error) {
	presenter := &Presenter{
		accountUsecase:       config.Dependency.AccountUsecase,
		notificationUsecase:  config.Dependency.NotificationUsecase,
		profileConfigUsecase: config.Dependency.ProfileConfigUsecase,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cusmiddleware.IPMiddleware)
	r.Use(cusmiddleware.CheckApiKey)

	r.Group(func(r chi.Router) {
		r.Use(cusmiddleware.SetAuthorization)

		r.Get("/account/profile", presenter.GetProfileByUserID)
		r.Post("/account/profile/5410801c-faaf-4776-95be-56472e044820", presenter.CreateProfile)
		r.Get("/account/profile/5410801c-faaf-4776-95be-56472e044820", presenter.GetProfileByUserID)

		r.Post("/account/otorisasi", presenter.Otorisasi)

		r.Put("/account", presenter.UpdateAccount)

		r.Post("/account/profile-config", presenter.CreateProfileCfg)
		r.Get("/account/profile-config/{config-name}", presenter.GetProfileCfgByNameAndID)
		r.Put("/account/profile-config/{config-name}", presenter.UpdateProfileCfg)

		r.Get("/account/notif", presenter.GetAllByProfileID)
		r.Get("/account/notif/{id}", presenter.GetByIDAndProfileID)
		r.Put("/account/notif/status/{id}", presenter.UpdateStatus)
		r.Delete("/account/notif/{id}", presenter.DeleteByIDAndProfileID)
		r.Delete("/account/notif", presenter.DeleteAllProfileID)
	})

	server := &http.Server{
		Addr:              infra.AppPort,
		Handler:           r,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	return server, nil
}
