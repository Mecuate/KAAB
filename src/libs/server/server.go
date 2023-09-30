package server

import (
	"fmt"
	"net/http"
	"os"

	cf "kaab/src/libs/config"
	"kaab/src/libs/handlers"
	"kaab/src/models"

	"github.com/rs/cors"
)

// func RunServer2(portNumber string) {
// 	router := NewRouter()

// 	// router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

// 	// handlers.StaticVideoHandler(router, "/media/")
// 	// handlers.StaticMediaHandler(router, "/static/")
// 	// handlers.StaticFormattedMedia(router, "/img/fmt/")
// 	// handlers.StaticFileHandler(router, "/")

// 	port := fmt.Sprintf(":%s", portNumber)

//		fmt.Printf("Server status [ON] \nlistening at port: %s \n", portNumber)
//		log.Fatal(http.ListenAndServe(port, router))
//	}
func NewServer(webServerConfig *models.ServiceConfig) *models.Server {
	server := &models.Server{
		Configuration: webServerConfig,
		Router:        handlers.NewRouter(),
	}

	return server
}

func RunServer() (err error) {

	config, err := cf.FromEnv()
	if err != nil {
		return err
	}

	if os.Getenv("ENVIROMENT") == "" {
		os.Setenv("ENVIROMENT", "development")
	}

	server := NewServer(config)
	envs := config.WebServerConfig

	if envs.CorsEnabled {
		CORSServer(config, server)
	} else {
		NormalServer(config, server)
	}

	return nil
}

func NormalServer(config *models.ServiceConfig, server *models.Server) {
	serverConfig := config.WebServerConfig
	fmt.Println("KAAB --Normal server running", os.Getenv("ENVIROMENT"))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), server.Router.Router); err != nil {
		panic(err)
	}
}

func CORSServer(config *models.ServiceConfig, server *models.Server) {
	serverConfig := config.WebServerConfig
	fmt.Println("KAAB --CORS server running", os.Getenv("ENVIROMENT"))
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-Width", "Authorization", "Content-Type", "Accept", "Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PATCH", "CREATE", "DELETE", "PUT", "UPDATE", "READ", "PUT"},
	})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), c.Handler(server.Router.Router)); err != nil {
		panic(err)
	}
}
