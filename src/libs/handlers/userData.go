package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// taskID, err := strconv.Atoi(vars["id"])
	taskID, err := vars["id"]

	if !err {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	fmt.Fprintf(w, "The task with ID %v has been processed successfully", taskID)
}
