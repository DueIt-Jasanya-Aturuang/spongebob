package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func marshal(data any) string {
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Msg("cannot marshal")
		os.Exit(1)
	}
	return string(byteData)
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
		ProfileID:   profileID_1,
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
		ProfileID:   profileID_1,
		ConfigName:  "USER PERIOD",
		ConfigValue: marshal(profileCfgPeriod1),
		Status:      "off",
		CreatedAt:   unixProfileCfg,
		CreatedBy:   "profileCfgid1",
		UpdatedAt:   unixProfileCfg,
		UpdatedBy:   sql.NullString{String: profileID_1, Valid: true},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
	profileConfig2 = model.ProfileCfg{
		ID:          "profileCfgid2",
		ProfileID:   profileID_1,
		ConfigName:  "DAILY NOTIF",
		ConfigValue: marshal(profileCfgDay1),
		Status:      "on",
		CreatedAt:   unixProfileCfg,
		CreatedBy:   profileID_1,
		UpdatedAt:   unixProfileCfg,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
)

func ProfileConfigREPO(t *testing.T) {

	t.Run("SUCCESS_StoreProfileCfg", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		err = ProfileCfgRepo.StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		err = ProfileCfgRepo.StoreProfileCfg(context.Background(), profileConfig1)
		assert.NoError(t, err)
		ProfileCfgRepo.EndTx(err)

		err = ProfileCfgRepo.StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		err = ProfileCfgRepo.StoreProfileCfg(context.Background(), profileConfig2)
		assert.NoError(t, err)
		ProfileCfgRepo.EndTx(err)
		ProfileCfgRepo.CloseConn()
	})

	t.Run("ERROR_StoreProfileCfg_PROFILECFGEXISTS", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		err = ProfileCfgRepo.StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		err = ProfileCfgRepo.StoreProfileCfg(context.Background(), profileConfig1)
		assert.Error(t, err)
		assert.Equal(t, err, model.ErrConflict)
		ProfileCfgRepo.EndTx(errors.New("PROFILECFGEXISTS"))
		ProfileCfgRepo.CloseConn()
	})

	t.Run("SUCCESS_GetProfileCfgByID", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		profileCfg, err := ProfileCfgRepo.GetProfileCfgByNameAndID(context.Background(), profileConfig1.ProfileID, profileConfig1.ConfigName)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.Equal(t, profileConfig1.ID, profileCfg.ID)
		ProfileCfgRepo.CloseConn()
	})

	t.Run("ERROR_GetProfileCfgByID_NOROW", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		profileCfg, err := ProfileCfgRepo.GetProfileCfgByNameAndID(context.Background(), "nil", "nil")
		assert.Error(t, err)
		assert.Equal(t, "", profileCfg.ID)
		assert.Equal(t, sql.ErrNoRows, err)
		ProfileCfgRepo.CloseConn()
	})

	t.Run("SUCCESS_GetProfileCfgByScheduler", func(t *testing.T) {
		scheduler := dto.ProfileCfgSche{
			Day:  "monday",
			Time: "02:00",
		}
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		profileCfgs, err := ProfileCfgRepo.GetProfileCfgByScheduler(context.Background(), scheduler)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfgs)
		if len(profileCfgs) < 1 {
			fmt.Println(len(profileCfgs))
			os.Exit(1)
		}
		ProfileCfgRepo.CloseConn()
	})

	t.Run("ERROR_GetProfileCfgByScheduler_NOROWS", func(t *testing.T) {
		scheduler := dto.ProfileCfgSche{
			Day:  "saturday",
			Time: "02:00",
		}
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		profileCfgs, err := ProfileCfgRepo.GetProfileCfgByScheduler(context.Background(), scheduler)
		assert.NoError(t, err)
		if len(profileCfgs) >= 1 {
			fmt.Println(len(profileCfgs))
			os.Exit(1)
		}
		ProfileCfgRepo.CloseConn()
	})

	t.Run("SUCCESS_UpdateProfileCfg", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		err = ProfileCfgRepo.StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		err = ProfileCfgRepo.UpdateProfileCfg(context.Background(), profileConfigUpdate1)
		assert.NoError(t, err)

		ProfileCfgRepo.EndTx(err)
		ProfileCfgRepo.CloseConn()
	})

	t.Run("SUCCESS_GetProfileCfgByID_AFTERUPDATE", func(t *testing.T) {
		err := ProfileCfgRepo.OpenConn(context.Background())
		assert.NoError(t, err)
		profileCfg, err := ProfileCfgRepo.GetProfileCfgByNameAndID(context.Background(), profileConfigUpdate1.ProfileID, profileConfigUpdate1.ConfigName)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		ProfileCfgRepo.CloseConn()
	})
}
