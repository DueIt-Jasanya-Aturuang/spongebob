package integration

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/tests/integration/utils"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

var (
	db          *sql.DB
	minioClient *minio.Client
)

func TestMain(t *testing.M) {
	pool := utils.InitDocker()

	resource, dbPg, url := utils.PostgresStart(pool)
	db = dbPg
	if db == nil {
		panic("db nil")
	}

	endpoint := utils.MinioStart(utils.InitDocker())
	minioConn, err := config.NewMinioConn(endpoint, "MYACCESSKEY", "MYSECRETKEY", false)
	if err != nil {
		panic(err)
	}
	minioClient = minioConn
	utils.StartMigration(url, db)
	// Run tests
	code := t.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Err(err).Msgf("Could not purge resource: %s", err)
		os.Exit(1)
	}
	os.Exit(code)
}
