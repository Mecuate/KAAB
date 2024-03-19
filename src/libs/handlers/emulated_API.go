package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/utils"
	"net/http"

	"github.com/robertkrimen/otto"
)

func EmulatedAPISimpleHandler(w http.ResponseWriter, r *http.Request) {
	params, err := ExtractPathParams(r, Params.EMULATED_API)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.instEndpointObject: %v", err))
		FailReq(w, 1)
	}

	instance_id, endpoint_name := params["instance_id"], params["file"]
	instEndpointObject, err := LoadEndpointData(instance_id, endpoint_name)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.instEndpointObject: %v", err))
		emptyResponse(w)
		return
	}
	endpoint, err := utils.ComposeEndpointJS(instEndpointObject, r)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.ComposeEndpointJS: %v", err))
		emptyResponse(w)
		return
	}
	vm := otto.New()
	vm.Run(endpoint)

	if value, err := vm.Get("response"); err == nil {
		responseBody, errj := JSON(value)
		if errj != nil {
			config.Err(fmt.Sprintf("Error running otto: %v", err))
			emptyResponse(w)
			return
		}
		Response(w, responseBody)
		return
	}

	config.Err(fmt.Sprintf("Empty response from emulated api: %s", r.URL.Path))
	emptyResponse(w)
}
