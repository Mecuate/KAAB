package server

import (
	"kaab/src/libs/handlers"
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
	typedRouter := handlers.StabilizeRouter(r.Router)

	crud.CreateSingleHandlerCRUD(typedRouter, v1.userPath(), handlers.UserDataSimpleHandler)
	crud.CreateSingleHandlerCRUD(typedRouter, v1.emulatedAPIPath(), handlers.EmulatedAPISimpleHandler)
	r.Router.HandleFunc(RoutesDirectory.LOGIN, handlers.LoginHandler)
	r.Router.HandleFunc(RoutesDirectory.LOGOUT, handlers.LogoutHandler)
}
