package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"net/http"
)

var AllowedDataCreateActions = AllowedBodyDataFunc{
	"nodes": {
		"item":  CreateNodeItem,
		"items": CreateNodeItems,
	},
	"instance": {
		"item":  CreateInstanceItem,
		"items": CreateInstanceItems,
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
	"endpoint": {
		"item":  CreateEndpointItem,
		"items": CreateFailed,
	},
}

func CreateFailed(args ...any) (any, error) {
	return DATA_FAIL, nil
}

func CreateEndpointItem(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)

	var payload models.CreateEndpointRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	ctrlData := CreateCtrlFields(id)
	newReferenceID := db.RandomRefID()
	instData := func() models.DataEntryIdentity {
		if args[2] == nil {
			return models.DataEntryIdentity{
				Name:   payload.Name,
				RefId:  newReferenceID,
				Status: db.STATUS.Activate(),
				Id:     ctrlData.Uuid,
			}
		}
		return args[2].(models.DataEntryIdentity)
	}()
	endpointItem := models.EndpointItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            []interface{}{payload.Value},
		RefId:            newReferenceID,
		Uuid:             ctrlData.Uuid,
		Size:             int64(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
		Status:           db.STATUS.Activate(),
		ApiBase:          ReqApi,
	}
	err = db.CreateEndpointItem(endpointItem, instData, subjectId, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

/* nodes */
func CreateNodeItem(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	instData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	if payload.RefName == "" || payload.Schema == "" {
		return DATA_FAIL, fmt.Errorf("cannot create node without required fields [ref_name, schema]")
	}
	err = VerifyNodeReferenceName(instData.Name, subjectId, payload.RefName, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}

	err = VerifyNodeFileName(instData.Name, subjectId, payload.Name, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}

	err = VerifySchemaExist(instData.Name, subjectId, payload.Schema, ReqApi, payload.Value)
	if err != nil {
		return DATA_FAIL, err
	}

	err = VerifySchemaConform(instData.Name, subjectId, payload.Schema, ReqApi, payload.Value)
	if err != nil {
		DATA_FAIL["Message"] = err.Error()
		return DATA_FAIL, err
	}

	ctrlData := CreateCtrlFields(id)
	newReferenceID := db.RandomRefID()

	nodeItem := models.NodeFileItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            payload.Value,
		RefId:            newReferenceID,
		Schema:           payload.Schema,
		Uuid:             ctrlData.Uuid,
		Size:             int64(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
		Status:           db.STATUS.Activate(),
		ApiBase:          ReqApi,
		RefName:          payload.RefName,
		Thumb:            payload.Thumb,
	}
	err = db.CreateNodeItem(nodeItem, instData, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateNodeItems(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	instData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload []models.CreateNodeRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		if item.RefName == "" || item.Schema == "" {
			return DATA_FAIL, fmt.Errorf("cannot create node without required fields [ref_name, schema]")
		}
		newReferenceID := db.RandomRefID()
		ctrlData := CreateCtrlFields(id)
		nodeItem := models.NodeFileItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			RefId:            newReferenceID,
			RefName:          item.RefName,
			Schema:           item.Schema,
			Uuid:             ctrlData.Uuid,
			Size:             int64(len(fmt.Sprintf("%v", item.Value))),
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
			Status:           db.STATUS.Activate(),
			ApiBase:          ReqApi,
		}
		err = db.CreateNodeItem(nodeItem, instData, subjectId)
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
	id := args[0].(string)
	r := args[1].(*http.Request)
	instanceData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		config.Err(fmt.Sprintf("payload.error: %s", err.Error()))
		return DATA_FAIL, err
	}

	if payload.RefName == "" || payload.Schema == "" {
		return DATA_FAIL, fmt.Errorf("cannot create content without required fields [ref_name, schema]")
	}
	err = VerifyContentReferenceName(instanceData.Name, subjectId, payload.RefName, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}

	err = VerifyContentFileName(instanceData.Name, subjectId, payload.Name, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}

	err = VerifySchemaExist(instanceData.Name, subjectId, payload.Schema, ReqApi, payload.Value)
	if err != nil {
		return DATA_FAIL, err
	}

	ctrlData := CreateCtrlFields(id)
	newReferenceID := db.RandomRefID()

	contentItem := models.TextFileItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            payload.Value,
		RefId:            newReferenceID,
		RefName:          payload.RefName,
		Schema:           payload.Schema,
		Uuid:             ctrlData.Uuid,
		Size:             int64(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
		Status:           db.STATUS.Activate(),
		ApiBase:          ReqApi,
	}
	err = db.CreateContentItem(contentItem, instanceData, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateContentItems(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	instanceData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload []models.CreateContentRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		err = VerifyContentFileName(instanceData.Name, subjectId, item.Name, ReqApi)
		if err != nil {
			return DATA_FAIL, err
		}

		err = VerifySchemaExist(instanceData.Name, subjectId, item.Schema, ReqApi, item.Value)
		if err != nil {
			return DATA_FAIL, err
		}
		newReferenceID := db.RandomRefID()
		ctrlData := CreateCtrlFields(id)

		contentItem := models.TextFileItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			RefId:            newReferenceID,
			RefName:          item.RefName,
			Schema:           item.Schema,
			Uuid:             ctrlData.Uuid,
			Size:             int64(len(fmt.Sprintf("%v", item.Value))),
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
			Status:           db.STATUS.Activate(),
			ApiBase:          ReqApi,
		}
		err = db.CreateContentItem(contentItem, instanceData, subjectId)
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
	id := args[0].(string)
	r := args[1].(*http.Request)
	instanceData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload models.CreateMediaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	err = VerifyMediaFileName(instanceData.Name, subjectId, payload.Name, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}
	newReferenceID := db.RandomRefID()
	ctrlData := CreateCtrlFields(id)
	instData := func() models.DataEntryIdentity {
		if args[2] == nil {
			return models.DataEntryIdentity{
				Name:   payload.Name,
				RefId:  newReferenceID,
				Status: db.STATUS.Activate(),
				Id:     ctrlData.Uuid,
			}
		}
		return args[2].(models.DataEntryIdentity)
	}()
	mediaAddress := CreateMediaCtrlFields(payload.ChallengeID, newReferenceID)
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
		Value:            []interface{}{mediaAddress},
		RefId:            newReferenceID,
		Ttype:            payload.Ttype,
		Duration:         payload.Duration,
		Dimensions:       payload.Dimensions,
		Service:          payload.Service,
		Thumb:            mediaAddress.Thumb,
		Url:              mediaAddress.Url,
		UriAddress:       mediaAddress.UriAddress,
		File:             mediaAddress.File,
		Status:           "active",
		ApiBase:          ReqApi,
	}
	err = db.CreateMediaItem(mediaItem, instData, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateMediaItems(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload []models.CreateMediaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		newReferenceID := db.RandomRefID()
		ctrlData := CreateCtrlFields(id)
		instData := func() models.DataEntryIdentity {
			if args[2] == nil {
				return models.DataEntryIdentity{
					Name:   item.Name,
					RefId:  newReferenceID,
					Status: db.STATUS.Activate(),
					Id:     ctrlData.Uuid,
				}
			}
			return args[2].(models.DataEntryIdentity)
		}()
		mediaAddress := CreateMediaCtrlFields(item.ChallengeID, newReferenceID)
		mediaItem := models.MediaFileItem{
			Uuid:             ctrlData.Uuid,
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
			Name:             item.Name,
			Description:      item.Description,
			Size:             item.Size,
			Value:            []interface{}{mediaAddress},
			RefId:            newReferenceID,
			Ttype:            item.Ttype,
			Duration:         item.Duration,
			Dimensions:       item.Dimensions,
			Service:          item.Service,
			Thumb:            mediaAddress.Thumb,
			Url:              mediaAddress.Url,
			UriAddress:       mediaAddress.UriAddress,
			File:             mediaAddress.File,
			Status:           db.STATUS.Activate(),
			ApiBase:          ReqApi,
		}
		err = db.CreateMediaItem(mediaItem, instData, subjectId)
		if err != nil {
			return DATA_FAIL, err
		}
		R := DATA_SUCC
		R["item"] = ctrlData.Uuid
		RES = append(RES, R)
	}
	return RES, nil
}

/* schemas */
func CreateSchemaItem(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	instanceData := args[2].(models.DataEntryIdentity)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	err = VerifySchemaItemName(instanceData.Name, subjectId, payload.Name, ReqApi)
	if err != nil {
		return DATA_FAIL, err
	}
	newReferenceID := db.RandomRefID()
	ctrlData := CreateCtrlFields(id)
	instData := func() models.DataEntryIdentity {
		if args[2] == nil {
			return models.DataEntryIdentity{
				Name:   payload.Name,
				RefId:  newReferenceID,
				Status: db.STATUS.Activate(),
				Id:     ctrlData.Uuid,
			}
		}
		return args[2].(models.DataEntryIdentity)
	}()
	schemaItem := models.SchemaItem{
		Name:             payload.Name,
		Description:      payload.Description,
		Value:            []interface{}{payload.Value},
		Uuid:             ctrlData.Uuid,
		Size:             int64(len(fmt.Sprintf("%v", payload.Value))),
		Versions:         ctrlData.Versions,
		CreationDate:     ctrlData.CreationDate,
		ModificationDate: ctrlData.ModificationDate,
		ModifiedBy:       ctrlData.ModifiedBy,
		CreatedBy:        ctrlData.CreatedBy,
		Status:           db.STATUS.Activate(),
		ApiBase:          ReqApi,
		RefId:            newReferenceID,
	}
	err = db.CreateSchemaItem(schemaItem, instData, subjectId)
	if err != nil {
		return DATA_FAIL, err
	}
	R := DATA_SUCC
	R["item"] = ctrlData.Uuid
	return R, nil
}

func CreateSchemaItems(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	var payload []models.CreateSchemaRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	RES := []any{}
	for _, item := range payload {
		newReferenceID := db.RandomRefID()
		ctrlData := CreateCtrlFields(id)
		instData := func() models.DataEntryIdentity {
			if args[2] == nil {
				return models.DataEntryIdentity{
					Name:   item.Name,
					RefId:  newReferenceID,
					Status: db.STATUS.Activate(),
					Id:     ctrlData.Uuid,
				}
			}
			return args[2].(models.DataEntryIdentity)
		}()
		schemaItem := models.SchemaItem{
			Name:             item.Name,
			Description:      item.Description,
			Value:            item.Value,
			Uuid:             ctrlData.Uuid,
			Size:             int64(len(fmt.Sprintf("%v", item.Value))),
			Versions:         ctrlData.Versions,
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
			Status:           db.STATUS.Activate(),
			ApiBase:          ReqApi,
			RefId:            newReferenceID,
		}
		err = db.CreateSchemaItem(schemaItem, instData, subjectId)
		if err != nil {
			return DATA_FAIL, err
		}
		R := DATA_SUCC
		R["item"] = ctrlData.Uuid
		RES = append(RES, R)
	}
	return RES, nil
}

/* instance */
func CreateInstanceItem(args ...any) (any, error) {
	id := args[0].(string)
	r := args[1].(*http.Request)
	subjectId := args[3].(string)
	ReqApi := args[4].(string)
	instanceCreation := args[5].(bool)
	publishTarget := args[6].(string)
	var payload models.CreateInstanceRequest
	err := GetBody(r, &payload)
	if err != nil {
		return DATA_FAIL, err
	}
	if payload.Name == "" {
		return DATA_FAIL, fmt.Errorf("cannot create instance without a name")
	}
	ctrlData := CreateCtrlFields(id)
	newReferenceID := db.RandomRefID()
	instData := func() models.DataEntryIdentity {
		if args[2] == nil {
			return models.DataEntryIdentity{
				Name:   payload.Name,
				RefId:  newReferenceID,
				Status: db.STATUS.Activate(),
				Id:     ctrlData.Uuid,
			}
		}
		return args[2].(models.DataEntryIdentity)
	}()
	instanceItem := models.InstanceCollection{
		ApiBase:        ReqApi,
		Name:           payload.Name,
		Versions:       ctrlData.Versions,
		Owner:          subjectId,
		Admin:          []string{subjectId},
		Members:        []string{subjectId},
		MediaFilesList: []models.DataEntryIdentity{},
		EndpointsList:  []models.DataEntryIdentity{},
		SchemasList:    []models.DataEntryIdentity{},
		TextFilesList:  []models.DataEntryIdentity{},
		NodesFilesList: []models.DataEntryIdentity{},
		Sys: models.SysData{
			CreationDate:     ctrlData.CreationDate,
			ModificationDate: ctrlData.ModificationDate,
			ModifiedBy:       ctrlData.ModifiedBy,
			CreatedBy:        ctrlData.CreatedBy,
			Status:           db.STATUS.Activate(),
		},
	}
	refID := ""
	if instanceCreation {
		rawRefID, err := db.CreateInstanceItem(instanceItem, models.DataEntryIdentity{Name: payload.Name}, subjectId, ReqApi, publishTarget, payload.Bump)
		refID = rawRefID.(string)
		if err != nil {
			return DATA_FAIL, err
		}
	} else {
		rawRefID, err := db.CreateInstanceItem(instanceItem, instData, subjectId, ReqApi, publishTarget, payload.Bump)
		refID = rawRefID.(string)
		if err != nil {
			return DATA_FAIL, err
		}
	}

	R := DATA_SUCC
	R["instance_id"] = ctrlData.Uuid
	R["instance_name"] = payload.Name
	R["instance_dev"] = refID
	return R, nil
}

func CreateInstanceItems(args ...any) (any, error) {
	return DATA_FAIL, nil
}
