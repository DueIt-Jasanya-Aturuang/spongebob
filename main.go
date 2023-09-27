package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest"
	cusmiddleware "github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/middleware"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/pkg/_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/pkg/_usecase"
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

	redisConn := infra.NewRedisConn()
	defer func() {
		err := redisConn.Client.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE REDIS CONN")
		}
	}()

	minioConn, err := infra.NewMinioConn(infra.MinIoEndpoint, infra.MinIoID, infra.MinIoSecretKey, infra.MinIoSSL)
	if err != nil {
		panic(err)
	}

	uow := _repository.NewUnitOfWorkRepositoryImpl(pgConn)
	profileRepoCfg := _repository.NewProfileConfigRepoImpl(uow)
	profileRepo := _repository.NewProfileRepoImpl(uow)
	userRepo := _repository.NewUserRepoImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioConn)

	accountUsecase := _usecase.NewAccountUsecaseImpl(profileRepo, userRepo, minioRepo)
	profileCfgUsecase := _usecase.NewProfileConfigUsecaseImpl(profileRepo, profileRepoCfg)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cusmiddleware.IPMiddleware)

	accountHandler := rest.NewAccountHandler(accountUsecase)
	profileCfgHandler := rest.NewProfileCfgHandler(profileCfgUsecase)

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
	})

	err = http.ListenAndServe(infra.AppPort, r)
	if err != nil {
		panic(err)
	}
}
