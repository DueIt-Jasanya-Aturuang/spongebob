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
	"github.com/stretchr/testify/assert"
)

var (
	gender      = "male"
	userColumns = []string{"id", "fullname", "gender", "image", "email", "username", "password", "phone_number", "email_verified_at", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	userRepo    = repositories.NewUserRepoImpl()
	image       = "default-male.png"
	email       = "ibanrama29@gmail.com"
)

func TestGetUserById(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	unix := time.Now().Unix()
	expectUser := domain.User{
		ID:              userId1,
		FullName:        "rama",
		Gender:          gender,
		Image:           image,
		Username:        "ibanrmaa",
		Email:           email,
		Password:        "123456",
		PhoneNumber:     sql.NullString{String: "12345678", Valid: true},
		EmailVerifiedAt: true,
		CreatedAt:       unix,
		CreatedBy:       userId1,
		UpdatedAt:       unix,
		UpdatedBy:       sql.NullString{},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := regexp.QuoteMeta("SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM auth.m_users WHERE id = $1 AND deleted_at IS $2")
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", email, "123456", "12345678", true, unix, userId1, unix, nil, nil, nil)
		rows.AddRow(userId2, "2rama", gender, image, "2ibanrmaa", "2ibanrmaa29@gmail.com", "1234567", "123456789", true, unix, userId2, unix, nil, nil, nil)

		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1, "NULL").WillReturnRows(rows)

		user, err := userRepo.GetUserById(context.TODO(), db, userId1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &expectUser, user)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error deleted not nul", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", email, "123456", "12345678", true, unix, userId1, unix, nil, unix, "admin")
		rows.AddRow(userId2, "2rama", gender, image, "2ibanrmaa", "2ibanrmaa29@gmail.com", "1234567", "123456789", true, unix, userId2, unix, nil, nil, nil)

		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1, "NULL").WillReturnRows(rows)

		user, err := userRepo.GetUserById(context.TODO(), db, userId1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.NotEqual(t, &expectUser, user)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1, isNull).WillReturnRows(sqlmock.NewRows(userColumns))

		user, err := userRepo.GetUserById(context.TODO(), db, userId1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.NotEqual(t, &expectUser, user)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	unix := time.Now().Unix()
	updateUser := domain.User{
		ID:              userId1,
		FullName:        "rama",
		Gender:          gender,
		Image:           image,
		Username:        "ibanrmaa",
		Email:           email,
		Password:        "123456",
		PhoneNumber:     sql.NullString{String: "12345678", Valid: true},
		EmailVerifiedAt: true,
		CreatedAt:       unix,
		CreatedBy:       userId1,
		UpdatedAt:       unix,
		UpdatedBy:       sql.NullString{String: userId1, Valid: true},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := regexp.QuoteMeta("UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, updated_at = $5, updated_by = $6 WHERE id = $7 AND deleted_at IS NULL")
	t.Run("success", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectExec(query).WithArgs(
			"rama",
			gender,
			image,
			"12345678",
			unix,
			userId1,
			userId1,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUser(context.TODO(), tx, updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, &updateUser)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateUserUsername(t *testing.T) {
	log.Logger = log.Output(zerolog.Nop())
	unix := time.Now().Unix()
	updateUser := domain.User{
		ID:              userId1,
		FullName:        "rama",
		Gender:          gender,
		Image:           image,
		Username:        "ibanrmaa",
		Email:           email,
		Password:        "123456",
		PhoneNumber:     sql.NullString{String: "12345678", Valid: true},
		EmailVerifiedAt: true,
		CreatedAt:       unix,
		CreatedBy:       userId1,
		UpdatedAt:       unix,
		UpdatedBy:       sql.NullString{String: userId1, Valid: true},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}

	db, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := regexp.QuoteMeta("SELECT username, id FROM auth.m_users WHERE username=$1 AND id<>$2 AND deleted_at IS NULL")
	query1 := regexp.QuoteMeta("UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 WHERE id = $4 AND deleted_at IS NULL")
	t.Run("success", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			updateUser.Username,
			updateUser.ID,
		).WillReturnRows(sqlmock.NewRows(userColumns))

		mocksql.ExpectPrepare(query1)
		mocksql.ExpectExec(query1).WithArgs(
			updateUser.Username,
			unix,
			updateUser.ID,
			updateUser.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUsername(context.TODO(), tx, updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, &updateUser)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("username alvailable", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", "ibanrama29@gmail.com", "123456", "12345678", true, unix, userId1, unix, nil, unix, "admin")

		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			updateUser.Username,
			updateUser.ID,
		).WillReturnRows(rows)

		mocksql.ExpectPrepare(query1)
		mocksql.ExpectExec(query1).WithArgs(
			updateUser.Username,
			unix,
			updateUser.ID,
			updateUser.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUsername(context.TODO(), tx, updateUser)
		t.Log(err)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.NotEqual(t, user, &updateUser)

		err = mocksql.ExpectationsWereMet()
		assert.Error(t, err)
	})
}
