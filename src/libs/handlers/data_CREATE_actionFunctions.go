package handlers

import (
	"fmt"
	// "kaab/src/models"
)

var AllowedDataCreateActions = AllowedDataFunc{
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
func CreateNodeItem(args ...any) any {
	fmt.Println("CreateNodeItem", args[0], args[1])
	return EMPTY_OBJECT
}
func CreateNodeItems(args ...any) any {
	fmt.Println("CreateNodeItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* content */
func CreateContentItem(args ...any) any {
	fmt.Println("CreateContentItem", args[0], args[1])
	return EMPTY_OBJECT
}
func CreateContentItems(args ...any) any {
	fmt.Println("CreateContentItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* dynamic */
func CreateDynamicItem(args ...any) any {
	fmt.Println("CreateDynamicItem", args[0], args[1])
	return EMPTY_OBJECT
}
func CreateDynamicItems(args ...any) any {
	fmt.Println("CreateDynamicItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* media */
func CreateMediaItem(args ...any) any {
	fmt.Println("CreateMediaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func CreateMediaItems(args ...any) any {
	fmt.Println("CreateMediaItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* schemas */
func CreateSchemaItem(args ...any) any {
	fmt.Println("CreateSchemaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func CreateSchemaItems(args ...any) any {
	fmt.Println("CreateSchemaItems", args[0], args[1])
	return EMPTY_ARRAY
}
