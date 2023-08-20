package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	domainerror "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-error"
	domainprofilecfg "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile-cfg"
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
	unix              = time.Now().Unix()
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
	profileConfig1 = domainprofilecfg.ProfileCfg{
		ID:          "profileCfgid1",
		ProfileId:   "profileid1",
		ConfigName:  "USER PERIOD",
		ConfigValue: marshal(profileCfgPeriod1),
		Status:      "on",
		CreatedAt:   unix,
		CreatedBy:   "profileCfgid1",
		UpdatedAt:   unix,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
	profileConfigUpdate1 = domainprofilecfg.ProfileCfg{
		ID:          "profileCfgid1",
		ProfileId:   "profileid1",
		ConfigName:  "USER PERIOD",
		ConfigValue: marshal(profileCfgPeriod1),
		Status:      "off",
		CreatedAt:   unix,
		CreatedBy:   "profileCfgid1",
		UpdatedAt:   unix,
		UpdatedBy:   sql.NullString{String: "profileid1", Valid: true},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
	profileConfig2 = domainprofilecfg.ProfileCfg{
		ID:          "profileCfgid2",
		ProfileId:   "profileid1",
		ConfigName:  "DAILY NOTIF",
		ConfigValue: marshal(profileCfgDay1),
		Status:      "on",
		CreatedAt:   unix,
		CreatedBy:   "profileid1",
		UpdatedAt:   unix,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
)

func TestProfileConfigRepo(t *testing.T) {
	t.Run("TestProfileRepo", ProfileRepo)
	fmt.Println("RUNNING TEST PROFILE CONFIG REPOSITORY")

	t.Run("SUCCESS_Store", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		err = profileCfgRepo.StoreProfileCfg(context.Background(), tx, profileConfig1)
		assert.NoError(t, err)
		err = profileCfgRepo.StoreProfileCfg(context.Background(), tx, profileConfig2)
		assert.NoError(t, err)
		tx.Commit()
	})

	t.Run("ERROR_Store", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		err = profileCfgRepo.StoreProfileCfg(context.Background(), tx, profileConfig1)
		assert.Error(t, err)
		assert.Equal(t, domainerror.ErrProfileConfigAlvailable, err)
		tx.Rollback()
	})

	t.Run("SUCCESS_Get-By-Id-Or-UserId", func(t *testing.T) {
		profileCfg, err := profileCfgRepo.GetProfileCfgById(context.Background(), db, profileConfig1.ID)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfg)
		assert.Equal(t, profileConfig1.ID, profileCfg.ID)
	})

	t.Run("ERROR_Get-By-Id-Or-UserId", func(t *testing.T) {
		profileCfg, err := profileCfgRepo.GetProfileCfgById(context.Background(), db, profileConfig1.ConfigName)
		assert.Error(t, err)
		assert.Nil(t, profileCfg)
		assert.Equal(t, sql.ErrNoRows, err)
	})

	t.Run("SUCCESS_Get-Scheduler", func(t *testing.T) {
		scheduler := domainprofilecfg.ProfileCfgScheduler{
			Day:  "monday",
			Time: "02:00",
		}

		profileCfgs, err := profileCfgRepo.GetProfileCfgByScheduler(context.Background(), db, scheduler)
		assert.NoError(t, err)
		assert.NotNil(t, profileCfgs)
		if len(*profileCfgs) < 1 {
			fmt.Println(len(*profileCfgs))
			os.Exit(1)
		}
	})

	t.Run("ERROR_Get-Scheduler", func(t *testing.T) {
		scheduler := domainprofilecfg.ProfileCfgScheduler{
			Day:  "saturday",
			Time: "02:00",
		}

		profileCfgs, err := profileCfgRepo.GetProfileCfgByScheduler(context.Background(), db, scheduler)
		assert.NoError(t, err)
		if len(*profileCfgs) >= 1 {
			fmt.Println(len(*profileCfgs))
			os.Exit(1)
		}
	})

	t.Run("SUCCESS_Update", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		err = profileCfgRepo.UpdateProfileCfg(context.Background(), tx, profileConfigUpdate1)
		assert.NoError(t, err)
		tx.Commit()
	})
}
