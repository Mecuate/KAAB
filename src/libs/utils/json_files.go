package utils

import (
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

type OpenFILE struct {
	Filename  string
	DataModel interface{}
}

func (f *OpenFILE) JSON() {
	data, err := os.ReadFile(f.Filename)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.JSON.ReadFile: %v", err))
	}
	err = ParseJSON(data, &f.DataModel)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.JSON.ParseJSON: %v", err))
	}
}
