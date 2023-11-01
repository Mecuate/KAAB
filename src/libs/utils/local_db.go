package utils

import (
	"errors"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
)

func PullUserData(user_id string) (models.UserData, error) {
	users_list := models.UserDataDB{}
	userFile := OpenFILE{fmt.Sprintf("%s/%s", config.APPENV.DbDir, "users.json"), &users_list}
	userFile.JSON()
	for _, usr := range users_list {
		if usr.Id == user_id {
			return usr, nil
		}
	}
	return models.UserData{}, errors.New("not found")
}

func PullInstanceCollection(instance_id string, endpoint_name string) (models.EndpointInstance, error) {
	dataInstance := models.EndpointInstance{}
	file_name := fmt.Sprintf("%s/%s.json", config.APPENV.DbDir, instance_id)
	current_instance := OpenFILE{file_name, &dataInstance}
	current_instance.JSON()

	return models.EndpointInstance{}, errors.New("not found")
}
