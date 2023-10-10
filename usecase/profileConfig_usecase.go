package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type ProfileConfigUsecase interface {
	Create(ctx context.Context, req *RequestCreateProfileConfig) (*ResponseProfileConfig, error)
	GetByNameAndID(ctx context.Context, req *RequestGetProfileConfig) (*ResponseProfileConfig, error)
	Update(ctx context.Context, req *RequsetUpdateProfileConfig) (*ResponseProfileConfig, error)
	SchedulerDailyNotify(ctx context.Context, minuteSecond string, day string) error
	SchedulerMonthlyPeriode(ctx context.Context, tgl int, id *string) (*string, error)
}

type RequestCreateProfileConfig struct {
	ConfigValue  string
	Days         []string
	ConfigName   string
	Status       string
	Token        string
	UserID       string
	ProfileID    string
	Value        string
	IanaTimezone string
}

type RequsetUpdateProfileConfig struct {
	ConfigValue  string
	Days         []string
	Status       string
	Token        string
	ProfileID    string
	UserID       string
	ConfigName   string
	Value        string
	IanaTimezone string
}

type RequestGetProfileConfig struct {
	UserID     string
	ConfigName string
	ProfileID  string
}

type ResponseProfileConfig struct {
	ID          string
	ProfileID   string
	ConfigName  string
	ConfigValue string
	Status      string
	Days        []string
	Token       string
}

func (req *RequestCreateProfileConfig) ToModel(configValue []byte) *repository.ProfileConfig {
	return &repository.ProfileConfig{
		ID:          util.NewUlid(),
		ProfileID:   req.ProfileID,
		ConfigName:  req.ConfigName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func (req *RequsetUpdateProfileConfig) ToModel(configValue []byte, configName, id string) *repository.ProfileConfig {
	return &repository.ProfileConfig{
		ID:          id,
		ProfileID:   req.ProfileID,
		ConfigName:  configName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{String: req.ProfileID},
		},
	}
}
