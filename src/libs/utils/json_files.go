package utils

import (
	"encoding/json"
	"fmt"
	"kaab/src/libs/config"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func SaveFile(dataBytes []byte, filepath string) {
	err := os.WriteFile(filepath, dataBytes, 0666)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.saveFile: %v", err))
	}
}

func JSON(t interface{}) string {
	res, err := json.Marshal(t)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.JSON.Marshal: %v", err))
	}
	return string(res)
}
