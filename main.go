package main

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/server"
)

func main() {
	config.LoadLogger()
	err := server.RunServer()
	if err != nil {

		fmt.Println("kaab server application error")
		fmt.Println(err)
	}
}
