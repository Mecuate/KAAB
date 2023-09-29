package server

import (
	"fmt"
	"net/http"
	"os"

	cf "kaab/src/libs/config"
	"kaab/src/libs/handlers"
	"kaab/src/models"
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

	serverConfig := config.WebServerConfig
	server := NewServer(config)

	fmt.Println("KAAB server running")

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", "", serverConfig.Port), server.Router.Router); err != nil {
		panic(err)
	}

	return nil
}
