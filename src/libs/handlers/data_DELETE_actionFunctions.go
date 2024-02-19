package handlers

import (
	"fmt"
	// "kaab/src/models"
)

var AllowedDataDeleteActions = AllowedDataFunc{
	"nodes": {
		"item":  DeleteNodeItem,
		"items": DeleteNodeItems,
	},
	"dynamic": {
		"item":  DeleteDynamicItem,
		"items": DeleteDynamicItems,
	},
	"content": {
		"item":  DeleteContentItem,
		"items": DeleteContentItems,
	},
	"media": {
		"item":  DeleteMediaItem,
		"items": DeleteMediaItems,
	},
	"schemas": {
		"item":  DeleteSchemaItem,
		"items": DeleteSchemaItems,
	},
}

/* nodes */
func DeleteNodeItem(args ...any) any {
	fmt.Println("DeleteNodeItem", args[0], args[1])
	return EMPTY_OBJECT
}
func DeleteNodeItems(args ...any) any {
	fmt.Println("DeleteNodeItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* content */
func DeleteContentItem(args ...any) any {
	fmt.Println("DeleteContentItem", args[0], args[1])
	return EMPTY_OBJECT
}
func DeleteContentItems(args ...any) any {
	fmt.Println("DeleteContentItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* dynamic */
func DeleteDynamicItem(args ...any) any {
	fmt.Println("DeleteDynamicItem", args[0], args[1])
	return EMPTY_OBJECT
}
func DeleteDynamicItems(args ...any) any {
	fmt.Println("DeleteDynamicItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* media */
func DeleteMediaItem(args ...any) any {
	fmt.Println("DeleteMediaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func DeleteMediaItems(args ...any) any {
	fmt.Println("DeleteMediaItems", args[0], args[1])
	return EMPTY_ARRAY
}

/* schemas */
func DeleteSchemaItem(args ...any) any {
	fmt.Println("DeleteSchemaItem", args[0], args[1])
	return EMPTY_OBJECT
}
func DeleteSchemaItems(args ...any) any {
	fmt.Println("DeleteSchemaItems", args[0], args[1])
	return EMPTY_ARRAY
}
