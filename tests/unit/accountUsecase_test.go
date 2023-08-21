package unit

import (
	"bytes"
	"context"
	"database/sql"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func multipartFileHeader() *multipart.FileHeader {
	fileContent := []byte("Contoh isi file")
	fileHeader := &multipart.FileHeader{
		Filename: "example.png",
		Size:     int64(len(fileContent)),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	part.Write(fileContent)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return fileHeader
}

func TestAccountUsecase(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	userRepoMock := &mocks.FakeUserRepo{}
	tsxSqlRepoMock := &mocks.FakeSqlTransactionRepo{}
	minioRepoMock := &mocks.FakeMinioRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()

	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepoMock, userRepoMock, tsxSqlRepoMock, minioRepoMock, timeOutCtx)

	// t.Run("SUCCESS_AccountUpdate", func(t *testing.T) {
	// 	profile := model.Profile{
	// 		ProfileId: "profileid_1",
	// 		UserId:    "userid_1",
	// 		Quote:     sql.NullString{String: ""},
	// 		CreatedAt: time.Now().Unix(),
	// 		CreatedBy: "profileid_1",
	// 		UpdatedAt: time.Now().Unix(),
	// 		UpdatedBy: sql.NullString{},
	// 		DeletedAt: sql.NullInt64{},
	// 		DeletedBy: sql.NullString{},
	// 	}

	// 	user := model.User{
	// 		ID:              "userid_1",
	// 		FullName:        "rama_1",
	// 		Gender:          "undefinied",
	// 		Image:           "default-male.png",
	// 		Username:        "ibanrmaa_1",
	// 		Email:           "1_ibanrama29@gmail.com",
	// 		Password:        "123456",
	// 		PhoneNumber:     sql.NullString{},
	// 		EmailVerifiedAt: true,
	// 		CreatedAt:       time.Now().Unix(),
	// 		CreatedBy:       "userid_1",
	// 		UpdatedAt:       time.Now().Unix(),
	// 		UpdatedBy:       sql.NullString{},
	// 		DeletedAt:       sql.NullInt64{},
	// 		DeletedBy:       sql.NullString{},
	// 	}

	// 	req := domaindto.UpdateAccountReq{
	// 		UserID:      "userid_1",
	// 		FullName:    "rama_update_1",
	// 		Gender:      "male",
	// 		Image:       multipartFileHeader(),
	// 		PhoneNumber: "1234567890",
	// 		Quote:       "semangat_update_1",
	// 	}

	// 	profileRepoMock.GetProfileByUserId(ctx, "userid_1")
	// 	profileRepoMock.GetProfileByUserIdReturns(&profile, nil)

	// 	userRepoMock.GetUserById(ctx, "userid_1")
	// 	userRepoMock.GetUserByIdReturns(&user, nil)

	// 	minioReturnMock := minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
	// 	minioRepoMock.GenerateFileNameReturns(minioReturnMock)

	// 	minioRepoMock.UploadFile(ctx, multipartFileHeader(), minioReturnMock, "files")
	// 	minioRepoMock.UploadFileReturns(nil)

	// 	profileRes, userRes, err := accountUsecase.AccountUpdate(ctx, req)
	// 	assert.NoError(t, err)
	// 	assert.NotNil(t, profileRes)
	// 	assert.NotNil(t, userRes)
	// })

	t.Run("SUCCESS_AccountUpdate_WithDeleteFile", func(t *testing.T) {
		profile := model.Profile{
			ProfileId: "profileid_1",
			UserId:    "userid_1",
			Quote:     sql.NullString{String: ""},
			CreatedAt: time.Now().Unix(),
			CreatedBy: "profileid_1",
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		}

		user := model.User{
			ID:              "userid_1",
			FullName:        "rama_1",
			Gender:          "undefinied",
			Image:           "rama.png",
			Username:        "ibanrmaa_1",
			Email:           "1_ibanrama29@gmail.com",
			Password:        "123456",
			PhoneNumber:     sql.NullString{},
			EmailVerifiedAt: true,
			CreatedAt:       time.Now().Unix(),
			CreatedBy:       "userid_1",
			UpdatedAt:       time.Now().Unix(),
			UpdatedBy:       sql.NullString{},
			DeletedAt:       sql.NullInt64{},
			DeletedBy:       sql.NullString{},
		}

		req := dto.UpdateAccountReq{
			UserID:      "userid_1",
			FullName:    "rama_update_1",
			Gender:      "male",
			Image:       multipartFileHeader(),
			PhoneNumber: "1234567890",
			Quote:       "semangat_update_1",
		}

		profileRepoMock.GetProfileByUserId(ctx, "userid_1")
		profileRepoMock.GetProfileByUserIdReturns(&profile, nil)

		userRepoMock.GetUserById(ctx, "userid_1")
		userRepoMock.GetUserByIdReturns(&user, nil)

		tsxSqlRepoMock.Transaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error { return nil })
		// minioReturnMock := minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
		// minioRepoMock.GenerateFileNameReturns(minioReturnMock)

		// minioRepoMock.UploadFile(ctx, multipartFileHeader(), minioReturnMock, "files")
		// minioRepoMock.UploadFileReturns(nil)

		profileRes, userRes, err := accountUsecase.AccountUpdate(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, profileRes)
		assert.NotNil(t, userRes)
	})
}
