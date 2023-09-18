package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/repository"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest"
	cusmiddleware "github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/middleware"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/_usecase"
)

func main() {
	config.LogInit()
	config.EnvInit()
	pgConn := config.NewPgConn()
	defer func() {
		err := pgConn.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE DB")
		}
	}()

	redisConn := config.NewRedisConn()
	defer func() {
		err := redisConn.Client.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE REDIS CONN")
		}
	}()

	minioConn, err := config.NewMinioConn(config.MinIoEndpoint, config.MinIoID, config.MinIoSecretKey, config.MinIoSSL)
	if err != nil {
		panic(err)
	}

	uow := repository.NewUnitOfWorkImpl(pgConn)
	profileRepoCfg := repository.NewProfileCfgRepoImpl(uow)
	profileRepo := repository.NewProfileRepoImpl(uow)
	userRepo := repository.NewUserRepoImpl(uow)
	minioRepo := repository.NewMinioImpl(minioConn)

	accountUsecase := _usecase.NewAccountUsecaseImpl(profileRepo, userRepo, minioRepo, 10*time.Second)
	profileUsecase := _usecase.NewProfileUsecaseImpl(profileRepo, userRepo, 10*time.Second)
	profileCfgUsecase := _usecase.NewProfileCfgUsecaseImpl(profileRepo, profileRepoCfg, 10*time.Second)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cusmiddleware.IPMiddleware)

	accountHandler := rest.NewAccountHandler(accountUsecase)
	profileHandler := rest.NewProfileHandler(profileUsecase)
	profileCfgHandler := rest.NewProfileCfgHandler(profileCfgUsecase)

	r.Put("/account/{profile-id}", accountHandler.UpdateAccount)
	r.Get("/account/profile", profileHandler.GetProfileByID)
	r.Post("/account/profile", profileHandler.StoreProfile)
	r.Post("/account/profile-config/{profile-id}", profileCfgHandler.CreateProfileCfg)
	r.Get("/account/profile-config/{profile-id}/{config-name}", profileCfgHandler.GetProfileCfgByNameAndID)
	r.Put("/account/profile-config/{profile-id}/{config-name}", profileCfgHandler.UpdateProfileCfg)

	err = http.ListenAndServe(config.AppPort, r)
	if err != nil {
		panic(err)
	}
}
