package handlers

import (
	"fmt"
	"kaab/src/libs/db"
	"kaab/src/libs/utils"
	"net/http"

	auth "github.com/Mecuate/auth_module"
	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

var data_sections = NewStringArray{[]string{"nodes", "dynamic_data", "document", "media", "system"}}

var data_action_read = NewStringArray{[]string{"list", "item", "items"}}
var data_action_create = NewStringArray{[]string{"item", "items"}}
var data_action_update = NewStringArray{[]string{"item", "items"}}
var data_action_delete = NewStringArray{[]string{"item", "items"}}

// func DataEntryHandler(w http.ResponseWriter, r *http.Request) {
// 	Method := r.Method
// 	ReqApi, rerr := getReqApi(r)
// 	if rerr != nil {
// 		FailReq(w, 7)
// 		return
// 	}
// 	params, err := ExtractPathParams(r, Params.DATA_ACTION)
// 	if err != nil {
// 		FailReq(w, 1)
// 		return
// 	}

// 	instanceId, section, action := params["instance_id"], params["section"], params["action"]
// 	fmt.Fprintf(w, "SECTION: [%v], ID: [%s], ACTION: [%v]", valid_section(section), instanceId, valid_action(action, Method))
// 	fmt.Fprintf(w, "%s", ReqApi)
// }

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
				FailReq(w, 4)
				return
			}
			instanceId, section, action := params["instance_id"], params["section"], params["action"]

			fmt.Println("realSection", validSection(section), !validSection(section))

			if !validSection(section) {
				fmt.Println("@@@ misstake")
				FailReq(w, 4)
				return
			}
			fmt.Println("@@@ pass 1")

			validInstanceId, err := db.VerifyInstanceExist(instanceId, ReqApi)
			fmt.Println("validInstanceId", validInstanceId, err)
			if err != nil {
				fmt.Println("@@@ mistaKE 2")
				FailReq(w, 5)
				return
			}
			fmt.Println("@@@ pass 2")
			fmt.Println("validInstanceId", validInstanceId)

			user_info, err := db.PullUserData(userId, validInstanceId)
			fmt.Println("user_info: ", user_info, err)
			if err != nil {
				FailReq(w, 5)
				return
			}

			if validAction(action, r.Method) && user_info.Id == userId {
				// resp := AllowedReadActions[action](user_info)
				// responseBody, err := JSON(resp)
				// if err != nil {
				// 	FailReq(w, 6)
				// 	return
				// }
				// fmt.Println("responseBody", responseBody)
				// Response(w, responseBody)
				rr, _ := JSON(user_info)
				Response(w, rr)
			} else {
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
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				FailReq(w, 4)
				return
			}
			id, action, section := params["subject_id"], params["action"], params["section"]
			realSection := validSection(section)
			if !realSection {
				FailReq(w, 4)
				return
			}
			user_info, err := utils.PullUserData(id)
			if err != nil {
				FailReq(w, 5)
				return
			}

			if IsReadAction(action) && user_info.Id == id {
				resp := AllowedReadActions[action](user_info)
				responseBody, err := JSON(resp)
				if err != nil {
					FailReq(w, 6)
					return
				}
				Response(w, fmt.Sprintf(`{"create": %s}`, responseBody))
			} else {
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
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				FailReq(w, 4)
				return
			}
			id, action := params["subject_id"], params["action"]
			user_info, err := utils.PullUserData(id)
			if err != nil {
				FailReq(w, 5)
				return
			}

			if IsReadAction(action) && user_info.Id == id {
				resp := AllowedReadActions[action](user_info)
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

func DataHandler_DELETE(path string) crud.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized, claims := auth.Authorized(w, r)
		if authorized && claims.Realms.Delete().Apis {
			params, err := ExtractPathParams(r, Params.DATA_ACTION)
			if err != nil {
				FailReq(w, 4)
				return
			}
			id, action := params["subject_id"], params["action"]
			user_info, err := utils.PullUserData(id)
			if err != nil {
				FailReq(w, 5)
				return
			}

			if IsReadAction(action) && user_info.Id == id {
				resp := AllowedReadActions[action](user_info)
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

func validSection(section string) bool {
	return data_sections.Contains(section)
}

func validAction(action string, reqType string) bool {
	switch reqType {
	case "READ":
		return data_action_read.Contains(action)
	case "UPDATE":
		return data_action_update.Contains(action)
	case "DELETE":
		return data_action_delete.Contains(action)
	case "CREATE":
		return data_action_create.Contains(action)

	}
	return false
}
