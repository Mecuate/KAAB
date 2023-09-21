package config

import (
	"fmt"
	"os"
	"strings"

	m "kaab/src/models"
)

func GetSysFlags() (*m.CLIflags, error) {
	params := os.Args

	inputValues := &m.CLIflags{}

	for i := 0; i < len(params); i++ {
		selected_item := strings.Split(params[i], ":")
		flag := selected_item[0]
		value := selected_item[1]

		fmt.Println(selected_item)

		switch flag {
		case "--port":
			inputValues.PORT = value
		case "--name":
			inputValues.NAME = value
		case "--log":
			inputValues.LOGDIR = value
		}

	}

	return inputValues, nil
}
