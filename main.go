package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/presentation/rapi"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/minio_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/notification_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/profileConfig_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/profile_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/uow_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository/user_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/account_usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/notification_usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/profileConfig_usecase"
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

	minioConn, err := infra.NewMinioConn(infra.MinIoEndpoint, infra.MinIoID, infra.MinIoSecretKey, infra.MinIoSSL)
	if err != nil {
		panic(err)
	}

	depen := dependency(pgConn, minioConn)

	var id *string
	day := time.Now().Day()
	go func() {
		for range time.Tick(1 * time.Minute) {
			id, err = depen.ProfileConfigUsecase.SchedulerMonthlyPeriode(context.Background(), day, id)
			if err != nil {
				// log.Warn().Msgf("failed push notification monthly notif user | err : %v", err)
			}

			log.Info().Msgf("id user : %v", id)
		}
	}()

	go func() {
		for range time.Tick(1 * time.Minute) {
			fmt.Println(fmt.Sprintf("%02d:%02d", time.Now().UTC().Hour(), time.Now().UTC().Minute()))
			err = depen.ProfileConfigUsecase.SchedulerDailyNotify(
				context.Background(),
				fmt.Sprintf("%02d:%02d", time.Now().UTC().Hour(), time.Now().UTC().Minute()),
				strings.ToLower(time.Now().UTC().Weekday().String()),
			)

			if err != nil {
				log.Warn().Msgf("failed push notification daily notif user | err : %v", err)
			}
		}
	}()

	httpServer, err := rapi.NewPresenter(rapi.PresenterConfig{
		Dependency: depen,
	})
	if err != nil {
		log.Fatal().Msgf("creating new presenter: %s", err.Error())
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		log.Info().Msg("Interrupt signal received, exiting...")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Minute)
		defer shutdownCancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Err(err).Msg("shutting down HTTP server")
		}
	}()

	log.Info().Msgf("Server is running on port %s", infra.AppPort)
	err = httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Msgf("serving HTTP server: %s", err.Error())
	}
}

func dependency(db *sql.DB, client *minio.Client) *rapi.Dependency {
	uow := uow_repository.NewUnitOfWorkRepositoryImpl(db)
	profileConfigRepo := profileConfig_repository.NewProfileConfigRepositoryImpl(uow)
	profileRepo := profile_repository.NewProfileRepositoryImpl(uow)
	userRepo := user_repository.NewUserRepositoryImpl(uow)
	minioRepo := minio_repository.NewMinioRepositoryImpl(client)
	notificationRepo := notification_repository.NewNotificationRepositoryImpl(uow)

	accountUsecase := account_usecase.NewAccountUsecaseImpl(profileRepo, userRepo, minioRepo)
	profileCfgUsecase := profileConfig_usecase.NewProfileConfigUsecaseImpl(profileRepo, profileConfigRepo, notificationRepo)
	notificationUsecase := notification_usecase.NewNotificationUsecaseImpl(notificationRepo)

	return &rapi.Dependency{
		AccountUsecase:       accountUsecase,
		NotificationUsecase:  notificationUsecase,
		ProfileConfigUsecase: profileCfgUsecase,
	}
}
