package integration

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/repositories"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

func TestMinioImpl(t *testing.T) {
	fileContent := []byte("Contoh isi file")
	fileHeader := &multipart.FileHeader{
		Filename: "example.png",
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

	err = minioClient.MakeBucket(context.Background(), "files", minio.MakeBucketOptions{})
	assert.NoError(t, err)

	minioIMPL := repositories.NewMinioImpl(minioClient)

	filename := minioIMPL.GenerateFileName(fileHeader, "user-images/public/", "")
	t.Run("SUCCESS_Upload", func(t *testing.T) {
		err = minioIMPL.UploadFile(context.Background(), fileHeader, filename, "files")
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_Delete", func(t *testing.T) {
		err = minioIMPL.DeleteFile(context.Background(), filename, "files")
		assert.NoError(t, err)
	})
}
