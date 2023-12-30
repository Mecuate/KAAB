package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kaab/src/libs/utils"
	"kaab/src/models"
	"net/http"
	"regexp"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

func LoadEndpointData(instance_id string, endpoint_name string) (models.EndpointInstance, error) {
	instanceId, err := utils.PullInstanceCollection(instance_id)
	if err != nil {
		return models.EndpointInstance{}, err
	}

	endpointInstance, err := utils.PullEndpoint(endpoint_name, instanceId)
	if err != nil {
		return models.EndpointInstance{}, err
	}

	return endpointInstance, nil
}

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

func ExtractPathParams(r *http.Request, params []string) (map[string]string, error) {
	mustMatch := NewStringArray{params}
	vars := mux.Vars(r)
	newParams := make(map[string]string)
	for _, v := range params {
		rex := regexp.MustCompile(`[^A-Za-z0-9-]`)
		query := rex.ReplaceAllString(vars[v], ``)
		if query == "" {
			return nil, errors.New("empty query")
		}
		newParams[v] = query
	}
	if len(newParams) == len(mustMatch.elements) {
		if mustMatch.hasAllOf(MapToStringSlice(vars)) {
			return newParams, nil
		}
	}
	return nil, errors.New("missing params")
}

func JSON(vals interface{}) (string, error) {
	res, err := json.Marshal(vals)
	if err != nil {
		return "[]", err
	}
	return string(res), nil
}

func Response(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func emptyResponse(w http.ResponseWriter) {
	fmt.Fprint(w, "[]")
}
