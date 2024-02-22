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
	return models.NodeItemResponse{
		Uuid:        nodeItem.Uuid,
		Name:        nodeItem.Name,
		Description: nodeItem.Description,
		Size:        nodeItem.Size,
		Versions:    nodeItem.Versions,
		Value:       nodeItem.Value,
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
			Value:       nodeItem.Value,
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
		Value:       contentItem.Value,
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
			Value:       contentItem.Value,
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
	fmt.Println("GetMediaItem", args[0], args[1])
	return DATA_FAIL
}
func GetMediaItems(args ...any) any {
	fmt.Println("GetMediaItems", args[0], args[1])
	return EMPTY_ARRAY
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
	fmt.Println("GetSchemaItem", args[0], args[1])
	return DATA_FAIL
}
func GetSchemaItems(args ...any) any {
	fmt.Println("GetSchemaItems", args[0], args[1])
	return EMPTY_ARRAY
}
