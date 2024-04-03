package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"strings"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

type APIRelationship struct {
	Source  string
	Publish string
}

var KAAB_VERSION = map[string]APIRelationship{}

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
	var published_vers = strings.Split(config.WEBENV.ApiPublishedVersions, ",")
	for i := 0; i < len(vers); i++ {
		var api_version = APIVersion{vers[i]}
		/* CRUD */
		UserDataCRUD(typedRouter.Router, api_version.userPath())
		DataEntryCRUD(typedRouter.Router, api_version.dataEntryPath())
		/* single handler */
		crud.CreateSingleHandlerCRUD(typedRouter, api_version.emulatedAPIPath(), EmulatedAPISimpleHandler)
	}
	for i := 0; i < len(published_vers); i++ {
		var published_api_version = APIVersion{published_vers[i]}
		/* CRUD */
		PublicDataCRUD(typedRouter.Router, published_api_version.dataEntryPath())
		/* single handler */
		crud.CreateSingleHandlerCRUD(typedRouter, published_api_version.emulatedAPIPath(), EmulatedAPISimpleHandler)
	}
	AssingVersionedPaths(vers, published_vers)
}

func AssingVersionedPaths(vers []string, published_vers []string) {
	if len(vers) == len(published_vers) {
		for i := 0; i < len(vers); i++ {
			if vers[i] != "" {
				KAAB_VERSION[vers[i]] = APIRelationship{Source: vers[i], Publish: published_vers[i]}
			}
		}
		config.Log(fmt.Sprintf("API versions and published versions: %v", KAAB_VERSION))
	} else {
		panic("API versions and published versions are not equal")
	}
}

func NR(r *mux.Router) crud.MuxRouter {
	router := crud.MuxRouter{
		Router: r,
	}
	return router
}
