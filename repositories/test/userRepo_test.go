package test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/stretchr/testify/assert"
)

var (
	unixUser   = time.Now().Unix()
	createUser = domain.User{
		ID:              "userId1",
		FullName:        "rama",
		Gender:          "undefinied",
		Image:           "default-male.png",
		Username:        "ibanrmaa",
		Email:           "ibanrama29@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{},
		EmailVerifiedAt: false,
		CreatedAt:       unixUser,
		CreatedBy:       "userId1",
		UpdatedAt:       unixUser,
		UpdatedBy:       sql.NullString{},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}
)

func createUserFunc() {
	SQL := "INSERT INTO auth.m_users (id, fullname, image, username, email, password, email_verified_at, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING gender"
	_, err := db.ExecContext(
		context.TODO(), SQL,
		createUser.ID,
		createUser.FullName,
		createUser.Image,
		createUser.Username,
		createUser.Email,
		createUser.Password,
		createUser.EmailVerifiedAt,
		createUser.CreatedAt,
		createUser.CreatedBy,
		createUser.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	SQL = "INSERT INTO auth.m_users (id, fullname, image, username, email, password, email_verified_at, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING gender"
	createUser.Username = "rama2"
	createUser.Email = "ibanrama292@gmail.com"
	createUser.ID = "userId2"
	_, err = db.ExecContext(
		context.TODO(), SQL,
		createUser.ID,
		createUser.FullName,
		createUser.Image,
		createUser.Username,
		createUser.Email,
		createUser.Password,
		createUser.EmailVerifiedAt,
		createUser.CreatedAt,
		createUser.CreatedBy,
		createUser.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestUserRepo(t *testing.T) {
	updateUser := domain.User{
		ID:              "userId1",
		FullName:        "rama",
		Gender:          "male",
		Image:           "default-male.png",
		Username:        "ibanrmaa",
		Email:           "ibanrama29@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{String: "12345678", Valid: true},
		EmailVerifiedAt: true,
		CreatedAt:       unixUser,
		CreatedBy:       "userId1",
		UpdatedAt:       unixUser,
		UpdatedBy:       sql.NullString{String: "userId1", Valid: true},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}
	createUserFunc()
	t.Run("success get by id", func(t *testing.T) {
		user, err := userRepo.GetUserById(context.Background(), db, createUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &createUser, user)
	})

	t.Run("error get by id", func(t *testing.T) {
		user, err := userRepo.GetUserById(context.Background(), db, "createUser.ID")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("update user by id success", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		user, err := userRepo.UpdateUser(context.TODO(), tx, updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &updateUser, user)
		assert.NotEqual(t, &createUser, user)
	})

	t.Run("error update username", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		updateUser.Username = "rama2"
		user, err := userRepo.UpdateUsername(context.TODO(), tx, updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, domain.ErrUsernameAlvailable)
	})

	t.Run("success update username", func(t *testing.T) {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: false})
		if err != nil {
			t.Fatal(err)
		}
		user, err := userRepo.UpdateUsername(context.TODO(), tx, updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, domain.ErrUsernameAlvailable)
	})
}
