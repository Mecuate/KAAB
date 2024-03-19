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
	typedRouter := StabilizeRouter(r.Router)
	var vers = strings.Split(config.WEBENV.ApiVersions, ",")
	for i := 0; i < len(vers); i++ {
		var api_version = APIVersion{vers[i]}
		/* CRUD */
		UserDataCRUD(typedRouter.Router, api_version.userPath())
		DataEntryCRUD(typedRouter.Router, api_version.dataEntryPath())
		/* single handler */
		crud.CreateSingleHandlerCRUD(typedRouter, api_version.emulatedAPIPath(), EmulatedAPISimpleHandler)
	}
}

func NR(r *mux.Router) crud.MuxRouter {
	router := crud.MuxRouter{
		Router: r,
	}
	return router
}
