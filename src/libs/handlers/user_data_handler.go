package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"net/http"
	"strconv"
	"time"
)

var read_actions = NewStringArray{[]string{"account", "profile", "permissions", "report", "security"}}
var create_actions = NewStringArray{[]string{"profile", "report"}}
var update_actions = NewStringArray{[]string{"account", "profile", "permissions"}}
var delete_actions = NewStringArray{[]string{"account", "profile"}}

func IsReadAction(t string) bool {
	return read_actions.Contains(t)
}
func IsCreateAction(t string) bool {
	return create_actions.Contains(t)
}
func IsUpdateAction(t string) bool {
	return update_actions.Contains(t)
}
func IsDeleteAction(t string) bool {
	return delete_actions.Contains(t)
}

type MUD = models.UserData
type AllowedFunc map[string]func(args ...any) interface{}

var AllowedReadActions = AllowedFunc{
	"account":     GetAccount,
	"profile":     GetProfile,
	"permissions": GetPermissions,
	"report":      GetReport,
	"security":    GetSecurity,
}

func GetProfile(args ...any) interface{} {
	user_info := args[0].(MUD)
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
		Uuid:                 user_info.Uuid,
	}
}

func GetPermissions(args ...any) interface{} {
	user_info := args[0].(MUD)
	return models.PermissionsConform{
		Permissions: user_info.Realm,
		UserRol:     user_info.UserRol,
		Token:       user_info.Token,
	}
}

func GetReport(args ...any) interface{} {
	return models.ReportConform{
		ReportFrame: fmt.Sprintf("%v", time.Now().Unix()),
	}
}

func GetSecurity(args ...any) interface{} {
	user_info := args[0].(MUD)
	expirationDate, _ := strconv.ParseInt(user_info.Account.ExpirationDate, 10, 64)
	exp := time.Unix(expirationDate, 0)
	lifetime := time.Until(exp)
	return models.SecurityConform{
		Password:  lifetime.String(),
		Monitored: user_info.Monitored,
		KnownHost: user_info.KnownHost,
		Uuid:      user_info.Uuid,
	}
}

func GetAccount(args ...any) interface{} {
	user_info := args[0].(MUD)
	return models.AccountConform{
		Account:     user_info.Account,
		Email:       user_info.Email,
		AccessToken: user_info.AccessToken,
		Uuid:        user_info.Uuid,
	}
}

var AllowedCreateActions = AllowedFunc{
	"profile": CreateProfile,
	"report":  CreateReport,
}

func CreateProfile(args ...any) interface{} {
	var Res interface{}

	user_info := args[0].(MUD)
	r := args[1].(*http.Request)
	var payload models.CreateUserRequestBody
	err := GetBody(r, &payload)
	if err != nil {
		config.Err(fmt.Sprintf("Error to GetBody: %v", err))
		return DATA_FAIL
	}
	Res, err = db.CreateUser(user_info, payload)
	if err != nil {
		config.Err(fmt.Sprintf("Error creating user: %v", err))
		return DATA_FAIL
	}
	return Res
}

func CreateReport(args ...any) interface{} {
	var Res interface{}
	return Res
}

var AllowedUpdateActions = AllowedFunc{
	"account":     UpdateUserAccount,
	"profile":     UpdateUserProfile,
	"permissions": UpdateUserPermissions,
}

func UpdateUserAccount(args ...any) interface{} {
	var Res interface{}
	return Res
}

func UpdateUserProfile(args ...any) interface{} {
	var Res interface{}

	user_info := args[0].(MUD)
	r := args[1].(*http.Request)
	var payload models.UpdateProfileRequestBody
	err := GetBody(r, &payload)
	if err != nil {
		config.Err(fmt.Sprintf("Error to GetBody: %v", err))
		return DATA_FAIL
	}
	Res, err = db.UpdateUserProfile(user_info, payload)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating profile: %v", err))
		return DATA_FAIL
	}
	return Res
}

func UpdateUserPermissions(args ...any) interface{} {
	var Res interface{}
	return Res
}
