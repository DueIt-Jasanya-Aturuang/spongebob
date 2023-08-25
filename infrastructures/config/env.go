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
		log.Err(err).Msg("CANNOT LOAD ENV FILE")
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

	dbInt, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisDB = dbInt
	RedisPass = os.Getenv("REDIS_PASS")

	miniosslBool, err := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	if err != nil {
		panic(err)
	}
	MinIoID = os.Getenv("MINIO_ID")
	MinIoSecretKey = os.Getenv("MINIO_SECRETKEY")
	MinIoEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinIoPort = os.Getenv("MINIO_PORT")
	MinIoBucket = os.Getenv("MINIO_BUCKET")
	MinIoSSL = miniosslBool

	DefaultImage = os.Getenv("DEFAULT_DEFAULT_IMAGE")
	AesCFB = os.Getenv("DEFAULT_AES_CFB_KEY")
	AesCBC = os.Getenv("DEFAULT_AES_CBC_KEY")
	AesCBCIV = os.Getenv("DEFAULT_AES_CBC_IV_KEY")

	log.Info().Msg("config initialization successfully")
}

var (
	AppPort      string
	AppStatus    string
	AppGracefull string

	PgHost string
	PgPort string
	PgUser string
	PgPass string
	PgName string
	PgSSL  string

	RedisHost string
	RedisPort string
	RedisDB   int
	RedisPass string

	MinIoID        string
	MinIoSecretKey string
	MinIoEndpoint  string
	MinIoPort      string
	MinIoBucket    string
	MinIoSSL       bool

	DefaultImage string
	AesCFB       string
	AesCBC       string
	AesCBCIV     string
)
