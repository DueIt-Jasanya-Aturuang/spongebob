package unit

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDREPO(t *testing.T) {
	gender := "male"
	userColumns := []string{"id", "fullname", "gender", "image", "email", "username", "password", "phone_number", "email_verified_at", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	image := "default-male.png"
	email := "ibanrama29@gmail.com"
	userId1 := "userid_1"
	userId2 := "userid_2"

	unix := time.Now().Unix()
	expectUser := model.User{
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

	userRepo := repository.NewUserRepoImpl(db)
	query := regexp.QuoteMeta("SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM auth.m_users WHERE id = $1 AND deleted_at IS NULL")

	t.Run("SUCCESS_GetUserByID", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", email, "123456", "12345678", true, unix, userId1, unix, nil, nil, nil)
		rows.AddRow(userId2, "2rama", gender, image, "2ibanrmaa", "2ibanrmaa29@gmail.com", "1234567", "123456789", true, unix, userId2, unix, nil, nil, nil)

		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1).WillReturnRows(rows)

		user, err := userRepo.GetUserByID(context.TODO(), userId1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &expectUser, user)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_GetUserByID_DATANILL", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(userId1).WillReturnRows(sqlmock.NewRows(userColumns))

		user, err := userRepo.GetUserByID(context.TODO(), userId1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.NotEqual(t, &expectUser, user)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateUserREPO(t *testing.T) {
	gender := "male"
	userColumns := []string{"id", "fullname", "gender", "image", "email", "username", "password", "phone_number", "email_verified_at", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	image := "default-male.png"
	email := "ibanrama29@gmail.com"
	userId1 := "userid_1"

	unix := time.Now().Unix()
	updateUser := model.User{
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

	userRepo := repository.NewUserRepoImpl(db)
	query := regexp.QuoteMeta("SELECT phone_number, id FROM auth.m_users WHERE phone_number=$1 AND id<>$2 AND deleted_at IS NULL")
	query2 := regexp.QuoteMeta("UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, updated_at = $5, updated_by = $6 WHERE id = $7 AND deleted_at IS NULL")

	t.Run("SUCCESS_UpdateUser", func(t *testing.T) {
		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{}))
		mocksql.ExpectPrepare(query2)
		mocksql.ExpectExec(query2).WithArgs(
			"rama",
			gender,
			image,
			"12345678",
			unix,
			userId1,
			userId1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mocksql.ExpectCommit()

		err = userRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUser(context.TODO(), updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, &updateUser)

		err = userRepo.Commit()
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_UpdateUser_PHONEEXISTS", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", email, "123456", "12345678", true, unix, userId1, unix, nil, unix, "admin")

		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs("12345678", userId1).WillReturnRows(rows)

		err = userRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUser(context.TODO(), updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, exception.Err400PhoneAlvailable, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateUserUsernameREPO(t *testing.T) {
	gender := "male"
	userColumns := []string{"id", "fullname", "gender", "image", "email", "username", "password", "phone_number", "email_verified_at", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	image := "default-male.png"
	email := "ibanrama29@gmail.com"
	userId1 := "userid_1"

	unix := time.Now().Unix()
	updateUser := model.User{
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

	userRepo := repository.NewUserRepoImpl(db)
	query := regexp.QuoteMeta("SELECT username, id FROM auth.m_users WHERE username=$1 AND id<>$2 AND deleted_at IS NULL")
	query1 := regexp.QuoteMeta("UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 WHERE id = $4 AND deleted_at IS NULL")

	t.Run("SUCCESS_UpdateUserUsername", func(t *testing.T) {
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
		mocksql.ExpectCommit()

		err = userRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUsername(context.TODO(), updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, &updateUser)

		err = userRepo.Commit()
		assert.NoError(t, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_UpdateUserUsername_USERNAMEEXISTS", func(t *testing.T) {
		rows := sqlmock.NewRows(userColumns)
		rows.AddRow(userId1, "rama", gender, image, "ibanrmaa", email, "123456", "12345678", true, unix, userId1, unix, nil, unix, "admin")

		mocksql.ExpectBegin()
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			updateUser.Username,
			updateUser.ID,
		).WillReturnRows(rows)

		err = userRepo.BeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)

		user, err := userRepo.UpdateUsername(context.TODO(), updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.NotEqual(t, user, &updateUser)
		assert.Equal(t, exception.Err400UsernameAlvailable, err)

		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
