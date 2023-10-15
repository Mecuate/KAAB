package handlers

import (
	"net/http"

	auth "github.com/Mecuate/auth_module"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	authorized, claims := auth.Authorized(w, r)
	if authorized && claims.Realms.Read().Apis {
		params, err := ExtractPathParams(r, Params.USER)
		if err {
			FailReq(w, 1)
		}

		id, action := params["id"], params["action"]

		resp := map[string]interface{}{
			"tokenData": claims,
			"id":        id,
			"action":    action,
		}

		responseBody := JSON(resp)
		Response(w, responseBody)

	} else {
		http.Header.Add(w.Header(), "WWW-Authenticate", `JWT realm="Restricted"`)
		http.Header.Add(w.Header(), "User-Token", `SESSION`)
		http.Error(w, "", http.StatusUnauthorized)
	}
}
