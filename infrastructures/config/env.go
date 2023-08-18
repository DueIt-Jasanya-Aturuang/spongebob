package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func EnvInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Err(err).Msg("cannot load env")
		os.Exit(1)
	}

	AppPort = os.Getenv("APPLICATION_PORT")
	AppStatus = os.Getenv("APPLICATION_STATUS")
	AppGracefull = os.Getenv("APPLICATION_GRACEFUL_MAX_SECOND")

	PgHost = os.Getenv("DB_POSTGRESQL_HOST")
	PgPort = os.Getenv("DB_POSTGRESQL_PORT")
	PgUser = os.Getenv("DB_POSTGRESQL_USER")
	PgPass = os.Getenv("DB_POSTGRESQL_PASS")
	PgName = os.Getenv("DB_POSTGRESQL_NAME")
	PgSSL = os.Getenv("DB_POSTGRESQL_SSL")

	dbInt, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisDb = dbInt
	RedisPass = os.Getenv("REDIS_PASS")

	MinIoAccessKey = os.Getenv("MINIO_ACCESSKEY")
	MinIoSecretKey = os.Getenv("MINIO_SECRETKEY")
	MinIoEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinIoPort = os.Getenv("MINIO_PORT")
	MinIoBucket = os.Getenv("MINIO_BUCKET")
	MinIoSSL = os.Getenv("MINIO_SSL")

	DefaultImage = os.Getenv("DEFAULT_DEFAULT_IMAGE")
	AesCFB = os.Getenv("DEFAULT_AES_CFB_KEY")
	AesCBC = os.Getenv("DEFAULT_AES_CBC_KEY")
	AesCBCIV = os.Getenv("DEFAULT_AES_CBC_IV_KEY")

	log.Info().Msg("config initialization successfully")
}
