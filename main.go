package main

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/rs/zerolog/log"
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
	repository.NewProfileCfgRepoImpl(uow)
	repository.NewProfileRepoImpl(uow)
	repository.NewUserRepoImpl(uow)
	repository.NewMinioImpl(minioConn)
}
