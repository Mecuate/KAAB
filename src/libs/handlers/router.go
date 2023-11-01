package handlers

import (
	"kaab/src/models"

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
	var v1 = APIVersion{"collibri"}
	typedRouter := StabilizeRouter(r.Router)

	crud.CreateSingleHandlerCRUD(typedRouter, v1.userPath(), UserDataSimpleHandler)
	crud.CreateSingleHandlerCRUD(typedRouter, v1.emulatedAPIPath(), EmulatedAPISimpleHandler)
	crud.CreateSingleHandlerCRUD(typedRouter, v1.dataEntryPath(), DataEntryHandler)
}
