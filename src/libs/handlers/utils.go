package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

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

func ExtractPathParams(r *http.Request, params []string) (map[string]string, bool) {
	vars := mux.Vars(r)

	newParams := make(map[string]string)

	for _, v := range params {
		rex := regexp.MustCompile(`[^A-Za-z0-9]`)
		query := rex.ReplaceAllString(vars[v], ``)
		if query == "" {
			return nil, true
		}
		newParams[v] = query
	}

	return newParams, false
}

func JSON(vals interface{}) string {
	res, err := json.Marshal(vals)
	if err != nil {
		return "[]"
	}
	return string(res)
}

func Response(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func HasAstrophytumCredentials(r *http.Request) bool {
	authHeader := r.Header.Get("Access-Control-Astrophytum-Credentials")
	return authHeader != ""
}
