package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"regexp"
)

func PullUserData(user_id string) (models.UserData, error) {
	users_list := models.UserDataDB{}
	userFile := OpenFILE{fmt.Sprintf("%s/%s", config.APPENV.DbDir, "users.json"), &users_list}
	userFile.parseJSON()
	for _, usr := range users_list {
		if usr.Id == user_id {
			return usr, nil
		}
	}
	return models.UserData{}, errors.New("not found")
}

func PullEndpoint(endpoint_name string, instance_data models.InstanceCollection) (models.EndpointInstance, error) {
	endpointData := models.EndpointInstance{}
	selectedEndpoint := models.EndpointItem{}
	for _, endp := range instance_data.EndpointsList {
		if endp.Name == endpoint_name {
			selectedEndpoint = endp
			break
		}
	}
	endpointFile := models.EndpointFile{}
	dataFile := OpenFILE{fmt.Sprintf("%s/%s", config.APPENV.DbDir, selectedEndpoint.File), &endpointFile}
	dataFile.parseJSON()

	if endpointFile.Value.Generic != "" || endpointFile.Value.Get != "" || endpointFile.Value.Post != "" || endpointFile.Value.Delete != "" {
		endpointData.EndpointCode = endpointFile.Value
		endpointData.Context = GatherContext(instance_data, endpointFile.Value)
		return endpointData, nil
	}
	return endpointData, errors.New("instance does not exist")
}

func GatherContext(instance_data models.InstanceCollection, code models.EndpointCode) string {
	selection := NewStringArray{}
	rex := regexp.MustCompile(`useContext\(["'](.*)["']\)`)
	allItems := rex.FindAllString(code.Generic, -1)
	for i := 0; i < len(allItems); i++ {
		match := rex.FindStringSubmatch(allItems[i])
		if len(match) != 0 {
			selection.elements = append(selection.elements, match[1])
		}
	}
	ctx := map[string]any{}
	for ii := 0; ii < len(instance_data.TextFilesList); ii++ {
		cur := instance_data.TextFilesList[ii]

		if key, ok := selection.ContainsKey(cur.RefId); ok {
			currSelFile := models.DBstorageFile{}
			activeFile := OpenFILE{fmt.Sprintf("%s/%s", config.APPENV.DbDir, cur.File), &currSelFile}
			activeFile.parseJSON()
			ctx[key] = currSelFile.Value
		}
	}
	res, err := json.Marshal(ctx)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.JSON.Marshal: %v", err))
		return "{\"error\": true}"
	}
	return string(res)
}

func PullInstanceCollection(instance_id string) (models.InstanceCollection, error) {
	dataInstance := models.InstanceCollection{}
	fileName := fmt.Sprintf("%s/%s.json", config.APPENV.DbDir, instance_id)
	currentInstance := OpenFILE{fileName, &dataInstance}
	currentInstance.parseJSON()

	if dataInstance.Name == instance_id {
		return dataInstance, nil
	}
	return dataInstance, errors.New("instance does not exist")
}
