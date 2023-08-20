package withoutintegration

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repositories"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var (
	id1         = uuid.NewV4().String()
	id2         = uuid.NewV4().String()
	id3         = uuid.NewV4().String()
	userId1     = "user id 1"
	userId2     = "user id 2"
	userId3     = "user id 3"
	isNull      = "NULL"
	columns     = []string{"id", "user_id", "quotes", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	profileRepo = repositories.NewProfileRepoImpl()
)

func TestGetProfileById(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectProfile := &domain.Profile{
		ProfileId: id1,
		UserId:    userId1,
		Quote:     sql.NullString{String: "semangat", Valid: true},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id1,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	prepareQuery := regexp.QuoteMeta("SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS $2")

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1, isNull).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileById(context.TODO(), db, id1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
	t.Run("error deleted not nil", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, time.Now().Unix(), nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1, isNull).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileById(context.TODO(), db, id1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
	t.Run("error data nil", func(t *testing.T) {
		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1, isNull).WillReturnRows(sqlmock.NewRows(columns))

		profile, err := profileRepo.GetProfileById(context.TODO(), db, id1)
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestGetProfileByUserId(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	expectProfile := &domain.Profile{
		ProfileId: id1,
		UserId:    userId1,
		Quote:     sql.NullString{String: "semangat", Valid: true},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id1,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	prepareQuery := regexp.QuoteMeta("SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS $2")

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(userId1, isNull).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileByUserId(context.TODO(), db, userId1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
	t.Run("error deleted not nil", func(t *testing.T) {
		rows := sqlmock.NewRows(columns)
		rows.AddRow(id1, userId1, "semangat", time.Now().Unix(), id1, time.Now().Unix(), nil, time.Now().Unix(), nil)
		rows.AddRow(id2, userId2, "semangat", time.Now().Unix(), id2, time.Now().Unix(), nil, nil, nil)
		rows.AddRow(id3, userId3, "semangat", time.Now().Unix(), id3, time.Now().Unix(), nil, nil, nil)

		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1, isNull).WillReturnRows(rows)

		profile, err := profileRepo.GetProfileByUserId(context.TODO(), db, id1)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
	t.Run("error data nil", func(t *testing.T) {
		mocksql.ExpectPrepare(prepareQuery)
		mocksql.ExpectQuery(prepareQuery).WithArgs(id1, isNull).
			WillReturnRows(sqlmock.NewRows(columns))

		profile, err := profileRepo.GetProfileByUserId(context.TODO(), db, id1)
		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.NotEqual(t, expectProfile, profile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestStoreProfile(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	unix := time.Now().Unix()
	createProfile := &domain.Profile{
		ProfileId: id1,
		UserId:    userId1,
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

	query := regexp.QuoteMeta("INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")

	t.Run("success", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			createProfile.ProfileId,
			createProfile.UserId,
			createProfile.Quote,
			createProfile.CreatedAt,
			createProfile.CreatedBy,
			createProfile.UpdatedAt,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		profile, err := profileRepo.StoreProfile(context.TODO(), tx, *createProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile.ProfileId, createProfile.ProfileId)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	unix := time.Now().Unix()
	updateProfile := &domain.Profile{
		ProfileId: id1,
		UserId:    userId1,
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

	query := regexp.QuoteMeta("UPDATE dueit.m_profiles SET quotes = $1, updated_by = $2, updated_at = $3 WHERE user_id = $4 AND id = $5 AND deleted_at IS NULL")

	t.Run("success", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			updateProfile.Quote,
			id1,
			unix,
			userId1,
			id1,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		profile, err := profileRepo.UpdateProfile(context.TODO(), tx, *updateProfile)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, profile, updateProfile)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
