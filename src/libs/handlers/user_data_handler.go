package handlers

import (
	"fmt"
	"kaab/src/libs/utils"
	"net/http"

	auth "github.com/Mecuate/auth_module"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	authorized, claims := auth.Authorized(w, r)
	fmt.Println("@@@ authorized", authorized)
	if authorized && claims.Realms.Read().Apis {
		params, err := ExtractPathParams(r, Params.USER)
		if err != nil {
			FailReq(w, 4)
			return
		}

		id, action := params["id"], params["action"]

		user_info, err := utils.PullUserData(id)
		if err != nil {
			FailReq(w, 5)
			return
		}
		fmt.Println(user_info)
		// TODO: [` add actions table `]-{2023-10-29}
		resp := map[string]interface{}{
			"tokenData": claims,
			"id":        id,
			"action":    action,
		}

		responseBody, err := JSON(resp)
		if err != nil {
			FailReq(w, 6)
			return
		}
		Response(w, responseBody)

	} else {
		http.Header.Add(w.Header(), "WWW-Authenticate", `JWT realm="Restricted"`)
		http.Header.Add(w.Header(), "User-Token", `SESSION`)
		http.Error(w, "", http.StatusUnauthorized)
	}
}

var AllowedActions = map[string]interface{}{
	"account":     true,
	"profile":     true,
	"permissions": true,
	"report":      true,
	"security":    true,
}
