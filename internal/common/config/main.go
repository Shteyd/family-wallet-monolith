package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string `mapstructure:"HTTP_PORT"`

	PostgresDsn string `mapstructure:"POSTGRES_DSN"`

	Environment string `mapstructure:"ENVIRONMENT"`
	IsDebug     bool   `mapstructure:"IS_DEBUG"`
}

func New(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile("/.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err.Error())
	}

	return config
}
