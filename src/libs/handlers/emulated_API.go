package handlers

import (
	"kaab/src/libs/utils"
	"net/http"

	"github.com/robertkrimen/otto"
)

func EmulatedAPISimpleHandler(w http.ResponseWriter, r *http.Request) {
	params, err := ExtractPathParams(r, Params.EMULATED_API)
	if err != nil {
		FailReq(w, 1)
	}

	instance_id, endpoint_name := params["id"], params["file"]

	instance, err := utils.PullInstanceCollection(instance_id, endpoint_name)
	if err != nil {
		emptyResponse(w)
		return
	}

	endpoint, err := utils.ComposeEndpointJS(instance)
	if err != nil {
		emptyResponse(w)
		return
	}

	vm := otto.New()
	vm.Run(endpoint)

	if value, err := vm.Get("response"); err == nil {
		responseBody, errj := JSON(value)
		if errj != nil {
			emptyResponse(w)
			return
		}
		Response(w, responseBody)
		return
	}

	emptyResponse(w)
}
