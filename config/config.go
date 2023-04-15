package config

import (
	"github.com/joeshaw/envdecode"
)

// Config define product-integrator config structure
type Config struct {
	AppEnv  string `env:"APP_ENV"`
	ApiPort uint16 `env:"API_PORT,default=8080"`
	MysqlOption
}

// MysqlOption contains mySQL connection options
type MysqlOption struct {
	Driver       string `env:"MYSQL_DRIVER,default="`
	Host         string `env:"MYSQL_HOST,required"`
	Port         string `env:"MYSQL_PORT,required"`
	Pool         int    `env:"MYSQL_POOL,required"`
	DatabaseName string `env:"MYSQL_DATABASE_NAME,required"`
	Username     string `env:"MYSQL_USERNAME,required"`
	Password     string `env:"MYSQL_PASSWORD"`
	TimeZone     string `env:"MYSQL_TIMEZONE,required"`
}

// NewConfig initialize new config
func NewConfig() *Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
