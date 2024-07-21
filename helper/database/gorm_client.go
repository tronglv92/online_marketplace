package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type gormSvc struct {
	dbConfig Config
	dbConn   *gorm.DB
	*DBOption
}

var (
	instance *gormSvc
	mutex    sync.Mutex
)

func NewGormConnect(dbConfig Config, opts ...Option) (Database, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if instance == nil {
		instance = &gormSvc{
			dbConfig: dbConfig,
			DBOption: &DBOption{},
		}
		for _, opt := range opts {
			opt(instance.DBOption)
		}
		err := instance.getConn()
		if err != nil {
			instance = nil
			return nil, err
		}
	}

	return instance, nil
}

func WithGormMigrate(items []any) Option {
	return func(s *DBOption) {
		s.Migrate = items
	}
}

func (s *gormSvc) GetGormClient() *gorm.DB {
	return s.dbConn
}

func (s *gormSvc) getConn() error {

	var db *gorm.DB
	var err error

	maxRetry := s.dbConfig.GetMaxRetry()
	retryDelay := s.dbConfig.GetRetryDelay()
	multiplier := s.dbConfig.GetMultiplier()
	for i := 0; i < maxRetry; i++ {
		db, err = gorm.Open(s.getDialect(), s.options())
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetry, err)

		var sleepTime int
		if i == 0 {
			sleepTime = retryDelay
		} else {
			sleepTime = retryDelay * multiplier
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
	if err != nil {
		return err
	}

	if len(s.Migrate) > 0 {
		if e := db.AutoMigrate(s.Migrate...); e != nil {
			return e
		}
	}

	s.dbConn = db
	sqlDB, e := s.dbConn.DB()
	if e != nil {
		return e
	}
	sqlDB.SetMaxIdleConns(s.dbConfig.GetMaxIdleConnection())
	sqlDB.SetMaxOpenConns(s.dbConfig.GetMaxOpenConnection())
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(s.dbConfig.GetConnectionMaxIdleTime()))
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(s.dbConfig.GetConnectionMaxIdleTime()))

	return nil

}

func (s *gormSvc) options() *gorm.Config {
	schemaName := ""
	if s.dbConfig.GetSchemaName() != "" {
		schemaName = s.dbConfig.GetSchemaName() + "."
	}
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   schemaName,
		},
	}
}

func (s *gormSvc) getDialect() gorm.Dialector {
	var dialect gorm.Dialector
	switch s.dbConfig.GetDriver() {
	case PostgresDBDriver:
		dialect = s.postgresOpen()
	case MysqlDBDriver:
		dialect = s.mysqlOpen()
	case SqliteDBDriver:
		dialect = s.sqliteOpen()
	}
	return dialect
}

func (s *gormSvc) postgresOpen() gorm.Dialector {
	return postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s TimeZone=%s",
		s.dbConfig.GetHost(),
		s.dbConfig.GetPort(),
		s.dbConfig.GetUsername(),
		s.dbConfig.GetPassword(),
		s.dbConfig.GetDBName(),
		"disable",
		s.dbConfig.GetSchemaName(),
		s.dbConfig.GetTimeZone(),
	))
}

func (s *gormSvc) mysqlOpen() gorm.Dialector {
	return mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		s.dbConfig.GetUsername(),
		s.dbConfig.GetPassword(),
		s.dbConfig.GetHost(),
		s.dbConfig.GetPort(),
		s.dbConfig.GetDBName(),
	))
}

func (s *gormSvc) sqliteOpen() gorm.Dialector {
	return sqlite.Open(s.dbConfig.GetDBName())
}
