package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest"
	cusmiddleware "github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/middleware"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	repository2 "github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	_usecase2 "github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func main() {
	infra.LogInit()
	infra.EnvInit()
	pgConn := infra.NewPgConn()
	defer func() {
		err := pgConn.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE DB")
		}
	}()

	// redisConn := infra.NewRedisConn()
	// defer func() {
	// 	err := redisConn.Client.Close()
	// 	if err != nil {
	// 		log.Err(err).Msg("ERROR CLOSE REDIS CONN")
	// 	}
	// }()

	minioConn, err := infra.NewMinioConn(infra.MinIoEndpoint, infra.MinIoID, infra.MinIoSecretKey, infra.MinIoSSL)
	if err != nil {
		panic(err)
	}

	uow := repository2.NewUnitOfWorkRepositoryImpl(pgConn)
	profileRepoCfg := repository2.NewProfileConfigRepoImpl(uow)
	profileRepo := repository2.NewProfileRepoImpl(uow)
	userRepo := repository2.NewUserRepoImpl(uow)
	minioRepo := repository2.NewMinioImpl(minioConn)
	notificationRepo := repository2.NewNotificationRepoImpl(uow)

	accountUsecase := _usecase2.NewAccountUsecaseImpl(profileRepo, userRepo, minioRepo)
	profileCfgUsecase := _usecase2.NewProfileConfigUsecaseImpl(profileRepo, profileRepoCfg, notificationRepo)
	notificationUsecase := _usecase2.NewNotificationUsecaseImpl(notificationRepo)

	var id *string
	day := time.Now().Day()
	go func() {
		for range time.Tick(1 * time.Minute) {
			id, err = profileCfgUsecase.SchedulerMonthlyPeriode(context.Background(), day, id)
			if err != nil {
				// log.Warn().Msgf("failed push notification monthly notif user | err : %v", err)
			}

			log.Info().Msgf("id user : %v", id)
		}
	}()

	go func() {
		for range time.Tick(1 * time.Minute) {
			fmt.Println(fmt.Sprintf("%02d:%02d", time.Now().UTC().Hour(), time.Now().UTC().Minute()))
			err = profileCfgUsecase.SchedulerDailyNotify(context.Background(), domain.ProfileConfigScheduler{
				Day:  strings.ToLower(time.Now().UTC().Weekday().String()),
				Time: fmt.Sprintf("%02d:%02d", time.Now().UTC().Hour(), time.Now().UTC().Minute()),
			})

			if err != nil {
				log.Warn().Msgf("failed push notification daily notif user | err : %v", err)
			}
		}
	}()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cusmiddleware.IPMiddleware)
	r.Use(cusmiddleware.CheckApiKey)

	accountHandler := rest.NewAccountHandler(accountUsecase)
	profileCfgHandler := rest.NewProfileCfgHandler(profileCfgUsecase)
	notificationHandler := rest.NewNotificationHandlerImpl(notificationUsecase)

	r.Group(func(r chi.Router) {
		r.Use(cusmiddleware.SetAuthorization)

		r.Get("/account/profile", accountHandler.GetProfileByUserID)
		r.Post("/account/profile/5410801c-faaf-4776-95be-56472e044820", accountHandler.CreateProfile)
		r.Get("/account/profile/5410801c-faaf-4776-95be-56472e044820", accountHandler.GetProfileByUserID)

		r.Post("/account/otorisasi", accountHandler.Otorisasi)

		r.Put("/account", accountHandler.UpdateAccount)

		r.Post("/account/profile-config", profileCfgHandler.CreateProfileCfg)
		r.Get("/account/profile-config/{config-name}", profileCfgHandler.GetProfileCfgByNameAndID)
		r.Put("/account/profile-config/{config-name}", profileCfgHandler.UpdateProfileCfg)

		r.Get("/account/notif", notificationHandler.GetAllByProfileID)
		r.Get("/account/notif/{id}", notificationHandler.GetByIDAndProfileID)
		r.Put("/account/notif/status/{id}", notificationHandler.UpdateStatus)
		r.Delete("/account/notif/{id}", notificationHandler.DeleteByIDAndProfileID)
		r.Delete("/account/notif", notificationHandler.DeleteAllProfileID)
	})

	err = http.ListenAndServe(infra.AppPort, r)
	if err != nil {
		panic(err)
	}
}
