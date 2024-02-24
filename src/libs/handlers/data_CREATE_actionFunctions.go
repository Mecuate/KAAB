package handlers

import (
	"fmt"
	"kaab/src/libs/db"
	"kaab/src/models"
	"net/http"
)

var AllowedDataCreateActions = AllowedBodyDataFunc{
	"nodes": {
		"item":  CreateNodeItem,
		"items": CreateNodeItems,
	},
	"dynamic": {
		"item":  CreateDynamicItem,
		"items": CreateDynamicItems,
	},
	"content": {
		"item":  CreateContentItem,
		"items": CreateContentItems,
	},
	"media": {
		"item":  CreateMediaItem,
		"items": CreateMediaItems,
	},
	"schemas": {
		"item":  CreateSchemaItem,
		"items": CreateSchemaItems,
	},
}

/* nodes */
func CreateNodeItem(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	ctrlData := CreateCtrlFields(id)
	nodeItem := models.NodeFileItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            payload.Value,
		RefId:            payload.RefId,
		Schema:           payload.Schema,
		Uuid:             ctrlData.Uuid,
		Size:             ctrlData.Size,
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
	}
	err = db.CreateNodeItem(nodeItem, instName, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	return EMPTY_OBJECT, nil
}
func CreateNodeItems(args ...any) (any, error) {
	fmt.Println("CreateNodeItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}

/* content */
func CreateContentItem(args ...any) (any, error) {
	fmt.Println("CreateContentItem", args[0], args[1])
	return EMPTY_OBJECT, nil
}
func CreateContentItems(args ...any) (any, error) {
	fmt.Println("CreateContentItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}

/* dynamic */
func CreateDynamicItem(args ...any) (any, error) {
	fmt.Println("CreateDynamicItem", args[0], args[1])
	return EMPTY_OBJECT, nil
}
func CreateDynamicItems(args ...any) (any, error) {
	fmt.Println("CreateDynamicItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}

/* media */
func CreateMediaItem(args ...any) (any, error) {
	fmt.Println("CreateMediaItem", args[0], args[1])
	return EMPTY_OBJECT, nil
}
func CreateMediaItems(args ...any) (any, error) {
	fmt.Println("CreateMediaItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}

/* schemas */
func CreateSchemaItem(args ...any) (any, error) {
	fmt.Println("CreateSchemaItem", args[0], args[1])
	return EMPTY_OBJECT, nil
}
func CreateSchemaItems(args ...any) (any, error) {
	fmt.Println("CreateSchemaItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}
