package server

import (
	// "fmt"
	// "net/http"
	"fmt"
	"net/http"
	"strconv"

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
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	fmt.Fprintf(w, "The task with ID %v has been processed successfully", taskID)
}
func InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/boss/{id}", T).Methods(http.MethodGet)
	// sub := r.PathPrefix("/middleware").Subrouter()

	// sub.HandleFunc("/healthcheck", HealthCheckHandler).
	// 	Methods(http.MethodGet).
	// 	Name("healthcheck")

}
