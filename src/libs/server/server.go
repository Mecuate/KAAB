package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	cf "kaab/src/libs/config"
	"kaab/src/libs/db"
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

	server := NewServer(config)
	envs := config.WebServerConfig
	cf.Log(fmt.Sprintf("-- Server working ENVIRONMENT: %s", envs.Environment))

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
		AllowedHeaders: []string{
			"Authorization",
			"Cookie",
			"Content-Length",
			"Host",
			"User-Agent",
			"Accept",
			"Accept-Encoding",
			"Connection",
			"X-Requested-Width",
			"Content-Type",
			"Origin",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
			"User-Token"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "READ", "POST", "CREATE", "UPDATE", "DELETE"},
	})

	db.DatabaseSetup(serverConfig.PubDbName, strings.Split(serverConfig.ApiVersions, ","))
	cf.Log(fmt.Sprintf("DB initialization completed:[%s]", serverConfig.PubDbName))
	cf.Log(fmt.Sprintf("CORS server will run on port: %s", serverConfig.Port))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), c.Handler(server.Router.Router)); err != nil {
		panic(err)
	}
}
