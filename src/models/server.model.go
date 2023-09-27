package models

type Server struct {
	Configuration *ServiceConfig
	Router        MuxRouter
}
