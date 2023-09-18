package config

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
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
		log.Err(err).Msg(message.ErrOpenConnMinio)
	}

	return minioConn, err
}
