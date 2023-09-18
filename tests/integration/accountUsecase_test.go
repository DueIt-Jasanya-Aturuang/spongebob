package integration

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
)

func AccountUpdateUSECASE(t *testing.T) {
	config.MinIoBucket = "files"
	minio := repository.NewMinioImpl(minioClient)
	timeOut := 2 * time.Second
	account := usecase.NewAccountUsecaseImpl(ProfileRepo, UserRepo, minio, timeOut)

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
		ProfileID:   profileID_1,
		UserID:      userID_1,
		FullName:    "rama_update_usecase",
		Gender:      "male",
		Image:       fileHeader,
		PhoneNumber: "123456782",
		Quote:       "semangat_update_usecase",
	}

	t.Run("SUCCESS_AccountUpdate", func(t *testing.T) {
		userResp, profileResp, err := account.UpdateAccount(context.Background(), &accountUpdate)
		assert.NoError(t, err)
		assert.NotNil(t, userResp)
		assert.NotNil(t, profileResp)
	})
}
