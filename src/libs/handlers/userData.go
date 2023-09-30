package handlers

import (
	"fmt"
	"net/http"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	if HasAuthHeader(r) {
		token, valid := VerifyToken(w, r)

		if valid {
			params, err := ExtractPathParams(r, Params.USER)
			if err {
				FailReq(w, 1)
			}

			id, action := params["id"], params["action"]

			fmt.Fprintf(w, "Request is valid with ID: [%s] ACTION: [%s] TOKEN: [%s]", id, action, token)
		} else {
			FailedToken(w, 0)
		}
	} else {
		http.Error(w, "", http.StatusUnauthorized)
	}
}
