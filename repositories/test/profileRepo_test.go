package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestProfileRepo(t *testing.T) {
	unix := time.Now().Unix()
	dataProfile := domain.Profile{
		ProfileId: "id1",
		UserId:    "userId1",
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: unix,
		CreatedBy: "id1",
		UpdatedAt: unix,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	t.Run("success store data", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		profile, err := profileRepo.StoreProfile(context.Background(), tx, dataProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		tx.Commit()
	})

	t.Run("success get data by id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileById(context.TODO(), db, dataProfile.ProfileId)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("error get data by id nil", func(t *testing.T) {
		profile, err := profileRepo.GetProfileById(context.TODO(), db, "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})
	t.Run("success get data by user id", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserId(context.TODO(), db, dataProfile.UserId)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("error get data by user id nil", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserId(context.TODO(), db, "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("success update data", func(t *testing.T) {
		updateProfile := domain.Profile{
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
