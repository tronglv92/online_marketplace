package database

import (
	"gorm.io/gorm"
)

type Database interface {
	GetGormClient() *gorm.DB
}

type Config interface {
	GetDriver() string
	GetHost() string
	GetPort() int
	GetDBName() string
	GetUsername() string
	GetPassword() string
	GetSchemaName() string
	GetTimeZone() string
	GetConnectTimeout() int
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnectionMaxLifeTime() int
	GetConnectionMaxIdleTime() int
	GetMaxRetry() int
	GetRetryDelay() int
	GetMultiplier() int
}
