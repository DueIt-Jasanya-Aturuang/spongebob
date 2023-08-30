package integration

import (
	"context"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ProfileCfgUSECASE(t *testing.T) {
	timeOut := 2 * time.Second
	ctx := context.Background()

	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(ProfileRepo, ProfileCfgRepo, timeOut)
	req := dto.CreateProfileCfgReq{
		ProfileID:   profileID_1,
		ConfigValue: "19:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		ConfigName:   "DAILY_NOTIFY",
		Status:       "on",
		Token:        "123",
		UserID:       userID_1,
		Value:        "19:00",
		IanaTimezone: "Asia/Jakarta",
	}

	reqUpdate := dto.UpdateProfileCfgReq{
		ProfileID:   profileID_1,
		ConfigValue: "20:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		Status:       "on",
		Token:        "123",
		UserID:       userID_1,
		ConfigName:   "DAILY_NOTIFY",
		Value:        "20:00",
		IanaTimezone: "Asia/Jakarta",
	}

	var profileCfgResp dto.ProfileCfgResp
	t.Run("SUCCESS_CreateProfileCfgUSECASE", func(t *testing.T) {
		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, req)
		t.Log(profileCfg)
		profileCfgResp = profileCfg
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("SUCCESS_GetProfileCfgByNameAndIDUSECASE", func(t *testing.T) {
		req := dto.GetProfileCfgReq{
			UserID:     userID_1,
			ConfigName: "DAILY_NOTIFY",
			ProfileID:  profileID_1,
		}
		profileCfg, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.Equal(t, profileCfgResp, profileCfg)
	})

	t.Run("SUCCESS_UpdateProfileCfgUSECASE", func(t *testing.T) {
		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, reqUpdate)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.NotEqual(t, profileCfgResp, profileCfg)
	})
}
