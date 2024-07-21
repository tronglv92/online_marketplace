package database

import "errors"

const (
	MysqlDBDriver    string = "mysql"
	PostgresDBDriver string = "postgres"
	SqliteDBDriver   string = "sqlite"
	MongoDBDriver    string = "mongodb"
)

type (
	DBOption struct {
		Migrate []any
	}

	Option func(s *DBOption)
)

func New(c Config, opts ...Option) (Database, error) {
	switch c.GetDriver() {
	case PostgresDBDriver, SqliteDBDriver, MysqlDBDriver:
		return NewGormConnect(c, opts...)
	}
	return nil, errors.New("the database driver does not support")
}

func Must(c Config, opts ...Option) Database {
	conn, err := New(c, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}
