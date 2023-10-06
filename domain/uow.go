package domain

import (
	"context"
	"database/sql"
)

type AuditInfo struct {
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}

type RequestGetAllByPaginate struct {
	ID        string
	ProfileID string
	Order     string
	Operation string
}

//counterfeiter:generate -o ./mocks . UnitOfWorkRepository
type UnitOfWorkRepository interface {
	StartTx(ctx context.Context, opts *sql.TxOptions, fn func() error) error
	GetTx() (*sql.Tx, error)
	GetDB() (*sql.DB, error)
}
