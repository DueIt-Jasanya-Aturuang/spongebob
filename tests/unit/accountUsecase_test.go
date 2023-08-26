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
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/dtoconv"
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

func TestAccounUpdateUsecase(t *testing.T) {
	uow := &mocks.FakeUnitOfWork{}
	profileRepoMock := &mocks.FakeProfileRepo{}
	userRepoMock := &mocks.FakeUserRepo{}
	minioRepoMock := &mocks.FakeMinioRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()
	image := "user-images/public/asd.png"
	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepoMock, userRepoMock, minioRepoMock, timeOutCtx)

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userid_1",
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
		Image:           "default-male.png",
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

	profileRepoMock.GetProfileByUserID(ctx, "userid_1")
	profileRepoMock.GetProfileByUserIDReturns(&profile, nil)

	userRepoMock.GetUserByID(ctx, "userid_1")
	userRepoMock.GetUserByIDReturns(&user, nil)

	profileRepoMock.UoW()
	profileRepoMock.UoWReturns(uow)

	profileConv, userConv := dtoconv.UpdateAccountToModel(req, profile.ProfileID, user.Image)
	profileRepoMock.UpdateProfile(ctx, profileConv)
	profileRepoMock.UpdateProfileReturns(&profile, nil)

	userRepoMock.UoW()
	userRepoMock.UoWReturns(uow)

	uow.GetTx()
	uow.GetTxReturns(&sql.Tx{}, nil)
	uow.CallTx(nil)
	uow.CallTxReturns(nil)

	userRepoMock.UpdateUser(ctx, userConv)
	userRepoMock.UpdateUserReturns(&user, nil)

	minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
	minioRepoMock.GenerateFileNameReturns(image)

	minioRepoMock.UploadFile(ctx, multipartFileHeader(), image, "files")
	minioRepoMock.UploadFileReturns(nil)

	uow.EndTx(nil)
	uow.EndTxReturns(nil)

	profileRes, userRes, err := accountUsecase.UpdateAccount(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, profileRes)
	assert.NotNil(t, userRes)
}

func TestAccounUpdateWithDeleteFileUsecase(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	uow := &mocks.FakeUnitOfWork{}
	userRepoMock := &mocks.FakeUserRepo{}
	minioRepoMock := &mocks.FakeMinioRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()
	image := "/files/user-images/public/asd.png"

	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepoMock, userRepoMock, minioRepoMock, timeOutCtx)

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userid_1",
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
		Image:           image,
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

	profileRepoMock.GetProfileByUserID(ctx, "userid_1")
	profileRepoMock.GetProfileByUserIDReturns(&profile, nil)

	userRepoMock.GetUserByID(ctx, "userid_1")
	userRepoMock.GetUserByIDReturns(&user, nil)

	profileRepoMock.UoW()
	profileRepoMock.UoWReturns(uow)

	profileConv, userConv := dtoconv.UpdateAccountToModel(req, profile.ProfileID, user.Image)
	profileRepoMock.UpdateProfile(ctx, profileConv)
	profileRepoMock.UpdateProfileReturns(&profile, nil)

	userRepoMock.UoW()
	userRepoMock.UoWReturns(uow)

	userRepoMock.UpdateUser(ctx, userConv)
	userRepoMock.UpdateUserReturns(&user, nil)

	minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
	minioRepoMock.GenerateFileNameReturns("user-images/public/asd.png")

	minioRepoMock.UploadFile(ctx, multipartFileHeader(), image, "files")
	minioRepoMock.UploadFileReturns(nil)

	minioRepoMock.DeleteFile(ctx, image, "files")
	minioRepoMock.DeleteFileReturns(nil)

	profileRes, userRes, err := accountUsecase.UpdateAccount(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, profileRes)
	assert.NotNil(t, userRes)
}
