package handlers

import (
	"kaab/src/libs/utils"
	"kaab/src/models"
	"net/http"
	"time"

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

		if IsReadAction(action) && user_info.Id == id {
			resp := AllowedGetActions[action](user_info)
			responseBody, err := JSON(resp)
			if err != nil {
				FailReq(w, 6)
				return
			}
			Response(w, responseBody)
		} else {
			FailReq(w, 99)
		}
	} else {
		RequestAuth(w)
	}
}

var get_actions = NewStringArray{[]string{"account", "profile", "permissions", "report", "security"}}

func IsReadAction(t string) bool {
	return get_actions.Contains(t)
}

type MUD = models.UserData
type AllowedFunc map[string]func(any MUD) interface{}

var AllowedGetActions = AllowedFunc{
	"account":     GetAccount,
	"profile":     GetProfile,
	"permissions": GetPermissions,
	"report":      GetReport,
	"security":    GetSecurity,
}

func GetProfile(user_info MUD) interface{} {
	return models.ProfileConform{
		Name:                 user_info.Name,
		LastName:             user_info.LastName,
		Nick:                 user_info.Nick,
		UserRol:              user_info.UserRol,
		LastLogin:            user_info.LastLogin,
		Modification_date:    user_info.Modification_date,
		Picture:              user_info.Account.Picture,
		PictureUrl:           user_info.Account.PictureUrl,
		PicModification_date: user_info.Account.Modification_date,
		ExpirationDate:       user_info.Account.ExpirationDate,
	}
}

func GetPermissions(user_info MUD) interface{} {
	return models.PermissionsConform{
		Permissions: user_info.Realm,
		UserRol:     user_info.UserRol,
		Token:       user_info.Token,
	}
}

func GetReport(user_info MUD) interface{} {
	// TODO: [` pull data from db to map trace user activity `]-{2023-11-06}
	return models.ReportConform{
		ReportFrame: "2023.01.15-2023.11.06",
	}
}

func GetSecurity(user_info MUD) interface{} {
	exp := time.Unix(user_info.Account.ExpirationDate, 0)
	lifetime := time.Until(exp)
	return models.SecurityConform{
		Password:  lifetime.String(),
		Monitored: user_info.Monitored,
		KnownHost: user_info.KnownHost,
	}
}

func GetAccount(user_info MUD) interface{} {
	return models.AccountConform{
		Account:     user_info.Account,
		Email:       user_info.Email,
		Id:          user_info.Id,
		AccessToken: user_info.AccessToken,
	}
}
