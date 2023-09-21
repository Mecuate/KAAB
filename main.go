package main

import (
	"fmt"

	"kaab/src/libs/config"
	"kaab/src/libs/server"
)

func main() {
	env, _ := config.GetSysFlags()
	fmt.Println("-- port is:: ", env)

	server.RunServer(env.PORT)
}
