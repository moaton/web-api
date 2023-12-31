package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/moaton/web-api/pkg/logger"
)

type Config struct {
	PostgresHost         string `envconfig:"POSTGRES_HOST"`
	PostgresPort         string `envconfig:"POSTGRES_PORT"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	Secret               string `envconfig:"SECRET"`
	AccessTokenExpMinute int64  `envconfig:"ACCESS_TOKEN_EXPIRY_MINUTE"`
	RefreshTokenExpDays  int64  `envconfig:"REFRESH_TOKEN_EXPIRY_DAYS"`

	IsDebug bool `envconfig:"IS_DEBUG"`
}

func GetConfig() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Fatalf("envconfig.Process err %v", err)
	}
	return &cfg
}
