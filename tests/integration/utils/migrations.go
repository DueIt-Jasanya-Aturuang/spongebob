package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func StartMigration(url string, db *sql.DB) {
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
		"file://helper/jasanya-auth",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		log.Err(err).Msgf("could not apply the jasanya-auth: %s", err)
		os.Exit(1)
	}

	migrates.Up()
}
