package main

import (
	"fmt"
	"kaab/src/libs/server"
)

func main() {
	err := server.RunServer()
	if err != nil {

		fmt.Println("main server application error")
		fmt.Println(err)
	}
}
