package server

import (
	"fmt"
	// "net/http"
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
	return &Router{router}
}

// func (r *Router) InitializeRoutes(envName string) {
// 	r.InitializeProfileRoutes(envName)

// 	sub := r.Router.PathPrefix("/middleware").Subrouter()

// 	sub.HandleFunc("/healthcheck", HealthCheckHandler).
// 		Methods(http.MethodGet).
// 		Name("healthcheck")

// 	// sub.HandleFunc("/conn/create", CreateContractHandler).
// 	// 	Methods(http.MethodPost).
// 	// 	Name("createContract")

// 	// sub.HandleFunc("/product/filter", GetFilteredProductsHandler).
// 	// 	Methods(http.MethodPost).
// 	// 	Name("getFilters")

// }

func InitializeProfileRoutes(envName string) string {

	var routeWorkaround string

	routeWorkaround = "middleware"

	if envName == "CaaS" {
		routeWorkaround = "ssasc"
	}

	return fmt.Sprintf("/%s/v1/", routeWorkaround)
}
