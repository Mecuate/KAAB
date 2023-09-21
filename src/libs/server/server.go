package server

import (
	"fmt"
	"log"
	"net/http"

	"kaab/src/libs/handlers"

	"github.com/gorilla/mux"
)

func RunServer(portNumber string) {
	router := mux.NewRouter()

	// router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

	handlers.StaticVideoHandler(router, "/media/")
	handlers.StaticMediaHandler(router, "/static/")
	handlers.StaticFormattedMedia(router, "/img/fmt/")
	handlers.StaticFileHandler(router, "/")

	port := fmt.Sprintf(":%s", portNumber)

	fmt.Printf("Server status [ON] \nlistening at port: %s \n", portNumber)
	log.Fatal(http.ListenAndServe(port, router))
}
