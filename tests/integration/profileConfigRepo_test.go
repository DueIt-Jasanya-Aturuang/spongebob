package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func marshal(data any) string {
	byte, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Msg("cannot marshal")
		os.Exit(1)
	}
	return string(byte)
}

var (
	unixProfileCfg    = time.Now().Unix()
	profileCfgPeriod1 = map[string]any{
		"config_date": 29,
		"token":       "123",
	}
	profileCfgDay1 = map[string]any{
		"config_time_user":       "19:00",
		"config_timezone_user":   "Asia/Jakarta",
		"config_time_notify":     fmt.Sprintf("%02d:%02d", 0o2, 0o0),
		"config_timezone_notify": "UTC",
		"days":                   []string{"monday", "sunday"},
		"token":                  "1234",
	}
	profileConfig1 = model.ProfileCfg{
		ID:          "profileCfgid1",
		ProfileID:   "profileid1",
		ConfigName:  "USER PERIOD",
		ConfigValue: marshal(profileCfgPeriod1),
		Status:      "on",
		CreatedAt:   unixProfileCfg,
		CreatedBy:   "profileCfgid1",
		UpdatedAt:   unixProfileCfg,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
	profileConfigUpdate1 = model.ProfileCfg{
		ID:          "profileCfgid1",
		ProfileID:   "profileid1",
		ConfigName:  "USER PERIOD",
		ConfigValue: marshal(profileCfgPeriod1),
		Status:      "off",
		CreatedAt:   unixProfileCfg,
		CreatedBy:   "profileCfgid1",
		UpdatedAt:   unixProfileCfg,
		UpdatedBy:   sql.NullString{String: "profileid1", Valid: true},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
	profileConfig2 = model.ProfileCfg{
		ID:          "profileCfgid2",
		ProfileID:   "profileid1",
		ConfigName:  "DAILY NOTIF",
		ConfigValue: marshal(profileCfgDay1),
		Status:      "on",
		CreatedAt:   unixProfileCfg,
		CreatedBy:   "profileid1",
		UpdatedAt:   unixProfileCfg,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
)

func TestProfileConfigREPO(t *testing.T) {
	profileCfgRepo := repository.NewProfileCfgRepoImpl(db)
	t.Run("TestProfileRepo", ProfileREPO)
	fmt.Println("RUNNING TEST PROFILE CONFIG REPOSITORY")

	t.Run("SUCCESS_StoreProfileCfg", func(t *testing.T) {
		err := profileCfgRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		err = profileCfgRepo.StoreProfileCfg(context.Background(), profileConfig1)
		assert.NoError(t, err)
		err = profileCfgRepo.Commit()
		assert.NoError(t, err)

		err = profileCfgRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		err = profileCfgRepo.StoreProfileCfg(context.Background(), profileConfig2)
		assert.NoError(t, err)
		err = profileCfgRepo.Commit()
		assert.NoError(t, err)
	})

	t.Run("ERROR_StoreProfileCfg_PROFILECFGEXISTS", func(t *testing.T) {
		err := profileCfgRepo.BeginTx(context.TODO(), &sql.TxOptions{})
		assert.NoError(t, err)
		err = profileCfgRepo.StoreProfileCfg(context.Background(), profileConfig1)
		assert.Error(t, err)
		assert.Equal(t, exception.Err400ProfileConfigAlvailable, err)

		err = profileCfgRepo.Rollback()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_GetProfileCfgByID", func(t *testing.T) {
		profileCfg, err := profileCfgRepo.GetProfileCfgByID(context.Background(), profileConfig1.ID)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.Equal(t, profileConfig1.ID, profileCfg.ID)
	})

	t.Run("ERROR_GetProfileCfgByID_NOROW", func(t *testing.T) {
		profileCfg, err := profileCfgRepo.GetProfileCfgByID(context.Background(), profileConfig1.ConfigName)
		assert.Error(t, err)
		assert.Nil(t, profileCfg)
		assert.Equal(t, sql.ErrNoRows, err)
	})

	t.Run("SUCCESS_GetProfileCfgByScheduler", func(t *testing.T) {
		scheduler := dto.ProfileCfgScheduler{
			Day:  "monday",
			Time: "02:00",
		}

		profileCfgs, err := profileCfgRepo.GetProfileCfgByScheduler(context.Background(), scheduler)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfgs)
		if len(*profileCfgs) < 1 {
			fmt.Println(len(*profileCfgs))
			os.Exit(1)
		}
	})

	t.Run("ERROR_GetProfileCfgByScheduler_NOROWS", func(t *testing.T) {
		scheduler := dto.ProfileCfgScheduler{
			Day:  "saturday",
			Time: "02:00",
		}

		profileCfgs, err := profileCfgRepo.GetProfileCfgByScheduler(context.Background(), scheduler)
		assert.NoError(t, err)
		if len(*profileCfgs) >= 1 {
			fmt.Println(len(*profileCfgs))
			os.Exit(1)
		}
	})

	t.Run("SUCCESS_UpdateProfileCfg", func(t *testing.T) {
		err := profileCfgRepo.BeginTx(context.TODO(), &sql.TxOptions{})
		assert.NoError(t, err)

		err = profileCfgRepo.UpdateProfileCfg(context.Background(), profileConfigUpdate1)
		assert.NoError(t, err)

		err = profileCfgRepo.Commit()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_GetProfileCfgByID_AFTERUPDATE", func(t *testing.T) {
		profileCfg, err := profileCfgRepo.GetProfileCfgByID(context.Background(), profileConfigUpdate1.ID)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		t.Log(profileCfg)
	})
}
