package handlers

import (
	"fmt"
	"kaab/src/models"
	"reflect"
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

func VerifySchemaItemName(instanceName string, subjectId string, name string, ReqApi string) error {
	SchemaData := GetSchemaList(instanceName, subjectId, "", "", ReqApi)
	SchemaList := SchemaData.(models.SchemasCollectionList)
	for _, item := range SchemaList {
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

func VerifyNodeFileName(instanceName string, subjectId string, name string, ReqApi string) error {
	Data := GetNodeList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.NodesFilesCollectionList)
	for _, item := range List {
		if item.Name == name {
			return fmt.Errorf("error name already in use")
		}
	}
	return nil
}

func VerifyNodeReferenceName(instanceName string, subjectId string, name string, ReqApi string) error {
	Data := GetNodeList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.NodesFilesCollectionList)
	for _, item := range List {
		if item.RefName == name {
			return fmt.Errorf("error ref_name already in use")
		}
	}
	return nil
}

func VerifyContentReferenceName(instanceName string, subjectId string, name string, ReqApi string) error {
	Data := GetContentList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.TextFilesCollectionList)
	for _, item := range List {
		if item.RefName == name {
			return fmt.Errorf("error ref_name already in use")
		}
	}
	return nil
}

func VerifySchemaConform(instanceName string, subjectId string, schema string, ReqApi string, data []interface{}) error {
	Data := GetSchemaList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.SchemasCollectionList)
	selected := models.SchemaItemResponse{}
	for _, item := range List {
		if item.Name == schema {
			schemaData := GetSchemaItem(instanceName, subjectId, item.Id, models.URLFilterSearchParams{}, ReqApi)
			selected = schemaData.(models.SchemaItemResponse)
		}
	}
	if ok := selected.Value.(models.MAPDATA); ok != nil {
		OK := selected.Value.(models.MAPDATA)

		for _, item := range data {
			if _, oki := item.(models.MAPDATA); !oki {
				return fmt.Errorf("wrong type in payload")
			}

			for k, v := range item.(models.MAPDATA) {
				if k == "$__i" {
					continue
				}
				if part := OK[k]; part != nil {
					testType := fmt.Sprintf("%v", reflect.TypeOf(v))
					switch testType {
					case "string":
						if part == "string" && len(fmt.Sprintf("%v", v)) < 96 {
							continue
						}
						if part == "long_string" {
							continue
						}
					case "bool":
						if part == "bool" {
							continue
						}
					case "float64":
						if part == "number" {
							continue
						}
					case "[]interface {}":
						if part == "array" {
							continue
						}
					case "map[string]interface {}":
						if part == "object" {
							continue
						}
					default:
						return fmt.Errorf("data does not conform to schema")
					}
					return fmt.Errorf("data does not conform to schema")
				} else {
					return fmt.Errorf("key-value pair out of range")
				}
			}
		}
	} else {
		return fmt.Errorf("data does not conform to schema")
	}
	return nil
}

func VerifySchemaExist(instanceName string, subjectId string, name string, ReqApi string, data []interface{}) error {
	Data := GetSchemaList(instanceName, subjectId, "", "", ReqApi)
	List := Data.(models.SchemasCollectionList)
	for _, item := range List {
		if item.Name == name {
			return nil
		}
	}
	return fmt.Errorf("error no schema by that reference found")
}
