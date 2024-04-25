package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"net/http"
)

var AllowedDataUpdateActions = AllowedDataFunc{
	"nodes": {
		"item":  UpdateNodeItem,
		"items": UpdateFailed,
	},
	"instance": {
		"item":  UpdateInstanceItem,
		"items": UpdateFailed,
	},
	"content": {
		"item":  UpdateContentItem,
		"items": UpdateFailed,
	},
	"media": {
		"item":  UpdateMediaItem,
		"items": UpdateFailed,
	},
	"schemas": {
		"item":  UpdateSchemaItem,
		"items": UpdateFailed,
	},
	"endpoint": {
		"item":  UpdateEndpointItem,
		"items": UpdateFailed,
	},
}

func UpdateFailed(args ...any) any {
	fmt.Println("UpdateFailed got called")
	return DATA_FAIL
}

func UpdateEndpointItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)
	publishApiTarget := args[5].(string)

	var payload models.CreateEndpointRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateEndpointItem(payload, instanceData, subjectId, itemId, ReqApi, publishApiTarget)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* nodes */
func UpdateNodeItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateNodeItem(payload, instanceData, subjectId, itemId, ReqApi)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* content */
func UpdateContentItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateContentItem(payload, instanceData, subjectId, itemId, ReqApi)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* media */
func UpdateMediaItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateMediaRequest
	err := GetBody(r, &payload)
	if err != nil {
		config.Err(fmt.Sprintf("payload.error: %s", err.Error()))
		return DATA_FAIL
	}
	mediaUpdateValue := models.InternalMediaCtrlFields{}
	if payload.ChallengeID != "" {
		mediaUpdateValue = CreateMediaCtrlFields(payload.ChallengeID, instanceData.Id)
	}
	R, err := db.UpdateMediaItem(payload, instanceData, subjectId, itemId, ReqApi, KAAB_VERSION[ReqApi].Publish, mediaUpdateValue)
	if err != nil {
		config.Err(fmt.Sprintf("db.error: %s", err.Error()))
		return DATA_FAIL
	}
	return R
}

/* schemas */
func UpdateSchemaItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateSchemaItem(payload, instanceData, subjectId, itemId, ReqApi)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* instance */
func UpdateInstanceItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceData := args[1].(models.DataEntryIdentity)
	subjectId := args[2].(string)
	itemId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateInstanceRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateInstanceItem(payload, instanceData, subjectId, itemId, ReqApi)
	if err != nil {
		return DATA_FAIL
	}
	return R
}
