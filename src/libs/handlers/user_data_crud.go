package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"net/http"

	auth "github.com/Mecuate/auth_module"
	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

func UserDataCRUD(r *mux.Router, path string) {
	var userDataHandlersCollection = crud.IndividualCRUDHandlers{
		crud.READ:   UserDataHandler_READ(path),
		crud.CREATE: UserDataHandler_CREATE(path),
		crud.UPDATE: UserDataHandler_UPDATE(path),
		crud.DELETE: UserDataHandler_DELETE(path),
	}
	crud.CreateMultiHandlerCRUD(NR(r), path, userDataHandlersCollection)
}

func UserDataHandler_READ(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)

		if authorized && claims.Realms.Read().Apis {
			id := claims.Id
			params, err := ExtractPathParams(r, Params.USER)
			if err != nil {
				FailReq(w, 4)
				return
			}
			instanceId, action := params["instance_id"], params["action"]
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			instanceInternalId, err := db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if IsReadAction(action) {
				user_info, err := db.PullUserData(id, instanceInternalId)
				if err != nil {
					FailReq(w, 5)
					return
				}
				fmt.Println("user_info: ", user_info)
				resp := AllowedReadActions[action](user_info)
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
}

func UserDataHandler_CREATE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Create().Apis {
			params, err := ExtractPathParams(r, Params.USER)
			if err != nil {
				FailReq(w, 4)
				return
			}
			instanceId, action := params["instance_id"], params["action"]
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			instanceInternalId, _ := db.VerifyInstanceExist(instanceId, ReqApi)

			if IsCreateAction(action) {
				var user_info = models.UserData{}
				if instanceId == ReqApi {
					user_info.Uuid = fmt.Sprintf("g-%s", claims.Id)
				} else {
					user_info, err = db.PullUserData(instanceId, instanceInternalId)
					if err != nil {
						FailReq(w, 5)
						return
					}
				}
				resp := AllowedCreateActions[action](user_info, r)
				responseBody, err := JSON(resp)
				if err != nil {
					FailReq(w, 6)
					return
				}
				Response(w, fmt.Sprintf(`{"CREATE": %s}`, responseBody))
			} else {
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}

func UserDataHandler_UPDATE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Update().Apis {
			params, err := ExtractPathParams(r, Params.USER)
			if err != nil {
				FailReq(w, 4)
				return
			}
			instanceId, action, subject := params["instance_id"], params["action"], params["subject_id"]
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			instanceInternalId, err := db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if IsUpdateAction(action) {
				user_info, err := db.PullUserData(instanceId, instanceInternalId)
				if err != nil {
					FailReq(w, 5)
					return
				}

				resp := AllowedUpdateActions[action](user_info, r, subject)
				responseBody, err := JSON(resp)
				if err != nil {
					FailReq(w, 6)
					return
				}
				Response(w, fmt.Sprintf(`{"update": %s}`, responseBody))
			} else {
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}

func UserDataHandler_DELETE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Delete().Apis {
			params, err := ExtractPathParams(r, Params.USER)
			if err != nil {
				FailReq(w, 4)
				return
			}
			instanceId, action, subject := params["instance_id"], params["action"], params["subject_id"]
			ReqApi, rerr := getReqApi(r)
			if rerr != nil {
				FailReq(w, 7)
				return
			}
			instanceInternalId, err := db.VerifyInstanceExist(instanceId, ReqApi)
			if err != nil {
				config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
				FailReq(w, 5)
				return
			}

			if IsUpdateAction(action) {
				user_info, err := db.PullUserData(instanceId, instanceInternalId)
				if err != nil {
					FailReq(w, 5)
					return
				}

				resp := AllowedDeleteActions[action](user_info, r, subject)
				responseBody, err := JSON(resp)
				if err != nil {
					FailReq(w, 6)
					return
				}
				Response(w, fmt.Sprintf(`{"delete": %s}`, responseBody))
			} else {
				FailReq(w, 99)
			}
		} else {
			RequestAuth(w)
		}
	}
}
