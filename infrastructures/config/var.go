package config

import "time"

const (
	setMaxIdleConnsDB    = 5
	setMaxOpenConnsDB    = 100
	SetConnMaxIdleTimeDB = 5 * time.Minute
	setConnMaxLifetimeDB = 60 * time.Minute
	pgPingTimeOut        = 5 * time.Second
)
