package handlers

import (
	"fmt"
	"kaab/src/models"
)

type any = interface{}
type ArgsObject map[string]string
type AllowedDataFunc map[string]map[string]func(...any) interface{}
type AllowedBodyDataFunc map[string]map[string]func(...any) (any, error)

var DATA_FAIL = map[string]any{"status": "failure", "result": false}
var DATA_SUCC = map[string]any{"status": "success", "result": true}
var EMPTY_ARRAY = []any{}
var EMPTY_OBJECT = any(map[string]any{})

func validSection(section string) bool {
	return data_sections.Contains(section)
}

func validDataAction(action string, reqType string, section string) bool {
	switch reqType {
	case "READ":
		return data_action_read.Contains(action) && validSection(section)
	case "UPDATE":
		return data_action_update.Contains(action) && validSection(section)
	case "DELETE":
		return data_action_delete.Contains(action) && validSection(section)
	case "CREATE":
		return data_action_create.Contains(action) && validSection(section)
	}

	return false
}

func VerifyMediaFileName(instanceName string, subjectId string, name string, ReqApi string) error {
	MediaData := GetMediaList(instanceName, subjectId, "", "", ReqApi)
	MedList := MediaData.(models.MediaFilesCollectionList)
	for _, item := range MedList {
		if item.Name == name {
			return fmt.Errorf("error name already in use")
		}
	}
	return nil
}

func VerifyContentFileName(instanceName string, subjectId string, name string, ReqApi string) error {
	Data := GetContentList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.TextFilesCollectionList)
	for _, item := range List {
		if item.Name == name {
			return fmt.Errorf("error name already in use")
		}
	}
	return nil
}
