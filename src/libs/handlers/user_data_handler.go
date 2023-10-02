package handlers

import (
	"fmt"
	"net/http"

	auth "github.com/Mecuate/auth_module"
)

func UserDataSimpleHandler(w http.ResponseWriter, r *http.Request) {
	authorized := auth.Authorized(w, r)
	if authorized && HasAstrophytumCredentials(r) {
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
			fmt.Println(responseBody)

			Response(w, TT)

		} else {
			FailedToken(w, 0)
		}
	} else {
		http.Header.Add(w.Header(), "WWW-Authenticate", `JWT realm="Restricted"`)
		http.Header.Add(w.Header(), "Access-Control-Astrophytum-Credentials", `SESSION`)
		http.Error(w, "", http.StatusUnauthorized)
	}
}

var TT = `{
	"ACCOUNT_DATA": {
		"email": "string(64)",
		"id": "uuid",
		"lastName": "string(48)",
		"name": "string(36)",
		"nick": "string(16)",
		"password": "uuid--PASSWORDS",
		"picture": {
			"creationDate": "Date",
			"modificationDate": "Date",
			"url": "string"
		}
	},
	"AUDIO": {
		"address": "string",
		"audioBitrate": "int32",
		"audioChannel": "const<int8>",
		"audioCodec": "const<int8>",
		"creationDate": "Date",
		"file": "string",
		"id": "uuid",
		"modificationDate": "Date",
		"size": [{ "int8": "int32" }],
		"url": "string"
	},
	"DATA_ENTRY_EVENT": {
		"creationDate": "Date",
		"editionTime": "Time",
		"eventType": "created|update|deletion",
		"file": "uuid--FILE",
		"id": "uuid",
		"payload": "string"
	},
	"FILE": {
		"creationDate": "Date",
		"deltas": "uuid--GIT_COMMIT",
		"owner": "uuid",
		"payload": "string",
		"type": "const<int8>"
	},
	"GIT_COMMIT": {
		"creationDate": "Date",
		"deletions": "int32",
		"hash": "string(64)",
		"id": "uuid",
		"parent": "uuid--GIT_COMMIT",
		"writes": "int32"
	},
	"KNOWN_HOSTS": {
		"creationDate": "Date",
		"id": "uuid",
		"lastVisit": "Date",
		"signature": "string(64)",
		"visits": "int32"
	},
	"MEDIA": {
		"address": "string",
		"creationDate": "Date",
		"file": "string",
		"id": "uuid",
		"modificationDate": "Date",
		"url": "string"
	},
	"PASSWORDS": {
		"creationDate": "Date",
		"expired": "bool",
		"id": "uuid",
		"value": "string"
	},
	"STATS": {
		"creationDate": "Date",
		"errors": "int32",
		"id": "uuid",
		"interrruptions": "int32",
		"modificationDate": "Date",
		"views": "int32"
	},
	"SUBTITLES": {
		"creationDate": "Date",
		"id": "uuid",
		"language": "const<int8>",
		"modificationDate": "Date",
		"payload": "string",
		"size": "int32",
		"url": "string"
	},
	"USER": {
		"account": "uuid--ACCOUNT_DATA",
		"data": {
			"approverOf": ["uuid--USER"],
			"approvedBy": ["uuid--USER"],
			"creationDate": "Date",
			"editions": "uuid--DATA_ENTRY_EVENT",
			"expirationDate": "Date",
			"knownHOST": "uuid--KNOWN_HOSTS",
			"lastLogin": "Date",
			"modificationDate": "Date",
			"monitored": "bool",
			"realm": ["const<int8>"],
			"token": "uuid--JWT",
			"userRol": "int8"
		},
		"foreignId": "short_hash(8)",
		"id": "uuid"
	},
	"VIDEO": {
		"address": "string",
		"audio": "const<int8>",
		"audioTack": "uuid--AUDIO",
		"colorSpace": "const<int8>",
		"constraint": ["const<int8>"],
		"creationDate": "Date",
		"duration": "int32",
		"editionTime": "Time",
		"file": "string",
		"id": "uuid",
		"modificationDate": "Date",
		"resolutions": ["int8"],
		"size": [{ "int8": "int32" }],
		"stats": "uuid--STATS",
		"status": "const<int8>",
		"subtitles": "uuid--SUBTITLES",
		"url": "string",
		"videoBitrate": "int32",
		"videoCodec": "const<int8>"
	},
	"Z_SESSION": "int32"
}
`
