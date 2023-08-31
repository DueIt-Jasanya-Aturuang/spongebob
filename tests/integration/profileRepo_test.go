package integration

import (
	"context"
	"database/sql"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ProfileREPO(t *testing.T) {
	t.Run("SuccessStoreProfile", SuccessStoreProfile)
	t.Run("SuccessGetProfileByID", SuccessGetProfileByID)
	t.Run("ErrorGetProfileByID", ErrorGetProfileByID)
	t.Run("SuccessGetProfileByUserID", SuccessGetProfileByUserID)
	t.Run("ErrorGetProfileByUserID", ErrorGetProfileByUserID)
	t.Run("SuccessUpdateProfile", SuccessUpdateProfile)
}

func SuccessStoreProfile(t *testing.T) {
	err := ProfileRepo.OpenConn(context.Background())
	assert.NoError(t, err)
	err = ProfileRepo.StartTx(context.Background(), &sql.TxOptions{ReadOnly: false})
	assert.NoError(t, err)
	profile, err := ProfileRepo.StoreProfile(context.Background(), profile_1)
	assert.NoError(t, err)
	assert.Equal(t, profile_1, profile)
	ProfileRepo.EndTx(err)
	ProfileRepo.CloseConn()
}

func SuccessGetProfileByID(t *testing.T) {
	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		err := ProfileRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		profile, err := ProfileRepo.GetProfileByID(context.TODO(), profile_1.ProfileID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile_1, profile)
		ProfileRepo.CloseConn()
	})
}

func ErrorGetProfileByID(t *testing.T) {
	t.Run("ERROR_GetProfileByID_NOROW", func(t *testing.T) {
		err := ProfileRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		profile, err := ProfileRepo.GetProfileByID(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Equal(t, "", profile.UserID)
		assert.Equal(t, err, sql.ErrNoRows)
		ProfileRepo.CloseConn()
	})
}

func SuccessGetProfileByUserID(t *testing.T) {
	t.Run("SUCCESS_GetProfileByUserID", func(t *testing.T) {
		err := ProfileRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		profile, err := ProfileRepo.GetProfileByUserID(context.TODO(), profile_1.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile_1, profile)
		ProfileRepo.CloseConn()
	})
}

func ErrorGetProfileByUserID(t *testing.T) {
	t.Run("ERROR_GetProfileByUserID_NOROW", func(t *testing.T) {
		err := ProfileRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		profile, err := ProfileRepo.GetProfileByUserID(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Equal(t, "", profile.UserID)
		assert.Equal(t, err, sql.ErrNoRows)
		ProfileRepo.CloseConn()
	})
}

func SuccessUpdateProfile(t *testing.T) {
	t.Run("SUCCESS_UpdateProfile", func(t *testing.T) {
		updateProfile := model.Profile{
			ProfileID: profileID_1,
			UserID:    userID_1,
			Quote:     sql.NullString{String: "semagat", Valid: true},
			Profesi:   sql.NullString{String: "semagat", Valid: true},
			CreatedAt: time.Now().Unix(),
			CreatedBy: "id1",
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{String: profileID_1, Valid: true},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		}
		err := ProfileRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		err = ProfileRepo.StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		profile, err := ProfileRepo.UpdateProfile(context.TODO(), updateProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.NotEqual(t, profile_1, profile)
		assert.Equal(t, updateProfile, profile)
		ProfileRepo.EndTx(err)
		ProfileRepo.CloseConn()
	})
}
