package models

type Server struct {
	Configuration *EnvConfigs
	Router        MuxRouter
}
