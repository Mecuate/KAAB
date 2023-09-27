package main

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/server"
)

func main() {
	env, _ := config.GetSysFlags()
	//  log port to work on to log_file
	fmt.Print(env)
	server.RunServer(env.PORT)
}
