package handlers

import (
	"net/http"
	// "os"
	// "strconv"
	// "strings"

	"github.com/gorilla/mux"
)

func StaticFileHandler(r *mux.Router, path string) {
	redirPath := "public/"

	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(redirPath))))
}
