package handlers

import (
	"fmt"
	// "kaab/src/models"
)

var AllowedDataUpdateActions = AllowedDataFunc{
	"nodes": {
		"item":  UpdateNodeItem,
		"items": UpdateNodeItems,
	},
	"dynamic": {
		"item":  UpdateDynamicItem,
		"items": UpdateDynamicItems,
	},
	"content": {
		"item":  UpdateContentItem,
		"items": UpdateContentItems,
	},
	"media": {
		"item":  UpdateMediaItem,
		"items": UpdateMediaItems,
	},
	"schemas": {
		"item":  UpdateSchemaItem,
		"items": UpdateSchemaItems,
	},
}

/* nodes */
func UpdateNodeItem(args ...any) any {
	fmt.Println("UpdateNodeItem", args[0], args[1])
	return EMPTY_OBJECT
}
func UpdateNodeItems(args ...any) any {
	fmt.Println("UpdateNodeItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* content */
func UpdateContentItem(args ...any) any {
	fmt.Println("UpdateContentItem", args[0], args[1])
	return EMPTY_OBJECT
}
func UpdateContentItems(args ...any) any {
	fmt.Println("UpdateContentItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* dynamic */
func UpdateDynamicItem(args ...any) any {
	fmt.Println("UpdateDynamicItem", args[0], args[1])
	return EMPTY_OBJECT
}
func UpdateDynamicItems(args ...any) any {
	fmt.Println("UpdateDynamicItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* media */
func UpdateMediaItem(args ...any) any {
	fmt.Println("UpdateMediaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func UpdateMediaItems(args ...any) any {
	fmt.Println("UpdateMediaItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* schemas */
func UpdateSchemaItem(args ...any) any {
	fmt.Println("UpdateSchemaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func UpdateSchemaItems(args ...any) any {
	fmt.Println("UpdateSchemaItems", args[0], args[1])
	return EMPTY_ARRAY
}