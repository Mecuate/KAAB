package handlers

import (
	"encoding/json"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"strings"
)

var AllowedDataReadActions = AllowedDataFunc{
	"nodes": {
		"list":  GetNodeList,
		"item":  GetNodeItem,
		"items": GetNodeItems,
	},
	"instance": {
		"list":  GetInstanceList,
		"item":  GetInstanceItem,
		"items": GetInstanceItems,
	},
	"content": {
		"list":  GetContentList,
		"item":  GetContentItem,
		"items": GetContentItems,
	},
	"media": {
		"list":  GetMediaList,
		"item":  GetMediaItem,
		"items": GetMediaItems,
	},
	"schemas": {
		"list":  GetSchemaList,
		"item":  GetSchemaItem,
		"items": GetSchemaItems,
	},
	"endpoint": {
		"list":  GetEndpointList,
		"item":  GetEndpointItem,
		"items": FailedGetItem,
	},
}

func FailedGetItem(args ...any) any {
	return DATA_FAIL
}

/* nodes */
func GetEndpointList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.EndpointsList
}

func GetEndpointItem(args ...any) any {
	endpointItem, err := db.GetEndpointItem(fmt.Sprintf("%v", args[2]))
	if err != nil {
		config.Err(fmt.Sprintf("Error getting endpointItem: %v", err))
		return DATA_FAIL
	}
	ReqSearch := args[3].(models.URLFilterSearchParams)
	return models.EndpointItemResponse{
		Uuid:        endpointItem.Uuid,
		Name:        endpointItem.Name,
		Description: endpointItem.Description,
		Size:        endpointItem.Size,
		Versions:    endpointItem.Versions,
		Value:       AssortData(endpointItem.Value, ReqSearch, endpointItem.Versions),
		RefId:       endpointItem.RefId,
		MemFile:     endpointItem.MemFile,
		Status:      endpointItem.Status,
	}
}

/* nodes */
func GetNodeList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.NodesFilesList
}

func GetNodeItem(args ...any) any {
	nodeItem, err := db.GetNodeItem(fmt.Sprintf("%v", args[2]))
	if err != nil {
		config.Err(fmt.Sprintf("Error getting nodeItem: %v", err))
		return DATA_FAIL
	}
	ReqSearch := args[3].(models.URLFilterSearchParams)
	return models.NodeItemResponse{
		Uuid:        nodeItem.Uuid,
		Name:        nodeItem.Name,
		Description: nodeItem.Description,
		Size:        nodeItem.Size,
		Versions:    nodeItem.Versions,
		Value:       AssortData(nodeItem.Value, ReqSearch, nodeItem.Versions),
		RefId:       nodeItem.RefId,
		Schema:      nodeItem.Schema,
		Status:      nodeItem.Status,
	}
}

func GetNodeItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[2]), "&")
	nodeItems := models.ManyNodeItemResponse{}
	for _, item := range items {
		nodeItem, err := db.GetNodeItem(item)
		if err != nil {
			config.Err(fmt.Sprintf("Error getting nodeItem: %v", err))
			return EMPTY_ARRAY
		}
		result := models.NodeItemResponse{
			Uuid:        nodeItem.Uuid,
			Name:        nodeItem.Name,
			Description: nodeItem.Description,
			Size:        nodeItem.Size,
			Versions:    nodeItem.Versions,
			Value:       nodeItem.Value[0:1],
			RefId:       nodeItem.RefId,
			Schema:      nodeItem.Schema,
			Status:      nodeItem.Status,
		}
		nodeItems = append(nodeItems, result)
	}
	return nodeItems
}

/* content */
func GetContentList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.TextFilesList
}

func GetContentItem(args ...any) any {
	contentItem, err := db.GetContentItem(fmt.Sprintf("%v", args[2]))
	if err != nil {
		config.Err(fmt.Sprintf("Error getting contentItem: %v", err))
		return DATA_FAIL
	}
	ReqSearch := args[3].(models.URLFilterSearchParams)
	return models.ContentItemResponse{
		Uuid:        contentItem.Uuid,
		Name:        contentItem.Name,
		Description: contentItem.Description,
		Size:        contentItem.Size,
		Versions:    contentItem.Versions,
		Value:       AssortData(contentItem.Value, ReqSearch, contentItem.Versions),
		RefId:       contentItem.RefId,
		Schema:      contentItem.Schema,
		Status:      contentItem.Status,
	}
}

func GetContentItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[2]), "&")
	contentItems := models.ManyContentItemResponse{}
	for _, item := range items {
		contentItem, err := db.GetContentItem(item)
		if err != nil {
			config.Err(fmt.Sprintf("Error getting contentItem: %v", err))
			return EMPTY_ARRAY
		}
		result := models.ContentItemResponse{
			Uuid:        contentItem.Uuid,
			Name:        contentItem.Name,
			Description: contentItem.Description,
			Size:        contentItem.Size,
			Versions:    contentItem.Versions,
			Value:       contentItem.Value[0:1],
			RefId:       contentItem.RefId,
			Schema:      contentItem.Schema,
			Status:      contentItem.Status,
		}
		contentItems = append(contentItems, result)
	}
	return contentItems
}

/* media */
func GetMediaList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.MediaFilesList
}

func GetMediaItem(args ...any) any {
	mediaItem, err := db.GetMediaItem(fmt.Sprintf("%v", args[2]))
	if err != nil {
		config.Err(fmt.Sprintf("Error getting mediaItem: %v", err))
		return DATA_FAIL
	}
	ReqSearch := args[3].(models.URLFilterSearchParams)
	return models.MediaItemResponse{
		Uuid:        mediaItem.Uuid,
		Name:        mediaItem.Name,
		Description: mediaItem.Description,
		Size:        mediaItem.Size,
		Versions:    mediaItem.Versions,
		Value:       AssortData(mediaItem.Value, ReqSearch, mediaItem.Versions),
		RefId:       mediaItem.RefId,
		Ttype:       mediaItem.Ttype,
		Duration:    mediaItem.Duration,
		Dimensions:  mediaItem.Dimensions,
		Service:     mediaItem.Service,
		Thumb:       mediaItem.Thumb,
		Url:         mediaItem.Url,
		UriAddress:  mediaItem.UriAddress,
		File:        mediaItem.File,
		Status:      mediaItem.Status,
	}
}

func GetMediaItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[2]), "&")
	mediaItems := models.ManyMediaItemResponse{}
	for _, item := range items {
		result := GetMediaItem("", "", item, args[3])
		if result == nil {
			config.Err(fmt.Sprintf("Error getting mediaItem: %v", item))
			return EMPTY_ARRAY
		}
		mediaItems = append(mediaItems, result.(models.MediaItemResponse))
	}
	return mediaItems
}

/* schemas */
func GetSchemaList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.SchemasList
}

func GetSchemaItem(args ...any) any {
	schemaItem, err := db.GetSchemaItem(fmt.Sprintf("%v", args[2]))
	if err != nil {
		config.Err(fmt.Sprintf("Error getting schemaItem: %v", err))
		return DATA_FAIL
	}
	ReqSearch := args[3].(models.URLFilterSearchParams)
	return models.SchemaItemResponse{
		Uuid:        schemaItem.Uuid,
		Name:        schemaItem.Name,
		Description: schemaItem.Description,
		Size:        schemaItem.Size,
		Versions:    schemaItem.Versions,
		Value:       AssortData(schemaItem.Value, ReqSearch, schemaItem.Versions),
		Status:      schemaItem.Status,
	}
}

func GetSchemaItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[2]), "&")
	schemaItems := models.ManySchemaItemResponse{}
	for _, item := range items {
		schemaItem, err := db.GetSchemaItem(item)
		if err != nil {
			config.Err(fmt.Sprintf("Error getting schemaItem: %v", err))
			return EMPTY_ARRAY
		}
		result := models.SchemaItemResponse{
			Uuid:        schemaItem.Uuid,
			Name:        schemaItem.Name,
			Description: schemaItem.Description,
			Size:        schemaItem.Size,
			Versions:    schemaItem.Versions,
			Value:       schemaItem.Value[0:1],
			Status:      schemaItem.Status,
		}
		schemaItems = append(schemaItems, result)
	}
	return schemaItems
}

/* instance */
func GetInstanceList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.Sys
}

func GetInstanceItem(args ...any) any {
	selected := fmt.Sprintf("%v", args[2])
	instance := GetInstanceList(args[0], args[1])
	if instance == nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", args[0]))
		return EMPTY_ARRAY
	}
	res, err := json.Marshal(instance)
	if err != nil {
		return DATA_FAIL
	}
	resp := models.SysData{}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return DATA_FAIL
	}
	switch selected {
	case "creation_date":
		return resp.CreationDate
	case "modification_date":
		return resp.ModificationDate
	case "created_by":
		return resp.CreatedBy
	case "modified_by":
		return resp.ModifiedBy
	case "status":
		return resp.Status
	}
	return DATA_FAIL
}

func GetInstanceItems(args ...any) any {
	return GetInstanceList(args[0], args[1])
}
