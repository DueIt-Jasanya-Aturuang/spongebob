package utils

import (
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"
)

func InitDocker() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Err(err).Msgf("Could not construct pool: %s", err)
		os.Exit(1)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Err(err).Msgf("Could not connect to Docker: %s", err)
		os.Exit(1)
	}
	return pool
}
