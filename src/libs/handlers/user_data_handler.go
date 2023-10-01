package handlers

import (
	"net/http"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	if HasAuthHeader(r) && HasAstrophytumCredentials(r) {
		token, valid := VerifyToken(w, r)

		if valid {
			params, err := ExtractPathParams(r, Params.USER)
			if err {
				FailReq(w, 1)
			}

			id, action := params["id"], params["action"]

			resp := map[string]interface{}{
				"token":  token,
				"id":     id,
				"action": action,
				"data": map[string]interface{}{
					"data": map[string]interface{}{
						"email": "test@test.com",
					},
					"name": "test",
				},
			}

			responseBody := JSON(resp)
			Response(w, responseBody)

		} else {
			FailedToken(w, 0)
		}
	} else {
		http.Header.Add(w.Header(), "WWW-Authenticate", `JWT realm="Restricted"`)
		http.Header.Add(w.Header(), "Access-Control-Astrophytum-Credentials", `SESSION`)
		http.Error(w, "", http.StatusUnauthorized)
	}
}
