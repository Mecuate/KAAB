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
		Size:             int16(len(fmt.Sprintf("%v", payload.Value))),
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
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateNodeItems(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload []models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		ctrlData := CreateCtrlFields(id)
		nodeItem := models.NodeFileItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			RefId:            item.RefId,
			Schema:           item.Schema,
			Uuid:             ctrlData.Uuid,
			Size:             int16(len(fmt.Sprintf("%v", item.Value))),
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
		R := DATA_SUCC
		R["item"] = ctrlData.Uuid
		RES = append(RES, R)
	}
	return RES, nil
}

/* content */
func CreateContentItem(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	ctrlData := CreateCtrlFields(id)
	nodeItem := models.TextFileItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            payload.Value,
		RefId:            payload.RefId,
		Schema:           payload.Schema,
		Uuid:             ctrlData.Uuid,
		Size:             int16(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
	}
	err = db.CreateContentItem(nodeItem, instName, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}
func CreateContentItems(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload []models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		ctrlData := CreateCtrlFields(id)
		nodeItem := models.TextFileItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			RefId:            item.RefId,
			Schema:           item.Schema,
			Uuid:             ctrlData.Uuid,
			Size:             int16(len(fmt.Sprintf("%v", item.Value))),
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
		}
		err = db.CreateContentItem(nodeItem, instName, subjectId)
		if err != nil {
			return DATA_FAIL, err
		}
		R := DATA_SUCC
		R["item"] = ctrlData.Uuid
		RES = append(RES, R)
	}
	return RES, nil
}

/* media */
func CreateMediaItem(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload models.CreateMediaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	ctrlData := CreateCtrlFields(id)
	mediaAddress := CreateMediaCtrlFields(payload.RefId)
	mediaItem := models.MediaFileItem{
		Uuid:             ctrlData.Uuid,
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
		Name:             payload.Name,
		Description:      payload.Description,
		Size:             payload.Size,
		Value:            payload.Value,
		RefId:            payload.RefId,
		Ttype:            payload.Ttype,
		Duration:         payload.Duration,
		Dimensions:       payload.Dimensions,
		Service:          payload.Service,
		Thumb:            mediaAddress.Thumb,
		Url:              mediaAddress.Url,
		UriAddress:       mediaAddress.UriAddress,
		File:             mediaAddress.File,
	}
	err = db.CreateMediaItem(mediaItem, instName, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateMediaItems(args ...any) (any, error) {
	fmt.Println("CreateMediaItems", args[0], args[1])
	return EMPTY_ARRAY, nil
}

/* schemas */
func CreateSchemaItem(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	ctrlData := CreateCtrlFields(id)
	schemaItem := models.SchemaItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            payload.Value,
		Uuid:             ctrlData.Uuid,
		Size:             int16(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
	}
	err = db.CreateSchemaItem(schemaItem, instName, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateSchemaItems(args ...any) (any, error) {
	id := fmt.Sprintf("%v", args[0])
	instName := fmt.Sprintf("%v", args[2])
	subjectId := fmt.Sprintf("%v", args[3])
	r := args[1].(*http.Request)
	var payload []models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		ctrlData := CreateCtrlFields(id)
		schemaItem := models.SchemaItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			Uuid:             ctrlData.Uuid,
			Size:             int16(len(fmt.Sprintf("%v", item.Value))),
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
		}
		err = db.CreateSchemaItem(schemaItem, instName, subjectId)
		if err != nil {
			return DATA_FAIL, err
		}
		R := DATA_SUCC
		R["item"] = ctrlData.Uuid
		RES = append(RES, R)
	}
	return RES, nil
}

/* dynamic */
func CreateDynamicItem(args ...any) (any, error) {
	return EMPTY_OBJECT, nil
}
func CreateDynamicItems(args ...any) (any, error) {
	return EMPTY_ARRAY, nil
}
