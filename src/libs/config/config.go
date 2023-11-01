package config

import (
	"fmt"
	"kaab/src/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "KAAB"
	logPrefix = "LOG"
	appPrefix = "APP"
)

var WEBENV = &models.WebConfigs{}
var APPENV = &models.AppConfig{}
var LOGENV = &models.LoggingConfig{}

func FromEnv() (config *models.EnvConfigs, err error) {
	args := os.Args
	var configFile, logFile, appFile = args[1], args[2], args[3]
	loadEnvFile(configFile)
	loadEnvFile(logFile)
	loadEnvFile(appFile)

	err = envconfig.Process(logPrefix, LOGENV)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process(appPrefix, APPENV)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process(envPrefix, WEBENV)
	if err != nil {
		return nil, err
	}

	config = &models.EnvConfigs{
		WebServerConfig: WEBENV,
		LoggingConfig:   LOGENV,
		EnvConfig:       APPENV,
	}
	LoadLogger()

	return config, nil
}

func loadEnvFile(cfgFileName string) {
	if cfgFileName != "" {
		err := godotenv.Load(cfgFileName)
		if err != nil {
			Err(fmt.Sprintf("Environment file:[%s] not found", cfgFileName))
		}
	}
}
