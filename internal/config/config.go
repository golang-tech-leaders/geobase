package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	AppPort       string `yaml:"port" env:"PORT"`
	ReqTimeoutSec int    `yaml:"timeout" env:"REQTIMEOUTSEC" env-default:"10"`
}

type LogConfig struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"INFO"`
}

type Config struct {
	AppConf AppConfig `yaml:"app"`
	LogConf LogConfig `yaml:"logging"`
}

func PrepareConfig(configFilePath string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configFilePath, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}
	return &cfg
}
