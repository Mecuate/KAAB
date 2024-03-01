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
	"dynamic": {
		"list":  GetDynamicList,
		"item":  GetDynamicItem,
		"items": GetDynamicItems,
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
	return models.ContentItemResponse{
		Uuid:        contentItem.Uuid,
		Name:        contentItem.Name,
		Description: contentItem.Description,
		Size:        contentItem.Size,
		Versions:    contentItem.Versions,
		Value:       contentItem.Value[0:1],
		RefId:       contentItem.RefId,
		Schema:      contentItem.Schema,
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
		}
		contentItems = append(contentItems, result)
	}
	return contentItems
}

/* dynamic */
func GetDynamicList(args ...any) any {
	instanceName, subjectId := fmt.Sprintf("%v", args[0]), fmt.Sprintf("%v", args[1])
	instance, err := db.GetInstanceInfo(instanceName, subjectId)
	if err != nil {
		config.Err(fmt.Sprintf("Error getting instance info: %v", err))
		return EMPTY_ARRAY
	}
	return instance.Sys
}

func GetDynamicItem(args ...any) any {
	selected := fmt.Sprintf("%v", args[2])
	instance := GetDynamicList(args[0], args[1])
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

func GetDynamicItems(args ...any) any {
	return GetDynamicList(args[0], args[1])
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
	return models.MediaItemResponse{
		Uuid:        mediaItem.Uuid,
		Name:        mediaItem.Name,
		Description: mediaItem.Description,
		Size:        mediaItem.Size,
		Versions:    mediaItem.Versions,
		Value:       mediaItem.Value[0:1],
		RefId:       mediaItem.RefId,
		Ttype:       mediaItem.Ttype,
		Duration:    mediaItem.Duration,
		Dimensions:  mediaItem.Dimensions,
		Service:     mediaItem.Service,
		Thumb:       mediaItem.Thumb,
		Url:         mediaItem.Url,
		UriAddress:  mediaItem.UriAddress,
		File:        mediaItem.File,
	}
}

func GetMediaItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[2]), "&")
	mediaItems := models.ManyMediaItemResponse{}
	for _, item := range items {
		result := GetMediaItem("", "", item)
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
	return models.SchemaItemResponse{
		Uuid:        schemaItem.Uuid,
		Name:        schemaItem.Name,
		Description: schemaItem.Description,
		Size:        schemaItem.Size,
		Versions:    schemaItem.Versions,
		Value:       schemaItem.Value[0:1],
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
		}
		schemaItems = append(schemaItems, result)
	}
	return schemaItems
}
