package integration

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func AccountUpdateUSECASE(t *testing.T) {
	config.MinIoBucket = "files"
	uow := repository.NewUnitOfWorkImpl(db)
	profileRepo := repository.NewProfileRepoImpl(uow)
	userRepo := repository.NewUserRepoImpl(uow)
	minio := repository.NewMinioImpl(minioClient)
	timeOut := 2 * time.Second
	account := usecase.NewAccountUsecaseImpl(profileRepo, userRepo, minio, timeOut)

	fileContent := []byte("file content")
	fileHeader := &multipart.FileHeader{
		Filename: "ramaUpdate.png",
		Size:     int64(len(fileContent)),
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	assert.NoError(t, err)
	part.Write(fileContent)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("file")
	assert.NoError(t, err)
	defer file.Close()

	accountUpdate := dto.UpdateAccountReq{
		ProfileID:   "profileid1",
		UserID:      "userId1",
		FullName:    "rama_update_usecase",
		Gender:      "male",
		Image:       fileHeader,
		PhoneNumber: "12345678",
		Quote:       "semangat_update_usecase",
	}

	t.Run("SUCCESS_AccountUpdate", func(t *testing.T) {
		userResp, profileResp, err := account.UpdateAccount(context.Background(), accountUpdate)
		assert.NoError(t, err)
		assert.NotNil(t, userResp)
		assert.NotNil(t, profileResp)
	})
}
