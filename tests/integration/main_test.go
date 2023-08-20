package integration

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/repositories"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/tests/integration/utils"
	"github.com/rs/zerolog/log"
)

var (
	db             *sql.DB
	profileRepo    = repositories.NewProfileRepoImpl()
	userRepo       = repositories.NewUserRepoImpl()
	profileCfgRepo = repositories.NewProfileCfgRepoImpl()
)

func TestMain(t *testing.M) {
	// log.Logger = log.Output(zerolog.Nop())
	pool := utils.InitDocker()

	// pulls an image, creates a container based on it and runs it
	resource, dbPg, url := utils.PostgresStart(pool)

	db = dbPg
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
