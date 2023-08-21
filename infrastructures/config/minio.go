package config

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func NewMinioConn(endPoint, id, secretKey string, ssl bool) (*minio.Client, error) {
	minio, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(id, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrMinioConn)
	}

	return minio, err
}
