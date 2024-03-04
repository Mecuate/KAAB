package handlers

import (
	"fmt"
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
	return DATA_FAIL
}

func UpdateEndpointItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])

	var payload models.CreateEndpointRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateEndpointItem(payload, instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* nodes */
func UpdateNodeItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])

	var payload models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateNodeItem(payload, instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* content */
func UpdateContentItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])

	var payload models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateContentItem(payload, instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* media */
func UpdateMediaItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])

	var payload models.CreateMediaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateMediaItem(payload, instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* schemas */
func UpdateSchemaItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])

	var payload models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateSchemaItem(payload, instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	return R
}

/* instance */
func UpdateInstanceItem(args ...any) any {
	r := args[0].(*http.Request)
	instanceName := fmt.Sprintf("%v", args[1])
	subjectId := fmt.Sprintf("%v", args[2])
	itemId := fmt.Sprintf("%v", args[3])
	reqApi := fmt.Sprintf("%v", args[4])

	var payload models.CreateInstanceRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL
	}
	R, err := db.UpdateInstanceItem(payload, instanceName, subjectId, itemId, reqApi)
	if err != nil {
		return DATA_FAIL
	}
	return R
}
