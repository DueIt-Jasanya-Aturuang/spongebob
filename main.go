package main

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
)

func main() {
	config.EnvInit()
	pgConn := config.NewPgConn()
	minioConn, err := config.NewMinioConn(config.MinIoEndpoint, config.MinIoID, config.MinIoSecretKey, config.MinIoSSL)
	if err != nil {
		panic(err)
	}
	_ = config.NewRedisConn()

	repository.NewMinioImpl(minioConn)
	repository.NewProfileCfgRepoImpl(pgConn)
	repository.NewProfileRepoImpl(pgConn)
	repository.NewUserRepoImpl(pgConn)
}
