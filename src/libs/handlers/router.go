package handlers

import (
	"kaab/src/libs/config"
	"kaab/src/models"
	"strings"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

func NewRouter() models.MuxRouter {
	router := models.MuxRouter{
		Router: mux.NewRouter(),
	}
	InitializeRoutes(router)
	return router
}

func InitializeRoutes(r models.MuxRouter) {
	var vers = strings.Split(config.WEBENV.ApiVersions, ",")
	var v1 = APIVersion{vers[0]}
	typedRouter := StabilizeRouter(r.Router)

	// crud.CreateSingleHandlerCRUD(typedRouter, v1.userPath(), UserDataSimpleHandler)
	UserDataCRUD(typedRouter.Router, v1.userPath())
	crud.CreateSingleHandlerCRUD(typedRouter, v1.emulatedAPIPath(), EmulatedAPISimpleHandler)
	crud.CreateSingleHandlerCRUD(typedRouter, v1.dataEntryPath(), DataEntryHandler)
}

func NR(r *mux.Router) crud.MuxRouter {
	router := crud.MuxRouter{
		Router: r,
	}
	return router
}
