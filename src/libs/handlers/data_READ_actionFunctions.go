package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	// "kaab/src/models"
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

// user_info, err := db.PullUserData(userId, validInstanceId)
// if err != nil {
// 	FailReq(w, 5)
// 	return
// }

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
	return models.NodeFileItemResponse{
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
	fmt.Println("GetNodeItems", args[0], args[1])
	return EMPTY_ARRAY
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
	fmt.Println("GetContentItem", args[0], args[1])
	return DATA_FAIL
}
func GetContentItems(args ...any) any {
	fmt.Println("GetContentItems", args[0], args[1])
	return EMPTY_ARRAY
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
	fmt.Println("GetDynamicItem", args[0], args[1])
	return DATA_FAIL
}
func GetDynamicItems(args ...any) any {
	fmt.Println("GetDynamicItems", args[0], args[1])
	return EMPTY_ARRAY
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
