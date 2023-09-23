package integration

import (
	"context"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/pkg/_usecase"
)

func ProfileUsecase(t *testing.T) {
	timeOut := 2 * time.Second
	profile := _usecase.NewProfileUsecaseImpl(ProfileRepo, UserRepo, timeOut)

	t.Run("SUCCESS_StoreProfile", func(t *testing.T) {
		profile, err := profile.StoreProfile(context.Background(), &dto.StoreProfileReq{
			UserID: userID_2,
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), &dto.GetProfileReq{
			UserID: userID_2,
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

}
