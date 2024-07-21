package config

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	Database DBConfig     `mapstructure:"database"`
	JWT      JWTConfig    `mapstructure:"jwt"`
}
type ServerConfig struct {
	Env  string     `mapstructure:"env"  default:"production"`
	Http RestConfig `mapstructure:"http" `
}
type RestConfig struct {
	Host string `mapstructure:"host" default:"localhost"`
	Port int    `mapstructure:"port" default:"8080"`
}

type DBConfig struct {
	Driver                string `mapstructure:"driver" default:"postgres"`
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	DBName                string `mapstructure:"name"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	SchemaName            string `mapstructure:"schema-name"`
	TimeZone              string `mapstructure:"time-zone" default:"Asia/Ho_Chi_Minh"`
	ConnectTimeout        int    `mapstructure:"connect-timeout" default:"5"`
	MaxIdleConnection     int    `json:"max-idle-connections" default:"20"`
	MaxOpenConnection     int    `mapstructure:"max-open-connections" default:"100"`
	ConnectionMaxLifeTime int    `mapstructure:"connection-max-lifetime" default:"1200"`
	ConnectionMaxIdleTime int    `mapstructure:"connection-max-idle-time" default:"1"`
	MaxRetry              int    `mapstructure:"max-retry" default:"5"`
	RetryDelay            int    `mapstructure:"retry-delay" default:"2"`
	Multiplier            int    `mapstructure:"multiplier" default:"2"`
}

// GetDriver returns the database driver
func (c *DBConfig) GetDriver() string {
	return c.Driver
}

// GetHost returns the database host
func (c *DBConfig) GetHost() string {
	return c.Host
}

// GetPort returns the database port
func (c *DBConfig) GetPort() int {
	return c.Port
}

// GetDBName returns the database name
func (c *DBConfig) GetDBName() string {
	return c.DBName
}

// GetUsername returns the database username
func (c *DBConfig) GetUsername() string {
	return c.Username
}

// GetPassword returns the database password
func (c *DBConfig) GetPassword() string {
	return c.Password
}

// GetSchemaName returns the database schema name
func (c *DBConfig) GetSchemaName() string {
	return c.SchemaName
}

// GetTimeZone returns the database time zone
func (c *DBConfig) GetTimeZone() string {
	return c.TimeZone
}

// GetConnectTimeout returns the database connection timeout
func (c *DBConfig) GetConnectTimeout() int {
	return c.ConnectTimeout
}

// GetMaxIdleConnection returns the maximum number of idle connections
func (c *DBConfig) GetMaxIdleConnection() int {
	return c.MaxIdleConnection
}

// GetMaxOpenConnection returns the maximum number of open connections
func (c *DBConfig) GetMaxOpenConnection() int {
	return c.MaxOpenConnection
}

// GetConnectionMaxLifeTime returns the maximum connection lifetime
func (c *DBConfig) GetConnectionMaxLifeTime() int {
	return c.ConnectionMaxLifeTime
}

// GetConnectionMaxIdleTime returns the maximum connection idle time
func (c *DBConfig) GetConnectionMaxIdleTime() int {
	return c.ConnectionMaxIdleTime
}

// GetConnectionMaxIdleTime returns the maximum connection idle time
func (c *DBConfig) GetMaxRetry() int {
	return c.MaxRetry
}

func (c *DBConfig) GetRetryDelay() int {
	return c.RetryDelay
}

func (c *DBConfig) GetMultiplier() int {
	return c.Multiplier
}

type JWTConfig struct {
	HashSecret           string `mapstructure:"hash-secret"`
	AccessTokenLifespan  int    `mapstructure:"access-token-lifespan" default:"1"`
	RefreshTokenLifespan int    `mapstructure:"refresh-token-lifespan" default:"24"`
}
