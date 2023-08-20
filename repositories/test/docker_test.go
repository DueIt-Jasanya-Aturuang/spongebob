package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repositories"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	db          *sql.DB
	profileRepo = repositories.NewProfileRepoImpl()
	userRepo    = repositories.NewUserRepoImpl()
)

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.Nop())
	pool := initDocker()

	// pulls an image, creates a container based on it and runs it
	resource, dbPg, url := postgresStart(pool)

	db = dbPg
	startMigration(url, db)
	// Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Err(err).Msgf("Could not purge resource: %s", err)
		os.Exit(1)
	}

	os.Exit(code)
}

func initDocker() *dockertest.Pool {
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

func postgresStart(dockerPool *dockertest.Pool) (*dockertest.Resource, *sql.DB, string) {
	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=pw",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dueit_db",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Err(err).Msgf("Could not start resource: %s", err)
		os.Exit(1)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:pw@%s/dueit_db?sslmode=disable", hostAndPort)

	log.Info().Msgf("Connecting to database on url: %s", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	dockerPool.MaxWait = 120 * time.Second
	var db *sql.DB
	if err = dockerPool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Err(err).Msgf("Could not connect to docker: %s", err)
		os.Exit(1)
	}

	return resource, db, databaseUrl
}

func startMigration(url string, db *sql.DB) {
	log.Info().Msgf("start create schema")
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS dueit;")
	if err != nil {
		log.Err(err).Msgf("Failed to create schema dueit: %s", err)
		os.Exit(1)
	}
	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS auth;")
	if err != nil {
		log.Err(err).Msgf("Failed to create schema auth: %s", err)
		os.Exit(1)
	}

	db, err = sql.Open("postgres", url+"&search_path=dueit")
	if err != nil {
		log.Err(err).Msgf("could not start db: %s", err)
		os.Exit(1)
	}
	log.Info().Msgf("start migrate")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Err(err).Msgf("could not init driver: %s", err)
		os.Exit(1)
	}

	migrates, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		log.Err(err).Msgf("could not apply the migration: %s", err)
		os.Exit(1)
	}
	migrates.Up()

	db, err = sql.Open("postgres", url+"&search_path=auth")
	if err != nil {
		log.Err(err).Msgf("could not start db: %s", err)
		os.Exit(1)
	}

	driver, err = postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Err(err).Msgf("could not init driver: %s", err)
		os.Exit(1)
	}

	migrates, err = migrate.NewWithDatabaseInstance(
		"file://../../../../Jasanya-Auth/migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		log.Err(err).Msgf("could not apply the migration: %s", err)
		os.Exit(1)
	}

	migrates.Up()
}
