package config

import (
	"monolith/pkg/slog"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string `mapstructure:"HTTP_PORT"`

	DatabaseDsn string `mapstructure:"DATABASE_DSN"`

	TokenSalt    string `mapstructure:"TOKEN_SALT"`
	PasswordSalt string `mapstructure:"PASSWORD_SALT"`

	Environment string `mapstructure:"ENVIRONMENT"`
	IsDebug     bool   `mapstructure:"IS_DEBUG"`
}

func New(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		slog.Fatal(err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		slog.Fatal(err.Error())
	}

	return config
}
