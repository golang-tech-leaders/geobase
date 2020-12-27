package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	AppPort        string `yaml:"port" env:"PORT"`
	ReqTimeoutSec  int    `yaml:"timeout" env:"REQTIMEOUTSEC" env-default:"10"`
	DataPath       string `yaml:"datapath" env:"LOCATION_DATA_PATH"`
	MetersInRadius int    `yaml:"meters_in_radius" env:"METERS_IN_RADIUS" env-default:"100"`
}

type LogConfig struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"INFO"`
}

type Config struct {
	AppConf AppConfig `yaml:"app"`
	LogConf LogConfig `yaml:"logging"`
}

func PrepareConfig(configFilePath string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configFilePath, &cfg); err != nil {
		return nil, fmt.Errorf("configuration read: %v", err)
	}
	return &cfg, nil
}
