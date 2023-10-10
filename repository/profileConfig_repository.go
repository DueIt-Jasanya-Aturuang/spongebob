package repository

import (
	"context"
)

type ProfileConfigRepository interface {
	Create(ctx context.Context, profileCfg *ProfileConfig) (bool, error)
	Update(ctx context.Context, profileCfg *ProfileConfig) error
	GetByNameAndID(ctx context.Context, profileID string, configName string) (*ProfileConfig, error)
	GetBySchedulerDailyNotify(ctx context.Context, minuteSecond string, day string) (*[]ProfileConfig, error)
	GetBySchedulerMonthlyPeriode(ctx context.Context, tgl int, id string) (*[]ProfileConfig, error)
	UnitOfWorkRepository
}

type ProfileConfig struct {
	ID          string
	ProfileID   string
	ConfigName  string
	ConfigValue string
	Status      string
	AuditInfo
}
