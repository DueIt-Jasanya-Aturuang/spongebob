package unit

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/dtoconv"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfileCfgUSECASE(t *testing.T) {
	uow := &mocks.FakeUnitOfWork{}
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
		profileRepoMock.GetProfileByID(ctx, request.ProfileID)
		profileRepoMock.GetProfileByIDReturns(&profile, nil)

		uow.StartTx(ctx, &sql.TxOptions{})
		uow.StartTxReturns(nil)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgConv := dtoconv.CreateProfileCfgToModel(request, []byte("asd"))
		profileCfgRepoMock.StoreProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.StoreProfileCfgReturns(nil)

		uow.EndTx(nil)
		uow.EndTxReturns(nil)

		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, request)

		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("ERROR_CreateProfileCfg_DATANIL", func(t *testing.T) {
		profileRepoMock.GetProfileByID(ctx, request.ProfileID)
		profileRepoMock.GetProfileByIDReturns(nil, sql.ErrNoRows)

		uow.StartTx(ctx, &sql.TxOptions{})
		uow.StartTxReturns(nil)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgConv := dtoconv.CreateProfileCfgToModel(request, []byte("asd"))
		profileCfgRepoMock.StoreProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.StoreProfileCfgReturns(nil)

		uow.EndTx(nil)
		uow.EndTxReturns(nil)

		profileCfg, err := profileCfgUsecase.CreateProfileCfg(ctx, request)

		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Nil(t, profileCfg)
	})
}

func TestGetProfileCfgByNameAndIDUSECASE(t *testing.T) {
	uow := &mocks.FakeUnitOfWork{}
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

	t.Run("SUCCESS_GetProfileCfgByNameAndID", func(t *testing.T) {
		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "profileid_1", "DAILY_NOTIF")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(&profileCfg, nil)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgResp, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, "profileid_1", "DAILY_NOTIF")
		assert.NoError(t, err)
		assert.NotNil(t, profileCfgResp)
	})

	t.Run("ERROR_GetProfileCfgByNameAndID_DATANIL", func(t *testing.T) {
		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "profileid_1", "DAILY_NOTIF")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(nil, sql.ErrNoRows)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgResp, err := profileCfgUsecase.GetProfileCfgByNameAndID(ctx, "profileid_1", "DAILY_NOTIF")
		assert.Error(t, err)
		assert.Nil(t, profileCfgResp)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}

func TestUpdateProfileCfgUSECASE(t *testing.T) {
	uow := &mocks.FakeUnitOfWork{}
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
	profileCfg := &model.ProfileCfg{
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
	t.Run("SUCCESS_UpdateProfileCfg", func(t *testing.T) {
		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "nil", "nil")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(profileCfg, nil)
		assert.Equal(t, 1, profileCfgRepoMock.GetProfileCfgByNameAndIDCallCount())

		uow.StartTx(ctx, &sql.TxOptions{})
		uow.StartTxReturns(nil)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgConv := dtoconv.UpdateProfileCfgToModel(request, []byte("asd"), "DAILY_NOTIF", "cfgid_1")
		profileCfgRepoMock.UpdateProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.UpdateProfileCfgReturns(nil)
		assert.Equal(t, 1, profileCfgRepoMock.UpdateProfileCfgCallCount())

		uow.EndTx(nil)
		uow.EndTxReturns(nil)
		assert.Equal(t, 1, uow.EndTxCallCount())

		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, request, "cfgid_1", "profileid_1")

		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
	})

	t.Run("ERROR_UpdateProfileCfg_DATANIL", func(t *testing.T) {
		profileCfgRepoMock.GetProfileCfgByNameAndID(ctx, "nil", "nil")
		profileCfgRepoMock.GetProfileCfgByNameAndIDReturns(nil, sql.ErrNoRows)

		uow.StartTx(ctx, &sql.TxOptions{})
		uow.StartTxReturns(nil)

		profileCfgRepoMock.UoW()
		profileCfgRepoMock.UoWReturns(uow)

		profileCfgConv := dtoconv.UpdateProfileCfgToModel(request, []byte("asd"), "DAILY_NOTIF", "cfgid_1")
		profileCfgRepoMock.UpdateProfileCfg(ctx, profileCfgConv)
		profileCfgRepoMock.UpdateProfileCfgReturns(nil)

		uow.EndTx(nil)
		uow.EndTxReturns(nil)

		profileCfg, err := profileCfgUsecase.UpdateProfileCfg(ctx, request, "cfgid_1", "profileid_1")
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Nil(t, profileCfg)
	})
}
