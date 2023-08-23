package repository

import (
	"context"
	"mime/multipart"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ./../mocks . MinioRepo
type MinioRepo interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, objectName, bucket string) error
	DeleteFile(ctx context.Context, objectName, bucket string) error
	GenerateFileName(file *multipart.FileHeader, path, prefix string) string
}
