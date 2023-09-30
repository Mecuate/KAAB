package config

import (
	"fmt"
	"kaab/src/models"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "KAAB"
	logPrefix = "LOG"
	appPrefix = "APP"
)

func FromEnv() (config *models.EnvConfigs, err error) {
	loadEnvFile("./config/config.env")
	loadEnvFile("./config/config.app.env")
	loadEnvFile("./config/config.log.env")

	webConf := &models.WebConfigs{}
	err = envconfig.Process(envPrefix, webConf)
	if err != nil {
		return nil, err
	}

	logConf := &models.LoggingConfig{}
	err = envconfig.Process(logPrefix, logConf)
	if err != nil {
		return nil, err
	}

	envConf := &models.AppConfig{}
	err = envconfig.Process(appPrefix, envConf)
	if err != nil {
		return nil, err
	}

	config = &models.EnvConfigs{
		WebServerConfig: webConf,
		LoggingConfig:   logConf,
		EnvConfig:       envConf,
	}

	return config, nil
}

func loadEnvFile(cfgFileName string) {
	if cfgFileName != "" {
		err := godotenv.Load(cfgFileName)
		if err != nil {
			fmt.Printf("Environment file:[%s] not found", cfgFileName)
		}
	}
}
