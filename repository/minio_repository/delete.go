package minio_repository

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (m *MinioRepositoryImpl) DeleteFile(ctx context.Context, objectName string) error {
	if err := m.client.RemoveObject(ctx, infra.MinIoBucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		log.Warn().Msgf(util.LogErrDelObjectMinio, err)
		return err
	}

	return nil
}
