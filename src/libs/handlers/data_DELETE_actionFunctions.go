package handlers

import (
	"fmt"
	"kaab/src/libs/db"
	"strings"
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
	itemId := fmt.Sprintf("%v", args[0])
	subjectId := fmt.Sprintf("%v", args[1])
	instanceName := fmt.Sprintf("%v", args[2])
	res, err := db.DeleteNodeItem(itemId)
	if err != nil {
		return DATA_FAIL
	}
	err = db.UnsetNodeList(instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	R := DATA_SUCC
	R["item"] = res.Id
	return R
}
func DeleteNodeItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[0]), "&")
	RES := []any{}
	for _, v := range items {
		R := DeleteNodeItem(v, args[1], args[2])
		RES = append(RES, R)
	}
	return RES
}

/* content */
func DeleteContentItem(args ...any) any {
	res, err := db.DeleteContentItem(fmt.Sprintf("%v", args[0]))
	if err != nil {
		return DATA_FAIL
	}
	itemId := fmt.Sprintf("%v", args[0])
	subjectId := fmt.Sprintf("%v", args[1])
	instanceName := fmt.Sprintf("%v", args[2])
	err = db.UnsetContentList(instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	R := DATA_SUCC
	R["item"] = res.Id
	return R
}
func DeleteContentItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[0]), "&")
	RES := []any{}
	for _, v := range items {
		R := DeleteContentItem(v, args[1], args[2])
		RES = append(RES, R)
	}
	return RES
}

/* media */
func DeleteMediaItem(args ...any) any {
	res, err := db.DeleteMediaItem(fmt.Sprintf("%v", args[0]))
	if err != nil {
		return DATA_FAIL
	}
	itemId := fmt.Sprintf("%v", args[0])
	subjectId := fmt.Sprintf("%v", args[1])
	instanceName := fmt.Sprintf("%v", args[2])
	err = db.UnsetMediaList(instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	R := DATA_SUCC
	R["item"] = res.Id
	return R
}
func DeleteMediaItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[0]), "&")
	RES := []any{}
	for _, v := range items {
		R := DeleteMediaItem(v, args[1], args[2])
		RES = append(RES, R)
	}
	return RES
}

/* schemas */
func DeleteSchemaItem(args ...any) any {
	res, err := db.DeleteSchemaItem(fmt.Sprintf("%v", args[0]))
	if err != nil {
		return DATA_FAIL
	}
	itemId := fmt.Sprintf("%v", args[0])
	subjectId := fmt.Sprintf("%v", args[1])
	instanceName := fmt.Sprintf("%v", args[2])
	err = db.UnsetSchemasList(instanceName, subjectId, itemId)
	if err != nil {
		return DATA_FAIL
	}
	R := DATA_SUCC
	R["item"] = res.Id
	return R
}
func DeleteSchemaItems(args ...any) any {
	items := strings.Split(fmt.Sprintf("%v", args[0]), "&")
	RES := []any{}
	for _, v := range items {
		R := DeleteSchemaItem(v, args[1], args[2])
		RES = append(RES, R)
	}
	return RES
}

/* dynamic */
func DeleteDynamicItem(args ...any) any {
	res, err := db.DeleteInstanceItem(fmt.Sprintf("%v", args[0]))
	if err != nil {
		return DATA_FAIL
	}
	R := DATA_SUCC
	R["item"] = res.Id
	return R
}
func DeleteDynamicItems(args ...any) any {
	return EMPTY_ARRAY
}
