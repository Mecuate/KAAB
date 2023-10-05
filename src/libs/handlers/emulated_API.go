package handlers

import (
	"fmt"
	"kaab/src/libs/fixtures"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertkrimen/otto"
)

func EmulatedAPISimpleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err1 := vars["id"]
	file_name, err2 := vars["file"]

	if !err1 || !err2 {
		emptyResponse(w)
		return
	}

	var user_context map[string]string
	var enp string

	ud := fixtures.TablesData

	for k, v := range ud {
		ox := ud[k]

		if ox.FileName == file_name && k == user_id {
			user_context = v.Data
			if r.Method == "READ" {
				enp = v.EndPoint[0]
			} else {
				enp = v.EndPoint[0]
			}

			break
		}
	}

	data, err_ := fixtures.ComposeEndpointJS(user_context, enp)

	if err_ != nil {
		emptyResponse(w)
		return
	}

	vm := otto.New()
	vm.Run(data)

	if value, err := vm.Get("response"); err == nil {
		fmt.Println(value)
		responseBody := JSON(value)
		Response(w, responseBody)
		return
	}

	emptyResponse(w)
}
