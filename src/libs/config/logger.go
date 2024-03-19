package config

import (
	"fmt"
	"kaab/src/models"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	LOG_PATH = ""
	ERR_PATH = ""
)

func Log(s string) {
	if WEBENV.Environment == "DEV" {
		fmt.Println("\033[33m○-", s+"\033[0m")
	}
	SaveToLog(fmt.Sprintf("├%s┼%s▌\n", time.Now(), s), true)
}

func Err(s string) {
	if WEBENV.Environment == "DEV" {
		fmt.Println("\033[35m#Err-", s+"\033[0m")
	}
	SaveToLog(fmt.Sprintf("├%s┼%s▌\n", time.Now(), s), false)
}

func SaveToLog(str string, diff bool) {
	if !fileExists(LOG_PATH) {
		fmt.Printf("err LOG_PATH: %v\n", LOG_PATH)
		makeFile(LOG_PATH, []byte("# log file\n"))
	}
	if !fileExists(ERR_PATH) {
		makeFile(ERR_PATH, []byte("# error file\n"))
	}

	if diff {
		appendToFile(LOG_PATH, str)
	} else {
		appendToFile(ERR_PATH, str)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func makeFile(fdir string, str []byte) {
	err := os.WriteFile(fdir, str, 0777)
	if err != nil {
		fmt.Printf("err makeFile: %v\n", err)
	}
}

func appendToFile(fdir string, payload string) {
	file, err := os.OpenFile(fdir, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(payload); err != nil {
		fmt.Printf("err appendToFile.WriteString : %v\n", err)
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
