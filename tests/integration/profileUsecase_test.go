package integration

import (
	"context"
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

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), "profileid1")
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("SUCCESS_GetProfileByUserID", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), "userId1")
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("SUCCESS_GetProfileByIDWithStore", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), "profileidnottfound")
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})
}
