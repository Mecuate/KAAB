package main

import (
	"kaab/src/libs/config"
	"kaab/src/libs/server"
)

func main() {
	config.LoadLogger()
	err := server.RunServer()
	if err != nil {
		config.Log("kaab server application error")
		config.Err(err.Error())
	}
}
