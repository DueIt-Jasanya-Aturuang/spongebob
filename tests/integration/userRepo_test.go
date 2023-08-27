package integration

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/stretchr/testify/assert"
)

var (
	unixUser   = time.Now().Unix()
	image      = "default-male.png"
	createUser = model.User{
		ID:              "userId1",
		FullName:        "rama",
		Gender:          "undefinied",
		Image:           image,
		Username:        "ibanrmaa",
		Email:           "ibanrama29@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{},
		EmailVerifiedAt: false,
		CreatedAt:       unixUser,
		CreatedBy:       "c",
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

func TestUserREPO(t *testing.T) {
	uow := repository.NewUnitOfWorkImpl(db)
	userRepo := repository.NewUserRepoImpl(uow)
	fmt.Println("RUNNING TEST USER REPOSITORY")
	updateUser := model.User{
		ID:              "userId1",
		FullName:        "rama",
		Gender:          "male",
		Image:           image,
		Username:        "ibanrmaaasd9",
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
	updateUser1 := model.User{
		ID:              "userId2",
		FullName:        "rama2",
		Gender:          "male",
		Image:           image,
		Username:        "ibanrmaa2",
		Email:           "ibanrama292@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{String: "12345678", Valid: true},
		EmailVerifiedAt: true,
		CreatedAt:       unixUser,
		CreatedBy:       "userId2",
		UpdatedAt:       unixUser,
		UpdatedBy:       sql.NullString{String: "userId2", Valid: true},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}
	createUserFunc()
	t.Run("SUCCESS_GetUserByID", func(t *testing.T) {
		user, err := userRepo.GetUserByID(context.Background(), createUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &createUser, user)
	})

	t.Run("ERROR_GetUserByID_NOWROW", func(t *testing.T) {
		user, err := userRepo.GetUserByID(context.Background(), "createUser.ID")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, sql.ErrNoRows)
	})

	t.Run("SUCCESS_UpdateUser", func(t *testing.T) {
		err := userRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		user, err := userRepo.UpdateUser(context.TODO(), updateUser1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, &updateUser1, user)
		assert.NotEqual(t, &createUser, user)
		userRepo.UoW().EndTx(err)
	})

	t.Run("ERROR_UpdateUser_PHONEEXISTS", func(t *testing.T) {
		err := userRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		user, err := userRepo.UpdateUser(context.TODO(), updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, exception.Err400PhoneAvailable)
		userRepo.UoW().EndTx(err)
	})

	t.Run("ERROR_UpdateUsername_USERNAMEEXISTS", func(t *testing.T) {
		err := userRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		updateUser.Username = "rama2"
		user, err := userRepo.UpdateUsername(context.TODO(), updateUser)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, err, exception.Err400UsernameAvailable)
		userRepo.UoW().EndTx(err)
	})

	t.Run("SUCCESS_UpdateUsername", func(t *testing.T) {
		err := userRepo.UoW().StartTx(context.TODO(), &sql.TxOptions{ReadOnly: false})
		assert.NoError(t, err)
		updateUser.Username = "updateusernamerama"
		user, err := userRepo.UpdateUsername(context.TODO(), updateUser)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		userRepo.UoW().EndTx(err)
	})
	t.Run("AccountUpdateUSECASE", AccountUpdateUSECASE)
}
