package integration

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repositories"
	"github.com/stretchr/testify/assert"
)

func ProfileRepo(t *testing.T) {
	profileRepo := repositories.NewProfileRepoImpl(db)
	fmt.Println("RUNNING TEST PROFILE REPOSITORY")
	unix := time.Now().Unix()
	dataProfile := model.Profile{
		ProfileId: "profileid1",
		UserId:    "userId1",
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: unix,
		CreatedBy: "profileid1",
		UpdatedAt: unix,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	t.Run("SUCCESS_Store", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		profile, err := profileRepo.StoreProfile(context.Background(), tx, dataProfile)
		assert.NoError(t, err)
		assert.Equal(t, dataProfile, profile)
		tx.Commit()
	})

	t.Run("ERROR_Store", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		profile, err := profileRepo.StoreProfile(context.Background(), tx, dataProfile)
		assert.Error(t, err)
		assert.NotEqual(t, dataProfile, profile)
		assert.Equal(t, exceptions.ErrProfileAlvailable, err)
		tx.Rollback()
	})

	t.Run("SUCCESS_Get-By-Id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileById(context.TODO(), dataProfile.ProfileId)
		t.Log(err)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("ERROR_Get-By-Id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileById(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})
	t.Run("SUCCESS_Get-By-User-Id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserId(context.TODO(), dataProfile.UserId)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("ERROR_Get-By-User-Id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserId(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("SUCCESS_Update", func(t *testing.T) {
		updateProfile := model.Profile{
			ProfileId: "id1",
			UserId:    "userId1",
			Quote:     sql.NullString{String: "semagat", Valid: true},
			CreatedAt: unix,
			CreatedBy: "id1",
			UpdatedAt: unix,
			UpdatedBy: sql.NullString{String: "id1", Valid: true},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		}
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		profile, err := profileRepo.UpdateProfile(context.TODO(), tx, updateProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.NotEqual(t, &dataProfile, profile)
		assert.Equal(t, &updateProfile, profile)
		tx.Commit()
	})
}
