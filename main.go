package main

import (
	"fmt"
	"os"
	"strings"

	"go_server/src/libs/server"
)

func main() {
	all_params := os.Args
	selected_port := strings.Split(os.Args[3], ":")[1]
	fmt.Println(len(os.Args), all_params)
	server.RunServer(selected_port)
}
