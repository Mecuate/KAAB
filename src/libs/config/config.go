package config

import (
	"fmt"
	"kaab/src/models"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "KAAB"
)

func FromEnv() (config *models.ServiceConfig, err error) {
	fromFileToEnv()

	cfg := &models.WebConfigs{}
	err = envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}

	lcfg := &models.LoggingConfig{}
	err = envconfig.Process("", lcfg)
	if err != nil {
		return nil, err
	}

	ecf := &models.EnvConfs{}
	err = envconfig.Process(envPrefix, ecf)
	if err != nil {
		return nil, err
	}

	config = &models.ServiceConfig{
		WebServerConfig: cfg,
		LoggingConfig:   lcfg,
		EnvConfig:       ecf,
	}

	return config, nil
}

func fromFileToEnv() {

	cfgFileName := "./config/config.env"
	if cfgFileName != "" {

		err := godotenv.Load(cfgFileName)
		if err != nil {
			fmt.Println("ENV_FILE not found")
		}
	}
	return
}
