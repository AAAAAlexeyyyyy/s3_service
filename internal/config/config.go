package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"github.com/AAAAAlexeyyyyy/s3_service/internal/config/modules"
)

type AppConfig struct {
	AppName string
	Logger  modules.Logger
}

var config *AppConfig

func LoadConfig() (*AppConfig, error) {
	_ = godotenv.Load()
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("error config parsing: %v", err)
	}
	return config, nil
}
