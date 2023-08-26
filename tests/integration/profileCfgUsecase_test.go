package integration

import (
	"context"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ProfileCfgUSECASE(t *testing.T) {
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo, profileCfgRepo := repository.NewProfileRepoImpl(uow), repository.NewProfileCfgRepoImpl(uow)
	timeOut := 2 * time.Second
	ctx := context.Background()

	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(profileRepo, profileCfgRepo, timeOut)
	req := dto.CreateProfileCfgReq{
		ProfileID:   "profileid1",
		ConfigValue: "19:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		ConfigName:   "DAILY_NOTIFY",
		Status:       "on",
		Token:        "123",
		Value:        "19:00",
		IanaTimezone: "Asia/Jakarta",
	}

	reqUpdate := dto.UpdateProfileCfgReq{
		ProfileID:   "profileid1",
		ConfigValue: "20:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		Status:       "on",
		Token:        "123",
		Value:        "20:00",
		IanaTimezone: "Asia/Jakarta",
	}

	var profileCfgResp *dto.ProfileCfgResp
	t.Run("SUCCESS_CreateProfileCfgUSECASE", func(t *testing.T) {
		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, req)
		t.Log(profileCfg)
		profileCfgResp = profileCfg
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("SUCCESS_GetProfileCfgByNameAndIDUSECASE", func(t *testing.T) {
		profileCfg, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, profileCfgResp.ID, "profileid1", "DAILY_NOTIFY")
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.Equal(t, profileCfgResp, profileCfg)
	})

	t.Run("SUCCESS_UpdateProfileCfgUSECASE", func(t *testing.T) {
		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, reqUpdate, profileCfgResp.ID, "DAILY_NOTIFY")
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.NotEqual(t, profileCfgResp, profileCfg)
	})
}
