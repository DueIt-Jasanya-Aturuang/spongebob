package integration

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/stretchr/testify/assert"
)

func ProfileREPO(t *testing.T) {
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)
	fmt.Println("RUNNING TEST PROFILE REPOSITORY")
	unix := time.Now().Unix()
	dataProfile := model.Profile{
		ProfileID: "profileid1",
		UserID:    "userId1",
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: unix,
		CreatedBy: "profileid1",
		UpdatedAt: unix,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	t.Run("SUCCESS_StoreProfile", func(t *testing.T) {
		err := profileRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		profile, err := profileRepo.StoreProfile(context.Background(), dataProfile)
		assert.NoError(t, err)
		assert.Equal(t, dataProfile, profile)
		profileRepo.UoW().EndTx(err)
	})

	t.Run("ERROR_StoreProfile_PROFILEEXISTS", func(t *testing.T) {
		err := profileRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		profile, err := profileRepo.StoreProfile(context.Background(), dataProfile)
		assert.Error(t, err)
		assert.NotEqual(t, dataProfile, profile)
		assert.Equal(t, exception.Err400ProfileAvailable, err)
		profileRepo.UoW().EndTx(err)
	})

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByID(context.TODO(), dataProfile.ProfileID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("ERROR_GetProfileByID_NOROW", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByID(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("SUCCESS_GetProfileByUserID", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserID(context.TODO(), dataProfile.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, &dataProfile, profile)
	})

	t.Run("ERROR_GetProfileByUserID_NOROW", func(t *testing.T) {
		profile, err := profileRepo.GetProfileByUserID(context.TODO(), "nil")
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("SUCCESS_UpdateProfile", func(t *testing.T) {
		updateProfile := model.Profile{
			ProfileID: "id1",
			UserID:    "userId1",
			Quote:     sql.NullString{String: "semagat", Valid: true},
			CreatedAt: unix,
			CreatedBy: "id1",
			UpdatedAt: unix,
			UpdatedBy: sql.NullString{String: "id1", Valid: true},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		}
		err := profileRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		profile, err := profileRepo.UpdateProfile(context.TODO(), updateProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.NotEqual(t, &dataProfile, profile)
		assert.Equal(t, &updateProfile, profile)
		profileRepo.UoW().EndTx(err)
	})

	t.Run("ProfileUSECASE", ProfileUsecase)
}
