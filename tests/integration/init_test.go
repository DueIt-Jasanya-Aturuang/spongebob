package integration

import (
	"database/sql"
	"testing"
	"time"

	repository2 "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/_repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/minio/minio-go/v7"
	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/repository"
)

var (
	db             *sql.DB
	minioClient    *minio.Client
	ProfileRepo    repository2.ProfileRepo
	UserRepo       repository2.UserRepo
	ProfileCfgRepo repository2.ProfileCfgRepo

	profileID_1 = uuid.NewV4().String()
	userID_1    = uuid.NewV4().String()
	userID_2    = uuid.NewV4().String()
	profile_1   = model.Profile{
		ProfileID: profileID_1,
		UserID:    userID_1,
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: time.Now().Unix(),
		CreatedBy: profileID_1,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}
)

func TestInit(t *testing.T) {
	UOW := repository.NewUnitOfWorkImpl(db)
	ProfileRepo = repository.NewProfileRepoImpl(UOW)
	UserRepo = repository.NewUserRepoImpl(UOW)
	ProfileCfgRepo = repository.NewProfileCfgRepoImpl(UOW)

	t.Run("PROFILE REPO", ProfileREPO)
	t.Run("USER REPO", UserREPO)
	t.Run("PROFILE CONFIG REPO", ProfileConfigREPO)
	t.Run("MINIO", Minio)
	t.Run("AccountUpdateUSECASE", AccountUpdateUSECASE)
	t.Run("PROFILE USECASE", ProfileUsecase)
	t.Run("ProfileCfg USECASE", ProfileCfgUSECASE)
}
