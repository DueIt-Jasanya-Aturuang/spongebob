package config

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

func NewPgConn() *sql.DB {
	fDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		PgHost, PgPort, PgUser, PgPass, PgName, PgSSL)

	log.Info().Msgf("postgres config %v", fDB)

	db, err := sql.Open("postgres", fDB)
	if err != nil {
		log.Err(err).Msg("cannot open db")
	}

	ctx, cancel := context.WithTimeout(context.Background(), pgPingTimeOut)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Err(err).Msg("cannot ping db")
	}

	db.SetMaxIdleConns(setMaxIdleConnsDB)
	db.SetMaxOpenConns(setMaxOpenConnsDB)
	db.SetConnMaxIdleTime(SetConnMaxIdleTimeDB)
	db.SetConnMaxLifetime(setConnMaxLifetimeDB)

	log.Info().Msgf("connection postgres successfully : %s", PgName)
	return db
}
