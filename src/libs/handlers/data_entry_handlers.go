package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"net/http"

	auth "github.com/Mecuate/auth_module"
	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

var data_sections = NewStringArray{[]string{"nodes", "schemas", "content", "media", "endpoint", "instance"}}

var data_action_read = NewStringArray{[]string{"list", "item", "items"}}
var data_action_create = NewStringArray{[]string{"item", "items"}}
var data_action_update = NewStringArray{[]string{"item", "items"}}
var data_action_delete = NewStringArray{[]string{"item", "items"}}

func DataEntryCRUD(r *mux.Router, path string) {
	var DataHandlersCollection = crud.IndividualCRUDHandlers{
		crud.READ:   DataHandler_READ(path),
		crud.CREATE: DataHandler_CREATE(path),
		crud.UPDATE: DataHandler_UPDATE(path),
		crud.DELETE: DataHandler_DELETE(path),
	}
	crud.CreateMultiHandlerCRUD(NR(r), path, DataHandlersCollection)
}

func DataHandler_READ(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)

		if authorized && claims.Realms.Read().Apis {
			userId := claims.Id
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				config.Err(fmt.Sprintf("Error Extracting Path Params: %v", err))
				FailReq(w, 4)
				return
			}
			instanceId, section, action, ref_id := params["instance_id"], params["section"], params["action"], params["ref_id"]
			ReqSearch, err := GetRequestFSS(r.RequestURI)
			if err != nil {
				config.Err(fmt.Sprintf("Bad URL: %v [%s]", err, r.RequestURI))
				FailReq(w, 5)
				return
			}

			_, err = db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if validDataAction(action, r.Method, section) {
				resp := AllowedDataReadActions[section][action](instanceId, userId, ref_id, ReqSearch)
				responseBody, err := JSON(resp)
				if err != nil {
					config.Err(fmt.Sprintf("Error JSON: %v", err))
					FailReq(w, 6)
					return
				}
				Response(w, responseBody)
			} else {
				config.Err(fmt.Sprintf("Error Validating Data Action: Invalid Action [%s] or Section [%s]", action, section))
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}

func DataHandler_CREATE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)

		if authorized && claims.Realms.Create().Apis {
			userId := claims.Id
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				config.Err(fmt.Sprintf("Error Extracting Path Params: %v", err))
				FailReq(w, 4)
				return
			}
			instanceId, section, action, _ := params["instance_id"], params["section"], params["action"], params["ref_id"]
			_, err = db.VerifyInstanceExist(instanceId, ReqApi)
			// TODO: [` add double check for permissions:: internalInstanceID `]-{2024-02-23}
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}
			if validDataAction(action, r.Method, section) {
				resp, err := AllowedDataCreateActions[section][action](userId, r, instanceId, userId, ReqApi)
				if err != nil {
					config.Err(fmt.Sprintf("Error getting body: %v", err))
					FailReq(w, 4)
					return
				}
				responseBody, err := JSON(resp)
				if err != nil {
					config.Err(fmt.Sprintf("Error JSON: %v", err))
					FailReq(w, 6)
					return
				}
				Response(w, responseBody)
			} else {
				config.Err(fmt.Sprintf("Error Validating Data Action: Invalid Action [%s] or Section [%s]", action, section))
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}

func DataHandler_UPDATE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Update().Apis {
			userId := claims.Id
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				config.Err(fmt.Sprintf("Error Extracting Path Params: %v", err))
				FailReq(w, 4)
				return
			}
			instanceId, section, action, ref_id := params["instance_id"], params["section"], params["action"], params["ref_id"]
			internalInstanceID, err := db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if validDataAction(action, r.Method, section) {
				resp := AllowedDataUpdateActions[section][action](r, instanceId, userId, ref_id, internalInstanceID)
				responseBody, err := JSON(resp)
				if err != nil {
					config.Err(fmt.Sprintf("Error JSON: %v", err))
					FailReq(w, 6)
					return
				}
				Response(w, responseBody)
			} else {
				config.Err(fmt.Sprintf("Error Validating Data Action: Invalid Action [%s] or Section [%s]", action, section))
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}

func DataHandler_DELETE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Delete().Apis {
			userId := claims.Id
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				config.Err(fmt.Sprintf("Error Extracting Path Params: %v", err))
				FailReq(w, 4)
				return
			}
			instanceId, section, action, ref_id := params["instance_id"], params["section"], params["action"], params["ref_id"]
			internalInstanceID, err := db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if validDataAction(action, r.Method, section) {
				resp := AllowedDataDeleteActions[section][action](ref_id, userId, instanceId, internalInstanceID)
				responseBody, err := JSON(resp)
				if err != nil {
					config.Err(fmt.Sprintf("Error JSON: %v", err))
					FailReq(w, 6)
					return
				}
				Response(w, responseBody)
			} else {
				config.Err(fmt.Sprintf("Error Validating Data Action: Invalid Action [%s] or Section [%s]", action, section))
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}
