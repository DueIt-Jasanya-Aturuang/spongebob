package repository

import (
	"context"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

type MinioImpl struct {
	c *minio.Client
}

func NewMinioImpl(c *minio.Client) repository.MinioRepo {
	minioPool := &MinioImpl{
		c: c,
	}

	return minioPool
}

func (m *MinioImpl) UploadFile(ctx context.Context, file *multipart.FileHeader, objectName, bucket string) error {
	fileReader, err := file.Open()
	if err != nil {
		log.Err(err).Msg(message.ErrOpenFile)
		return err
	}
	defer func() {
		errCloseFile := fileReader.Close()
		if errCloseFile != nil {
			log.Err(err).Msg(message.ErrCloseFile)
		}
	}()

	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err := m.c.PutObject(ctx, bucket, objectName, fileReader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Err(err).Msg(message.ErrPutMinio)
		return err
	}

	log.Info().Msgf("info upload : %v", info)
	return nil
}

func (m *MinioImpl) DeleteFile(ctx context.Context, objectName, bucket string) error {
	if err := m.c.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		log.Err(err).Msg(message.ErrDelMinio)
		return err
	}

	return nil
}

func (m *MinioImpl) GenerateFileName(file *multipart.FileHeader, path, prefix string) string {
	nameFile := fmt.Sprintf("%s%s%d%s", path, prefix, time.Now().UnixNano(), filepath.Ext(file.Filename))
	return nameFile
}
