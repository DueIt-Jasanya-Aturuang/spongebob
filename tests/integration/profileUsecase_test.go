package integration

import (
	"context"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func ProfileUsecase(t *testing.T) {
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)
	timeOut := 2 * time.Second
	profile := usecase.NewProfileUsecaseImpl(profileRepo, timeOut)

	t.Run("SUCCESS_StoreProfile", func(t *testing.T) {
		profile, err := profile.StoreProfile(context.Background(), dto.StoreProfileReq{
			UserID: "userId2",
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), dto.GetProfileReq{
			UserID: "userId2",
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("CreateProfileCfgUSECASE", ProfileCfgUSECASE)
}
