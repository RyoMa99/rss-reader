package config

import (
	"app/infra/logger"
	"os"

	"github.com/spf13/viper"
)

func Init() *Config {
	conf := &Config{}

	envFile := os.Getenv("CONFIG_FILE")
	if envFile == "" {
		envFile = "../config.dev.yaml"
	}
	viper.SetConfigFile(envFile)
	if err := viper.ReadInConfig(); err != nil {
		logger.Panic(err.Error())
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(conf); err != nil {
		logger.Panic(err.Error())
	}

	return conf
}
