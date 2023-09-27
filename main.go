package main

import (
	"kaab/src/libs/config"
	"kaab/src/libs/server"
)

func main() {
	env, _ := config.GetSysFlags()
	//  log port to work on to log_file

	server.RunServer(env.PORT)
}
