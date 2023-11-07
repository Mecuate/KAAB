package models

import "github.com/gorilla/mux"

type Server struct {
	Configuration *EnvConfigs
	Router        MuxRouter
}

type MuxRouter struct {
	Router *mux.Router
}
