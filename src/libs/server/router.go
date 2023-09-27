package server

import (
	// "fmt"
	// "net/http"
	"fmt"
	"kaab/src/libs/handlers"
	"net/http"

	// "strconv"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

const (
	protobufContentType = "application/x-protobuf"
	jsonContentType     = "application/json"
	connScope           = "conn-pass"
)

var (
	routeAuthScopes = map[string]string{
		"healthcheck":      connScope,
		"protobuf_enable":  connScope,
		"protobuf_disable": connScope,
	}

	authExemptEndpoints = map[string]string{
		"healthcheck": "notDefined.s",
	}

	jwtClaimsOptIn = map[string]string{
		"postProfile": "user",
	}
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	router := mux.NewRouter()
	InitializeRoutes(router)
	return &Router{router}
}
func T(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// taskID, err := strconv.Atoi(vars["id"])
	taskID, err := vars["id"]

	if !err {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	fmt.Fprintf(w, "The task with ID %v has been processed successfully", taskID)
}

func P(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// taskID, err := strconv.Atoi(vars["id"])
	taskID, err := vars["id"]

	if !err {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	fmt.Fprintf(w, "The task with ID %v has been processed successfully", taskID)
}

func InitializeRoutes(r *mux.Router) {
	var v1 = APIVersion{"collibri"}
	typedRouter := handlers.StabilizeRouter(r)

	crud.CreateSingleHandlerCRUD(typedRouter, v1.userPath(), T)
	crud.CreateSingleHandlerCRUD(typedRouter, v1.emulatedAPIPath(), P)

}
