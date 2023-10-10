package minio_repository

import (
	"github.com/minio/minio-go/v7"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type MinioRepositoryImpl struct {
	client *minio.Client
}

func NewMinioRepositoryImpl(client *minio.Client) repository.MinioRepository {
	return &MinioRepositoryImpl{
		client: client,
	}
}
