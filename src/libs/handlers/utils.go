package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

func StabilizeRouter(r *mux.Router) crud.MuxRouter {
	router := crud.MuxRouter{
		Router: r,
	}
	return router
}

func GetBody(r *http.Request, mo interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &mo)
	if err != nil {
		return err
	}

	return nil
}
