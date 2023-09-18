package unit

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/repository"
)

func GetProfileByIDREPO(t *testing.T) {
	id1 := uuid.NewV4().String()
	id2 := uuid.NewV4().String()
	id3 := uuid.NewV4().String()
	userId1 := "userid_1"
	userId2 := "userid_2"
	userId3 := "userid_3"
	columns := []string{"id", "user_id", "quotes", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)

	expectProfile := model.Profile{
		ProfileID: id1,
		UserID:    userId1,
		Quote:     sql.NullString{String: "semangat", Valid: true},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id1,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	prepareQuery := regexp.QuoteMeta("SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL")

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileByID(context.TODO(), id1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_GetProfileByID_DATANILL", func(t *testing.T) {
		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1).WillReturnRows(sqlmock.NewRows(columns))

		profile, err := profileRepo.GetProfileByID(context.TODO(), id1)
		assert.Error(t, err)
		assert.Equal(t, "", profile.ProfileID)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func GetProfileByUserIDREPO(t *testing.T) {
	id1 := uuid.NewV4().String()
	id2 := uuid.NewV4().String()
	id3 := uuid.NewV4().String()
	userId1 := "user id 1"
	userId2 := "user id 2"
	userId3 := "user id 3"
	columns := []string{"id", "user_id", "quotes", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)
	expectProfile := model.Profile{
		ProfileID: id1,
		UserID:    userId1,
		Quote:     sql.NullString{String: "semangat", Valid: true},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id1,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	query := regexp.QuoteMeta("SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL")

	t.Run("SUCCESS_GetProfileByUserID", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileByUserID(context.TODO(), userId1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_GetProfileByUserID_DATANILL", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(id1).
			WillReturnRows(sqlmock.NewRows(columns))

		profile, err := profileRepo.GetProfileByUserID(context.TODO(), id1)
		assert.Error(t, err)
		assert.Equal(t, "", profile.ProfileID)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func StoreProfileREPO(t *testing.T) {
	id1 := uuid.NewV4().String()
	userId1 := "user id 1"

	unix := time.Now().Unix()
	createProfile := &model.Profile{
		ProfileID: id1,
		UserID:    userId1,
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: unix,
		CreatedBy: id1,
		UpdatedAt: unix,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)

	query := regexp.QuoteMeta("SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1)")
	query2 := regexp.QuoteMeta("INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	t.Run("SUCCESS_StoreProfile", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		rows := sqlmock.NewRows([]string{"exists"})
		rows.AddRow(false)
		mocksql.ExpectQuery(query).WithArgs(createProfile.UserID).WillReturnRows(rows)
		mocksql.ExpectPrepare(query2)
		mocksql.ExpectExec(query2).WithArgs(
			createProfile.ProfileID,
			createProfile.UserID,
			createProfile.Quote,
			createProfile.CreatedAt,
			createProfile.CreatedBy,
			createProfile.UpdatedAt,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		profile, err := profileRepo.StoreProfile(context.TODO(), *createProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile.ProfileID, createProfile.ProfileID)
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_StoreProfile_DATAEXISTS", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		rows := sqlmock.NewRows([]string{"exists"})
		rows.AddRow(true)
		mocksql.ExpectQuery(query).WithArgs(createProfile.UserID).WillReturnRows(rows)
		mocksql.ExpectRollback()

		_, err = profileRepo.StoreProfile(context.TODO(), *createProfile)
		assert.Error(t, err)
		assert.Equal(t, err, model.ErrConflict)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func UpdateProfileREPO(t *testing.T) {
	id1 := uuid.NewV4().String()
	userId1 := "user id 1"

	unix := time.Now().Unix()
	updateProfile := model.Profile{
		ProfileID: id1,
		UserID:    userId1,
		Quote:     sql.NullString{String: "semagat", Valid: true},
		CreatedAt: unix,
		CreatedBy: id1,
		UpdatedAt: unix,
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)

	query := regexp.QuoteMeta("UPDATE dueit.m_profiles SET quotes = $1, updated_by = $2, updated_at = $3 WHERE user_id = $4 AND id = $5 AND deleted_at IS NULL")

	t.Run("SUCCESS_UpdateProfile", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			updateProfile.Quote,
			id1,
			unix,
			userId1,
			id1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		profile, err := profileRepo.UpdateProfile(context.TODO(), updateProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile, updateProfile)

		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
