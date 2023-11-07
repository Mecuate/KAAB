package handlers

import (
	"kaab/src/libs/utils"
	"kaab/src/models"
	"net/http"

	auth "github.com/Mecuate/auth_module"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	authorized, claims := auth.Authorized(w, r)
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
		resp := AllowedActions[action](user_info)
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

type AllowedFunc func(any models.UserData) interface{}

var AllowedActions = map[string]AllowedFunc{
	"account": GetAccount,
	// "profile":     true,
	// "permissions": true,
	// "report":      true,
	// "security":    true,
}

func GetAccount(user_info models.UserData) interface{} {
	acc := models.AccountConform{
		Account:     user_info.Account,
		Email:       user_info.Email,
		Id:          user_info.Id,
		AccessToken: user_info.AccessToken,
	}

	return acc
}
