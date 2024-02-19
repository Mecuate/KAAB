package handlers

type any = interface{}
type ArgsObject map[string]string
type AllowedDataFunc map[string]map[string]func(...any) interface{}

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
