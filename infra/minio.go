package infra

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func NewMinioConn(endPoint, id, secretKey string, ssl bool) (*minio.Client, error) {
	minioConn, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(id, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Warn().Msgf("failed open minio connection | err : %v", err)
	}

	return minioConn, err
}
