package server

import (
	"fmt"
	"net/http"
	"os"

	cf "kaab/src/libs/config"
	"kaab/src/libs/handlers"
	"kaab/src/libs/utils"
	"kaab/src/models"

	"github.com/rs/cors"
)

func NewServer(serverConfig *models.EnvConfigs) *models.Server {
	server := &models.Server{
		Configuration: serverConfig,
		Router:        handlers.NewRouter(),
	}

	return server
}

func RunServer() (err error) {
	config, err := cf.FromEnv()
	if err != nil {
		return err
	}

	if os.Getenv("ENVIRONMENT") == "" {
		os.Setenv("ENVIRONMENT", "DEV")
		cf.Log(fmt.Sprintf("-- Server working ENVIRONMENT: %s", os.Getenv("ENVIRONMENT")))
	}

	server := NewServer(config)
	envs := config.WebServerConfig

	if utils.Boolean(envs.CorsEnabled).String() {
		CORSServer(config, server)
	} else {
		NormalServer(config, server)
	}

	return nil
}

func NormalServer(config *models.EnvConfigs, server *models.Server) {
	serverConfig := config.WebServerConfig
	cf.Log(fmt.Sprintf("KAAB --Normal server running: %s", os.Getenv("ENVIRONMENT")))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), server.Router.Router); err != nil {
		panic(err)
	}
}

func CORSServer(config *models.EnvConfigs, server *models.Server) {
	serverConfig := config.WebServerConfig
	cf.Log(fmt.Sprintf("KAAB --CORS server running: %s", os.Getenv("ENVIRONMENT")))

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-Width", "Authorization", "Content-Type", "Accept", "Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "User-Token"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "READ", "POST", "CREATE", "UPDATE", "DELETE"},
	})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), c.Handler(server.Router.Router)); err != nil {
		panic(err)
	}
}
