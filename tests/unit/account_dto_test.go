package unit

import (
	"bytes"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

//goland:noinspection ALL
func TestAccountDTO(t *testing.T) {
	fileContent := []byte("image/png")
	fileHeader := &multipart.FileHeader{
		Filename: "rama.png",
		Size:     int64(len(fileContent)),
		Header:   make(textproto.MIMEHeader),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", fileHeader.Filename)
	assert.NoError(t, err)
	_, err = part.Write(fileContent)
	assert.NoError(t, err)
	err = writer.Close()
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("image")
	assert.NoError(t, err)
	defer func(file multipart.File) {
		err := file.Close()
		assert.NoError(t, err)
	}(file)
	fileHeader.Header.Set("Content-Type", "image/jpeg")
	t.Logf("data %v", fileHeader.Header.Get("Content-Type"))

	t.Run("SUCCESS_AccountDTO", func(t *testing.T) {
		reqSuccess := dto.UpdateAccountReq{
			ProfileID:   "699137ef-1f24-46d7-82bf-862fde7b36d8",
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			FullName:    "rama",
			Profesi:     "123456",
			Gender:      "male",
			Image:       fileHeader,
			PhoneNumber: "123456789012",
			Quote:       "semangat rama",
		}

		err := validation.UpdateAccountValidate(&reqSuccess)
		assert.NoError(t, err)
	})

	t.Run("ERROR_AccountDTO", func(t *testing.T) {
		fileHeader.Header.Set("Content-Type", "image/asd")
		reqSuccess := dto.UpdateAccountReq{
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			FullName:    "",
			Gender:      "maleasdasd",
			Image:       fileHeader,
			PhoneNumber: "1234asd56789012",
			Quote:       "sama",
		}
		err := validation.UpdateAccountValidate(&reqSuccess)
		assert.Error(t, err)
		t.Log(err)
	})
}
