package domain

import (
	"context"
)

type ProfileConfig struct {
	ID          string
	ProfileID   string
	ConfigName  string
	ConfigValue string
	Status      string
	AuditInfo
}

type RequestCreateProfileConfig struct {
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	ConfigName   string   `json:"config_name"`  // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	UserID       string   // request header
	ProfileID    string   // request param
	Value        string   // helper
	IanaTimezone string   // helper
}

type RequsetUpdateProfileConfig struct {
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	ProfileID    string   // url parameter
	UserID       string   // request header
	ConfigName   string   // url parameter
	Value        string   // helper
	IanaTimezone string   // helper
}

type RequestGetProfileConfig struct {
	UserID     string // request header
	ConfigName string // url parameter config-name
	ProfileID  string // url 		parameter profile-id
}

type ResponseProfileConfig struct {
	ID          string   `json:"profile_config_id"`
	ProfileID   string   `json:"profile_id"`
	ConfigName  string   `json:"config_name"`
	ConfigValue string   `json:"config_value"`
	Status      string   `json:"status"`
	Days        []string `json:"days,omitempty"`
}

type ProfileConfigScheduler struct {
	Day  string
	Time string
}

//counterfeiter:generate -o ./../mocks . ProfileConfigRepo
type ProfileConfigRepo interface {
	Create(ctx context.Context, profileCfg *ProfileConfig) (bool, error)
	Update(ctx context.Context, profileCfg *ProfileConfig) error
	GetByNameAndID(ctx context.Context, profileID string, configName string) (*ProfileConfig, error)
	GetBySchedulerDailyNotify(ctx context.Context, ProfileConfigScheduler ProfileConfigScheduler) (*[]ProfileConfig, error)
	GetBySchedulerMonthlyPeriode(ctx context.Context, tgl int, id string) (*[]ProfileConfig, error)
	UnitOfWorkRepository
}

//counterfeiter:generate -o ./../mocks . ProfileCfgUsecase
type ProfileConfigUsecase interface {
	Create(ctx context.Context, req *RequestCreateProfileConfig) (*ResponseProfileConfig, error)
	GetByNameAndID(ctx context.Context, req *RequestGetProfileConfig) (*ResponseProfileConfig, error)
	Update(ctx context.Context, req *RequsetUpdateProfileConfig) (*ResponseProfileConfig, error)
	SchedulerDailyNotify(ctx context.Context, ProfileConfigScheduler ProfileConfigScheduler) error
	SchedulerMonthlyPeriode(ctx context.Context, tgl int, id string) (string, error)
}
