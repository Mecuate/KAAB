package config

import (
	"fmt"
	"kaab/src/models"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var (
	LOG_PATH = ""
	ERR_PATH = ""
)

func Log(s string) {
	SaveToLog([]byte(s), true)
}

func Err(s string) {
	SaveToLog([]byte(s), false)
}

func SaveToLog(str []byte, t bool) {
	var err interface{}

	if t {
		err = os.WriteFile(LOG_PATH, str, 0666)
	} else {
		err = os.WriteFile(ERR_PATH, str, 0666)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func LoadLogger() (err error) {
	loadEnvFile("./config/config.log.env")
	logConf := &models.LoggingConfig{}
	err = envconfig.Process(logPrefix, logConf)
	if err != nil {
		return err
	}
	LOG_PATH = fmt.Sprintf("%s/%s", logConf.LogPath, logConf.LogFileName)
	ERR_PATH = fmt.Sprintf("%s/%s", logConf.LogPath, logConf.ErrorFileName)
	return nil
}
