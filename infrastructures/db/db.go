package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPgConn() *sql.DB {
	fDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PgHost, config.PgPort, config.PgUser, config.PgPass, config.PgName, config.PgSSL)

	log.Info().Msgf("postgres config %v", fDB)

	db, err := sql.Open("postgres", fDB)
	if err != nil {
		log.Err(err).Msg("cannot open db")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Err(err).Msg("cannot ping db")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Info().Msgf("connection postgres successfully : %s", config.PgName)
	return db
}
