package config

import "time"

type Config struct {
	ServerAddress             string        `mapstructure:"SERVER_ADDRESS"`
	DBUri                     string        `mapstructure:"DB_URI"`
	AccessTokenDuration       time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}
