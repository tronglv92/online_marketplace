package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	defstruct "github.com/online_marketplace/helper/default_struct"
	"github.com/spf13/viper"

	"strings"
	"sync"
)

var m = &sync.Mutex{}
var c *Config

type CfgOption struct {
	ConfigPath string
	EnvPrefix  string
}

type Option func(f *CfgOption)

func WithConfigPath(configPath string) Option {
	return func(f *CfgOption) {
		f.ConfigPath = configPath
	}
}

func WithEnvPrefix(prefix string) Option {
	return func(f *CfgOption) {
		f.EnvPrefix = prefix
	}
}

func Init(opts ...Option) *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	kdOption := &CfgOption{ConfigPath: "etc/", EnvPrefix: ""}
	for _, applyOpt := range opts {
		applyOpt(kdOption)
	}

	k := NewViper(kdOption)
	err = k.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if e := defstruct.Set(c); e != nil {
		panic(fmt.Errorf("Fatal error when set the default config: %s \n", err))
	}

	return c
}

func Get(opts ...Option) *Config {
	if c == nil {
		m.Lock()
		defer m.Unlock()
		return Init(opts...)
	}
	return c
}

func NewViper(opt *CfgOption) *viper.Viper {
	// profile := "dev"
	// if os.Getenv("PMC_ENV") == "prd" || os.Getenv("PMC_ENV") == "stg" {
	// 	profile = "prd"
	// }

	var configFileName string
	configFileName = configFileName + "app"

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(configFileName)
	v.AddConfigPath(opt.ConfigPath)
	v.SetEnvPrefix(opt.EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Replace placeholders with actual environment values

	replacePlaceholders(v, v.AllSettings(), "")
	return v
}
func replacePlaceholders(viper *viper.Viper, settings map[string]interface{}, prefix string) {
	for key, value := range settings {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}
		switch v := value.(type) {
		case map[string]interface{}:
			replacePlaceholders(viper, v, fullKey)
		case string:
			replaced := os.ExpandEnv(v)
			viper.Set(fullKey, replaced)
		}
	}
}
