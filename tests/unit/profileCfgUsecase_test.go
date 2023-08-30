package unit

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfileCfgUSECASE(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	profileCfgRepoMock := &mocks.FakeProfileCfgRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()

	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(profileRepoMock, profileCfgRepoMock, timeOutCtx)

	request := dto.CreateProfileCfgReq{
		ProfileID:   "profileid_1",
		ConfigValue: "19:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		ConfigName:   "DAILY_NOTIFY",
		Status:       "on",
		UserID:       "userid_1",
		Token:        "123",
		Value:        "19:00",
		IanaTimezone: "Asia/Jakarta",
	}

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userid_1",
		Quote:     sql.NullString{String: ""},
		CreatedAt: time.Now().Unix(),
		CreatedBy: "profileid_1",
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	t.Run("SUCCESS_CreateProfileCfg", func(t *testing.T) {
		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
		profileRepoMock.StartTxReturns(nil)

		profileRepoMock.GetProfileByID(ctx, request.ProfileID)
		profileRepoMock.GetProfileByIDReturns(profile, nil)

		profileCfgConv := helpers.CreateProfileCfgToModel(request, []byte("asd"))
		profileCfgRepoMock.StoreProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.StoreProfileCfgReturns(nil)

		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, request)
		profileRepoMock.EndTx(err)
		profileRepoMock.EndTxReturns(nil)

		profileRepoMock.CloseConn()

		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("ERROR_CreateProfileCfg_DATANIL", func(t *testing.T) {
		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
		profileRepoMock.StartTxReturns(nil)

		profileRepoMock.GetProfileByID(ctx, request.ProfileID)
		profileRepoMock.GetProfileByIDReturns(model.Profile{}, sql.ErrNoRows)

		profileCfgConv := helpers.CreateProfileCfgToModel(request, []byte("asd"))
		profileCfgRepoMock.StoreProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.StoreProfileCfgReturns(nil)

		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, request)
		profileRepoMock.EndTx(err)
		profileRepoMock.EndTxReturns(nil)

		profileRepoMock.CloseConn()
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, "", profileCfg.ProfileID)
	})
}

func TestGetProfileCfgByNameAndIDUSECASE(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	profileCfgRepoMock := &mocks.FakeProfileCfgRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()

	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(profileRepoMock, profileCfgRepoMock, timeOutCtx)

	configValue, _ := json.Marshal(map[string]any{
		"config_time_user":       "value",
		"config_timezone_user":   "ianaTimezone",
		"config_time_notify":     fmt.Sprintf("%02s:%02s", "19", "00"),
		"config_timezone_notify": "UTC",
		"days":                   []string{"monday"},
	})
	profileCfg := model.ProfileCfg{
		ID:          "cfgid_1",
		ProfileID:   "profileid_1",
		ConfigName:  "DAILY_NOTIF",
		ConfigValue: string(configValue),
		Status:      "on",
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   "profileid_1",
		UpdatedAt:   time.Now().Unix(),
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userId1",
		Quote:     sql.NullString{String: "Semangat"},
		CreatedAt: 0,
		CreatedBy: "profileid_1",
		UpdatedAt: 0,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	t.Run("SUCCESS_GetProfileCfgByNameAndID", func(t *testing.T) {
		req := dto.GetProfileCfgReq{
			UserID:     "userId1",
			ConfigName: "DAILY_NOTIFY",
			ProfileID:  "profileid_1",
		}

		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.GetProfileByID(ctx, req.ProfileID)
		profileRepoMock.GetProfileByIDReturns(profile, nil)

		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, req.ProfileID, req.ConfigName)
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(profileCfg, nil)

		profileCfgResp, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, req)
		profileRepoMock.CloseConn()
		assert.NoError(t, err)
		assert.NotNil(t, profileCfgResp)
	})

	t.Run("ERROR_GetProfileCfgByNameAndID_DATANIL", func(t *testing.T) {
		req := dto.GetProfileCfgReq{
			UserID:     "nil",
			ConfigName: "DAILY_NOTIFY",
			ProfileID:  "nil",
		}

		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.GetProfileByID(ctx, req.ProfileID)
		profileRepoMock.GetProfileByIDReturns(model.Profile{}, sql.ErrNoRows)

		profileCfgResp, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, req)
		profileRepoMock.CloseConn()
		assert.Error(t, err)
		assert.Equal(t, "", profileCfgResp.ProfileID)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}

func TestUpdateProfileCfgUSECASE(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	profileCfgRepoMock := &mocks.FakeProfileCfgRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()

	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(profileRepoMock, profileCfgRepoMock, timeOutCtx)

	request := dto.UpdateProfileCfgReq{
		ProfileID:   "profileid_1",
		ConfigValue: "19:00 Asia/Jakarta",
		Days: []string{
			"monday",
		},
		Status:       "on",
		Token:        "123",
		UserID:       "userId1",
		ConfigName:   "profileid_1",
		Value:        "19:00",
		IanaTimezone: "Asia/Jakarta",
	}

	configValue, _ := json.Marshal(map[string]any{
		"config_time_user":       "value",
		"config_timezone_user":   "ianaTimezone",
		"config_time_notify":     fmt.Sprintf("%02s:%02s", "19", "00"),
		"config_timezone_notify": "UTC",
		"days":                   []string{"monday"},
	})
	profileCfg := model.ProfileCfg{
		ID:          "cfgid_1",
		ProfileID:   "profileid_1",
		ConfigName:  "DAILY_NOTIF",
		ConfigValue: string(configValue),
		Status:      "on",
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   "profileid_1",
		UpdatedAt:   time.Now().Unix(),
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userId1",
		Quote:     sql.NullString{},
		CreatedAt: 0,
		CreatedBy: "",
		UpdatedAt: 0,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}
	t.Run("SUCCESS_UpdateProfileCfg", func(t *testing.T) {
		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
		profileRepoMock.StartTxReturns(nil)

		profileRepoMock.GetProfileByID(ctx, "profileid_1")
		profileRepoMock.GetProfileByIDReturns(profile, nil)

		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "nil", "nil")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(profileCfg, nil)
		assert.Equal(t, 1, profileCfgRepoMock.GetProfileCfgByNameAndIDCallCount())

		profileCfgConv := helpers.UpdateProfileCfgToModel(request, []byte("asd"), "DAILY_NOTIF", "cfgid_1")
		profileCfgRepoMock.UpdateProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.UpdateProfileCfgReturns(nil)
		assert.Equal(t, 1, profileCfgRepoMock.UpdateProfileCfgCallCount())

		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, request)
		profileRepoMock.EndTx(err)
		profileRepoMock.EndTxReturns(nil)
		profileRepoMock.CloseConn()
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("ERROR_UpdateProfileCfg_DATANIL", func(t *testing.T) {
		profileRepoMock.OpenConn(ctx)
		profileRepoMock.OpenConnReturns(nil)

		profileRepoMock.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
		profileRepoMock.StartTxReturns(nil)
		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "nil", "nil")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(model.ProfileCfg{}, sql.ErrNoRows)

		profileCfgConv := helpers.UpdateProfileCfgToModel(request, []byte("asd"), "DAILY_NOTIF", "cfgid_1")
		profileCfgRepoMock.UpdateProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.UpdateProfileCfgReturns(nil)

		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, request)
		profileRepoMock.EndTx(err)
		profileRepoMock.EndTxReturns(nil)
		profileRepoMock.CloseConn()
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, "", profileCfg.ProfileID)
	})
}
