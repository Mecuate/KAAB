package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/libs/db"
	"kaab/src/models"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
)

func PullEndpoint(endpointName string, instance_data models.InstanceCollection) (models.EndpointInstance, error) {
	endpointData := models.EndpointInstance{}
	selectedEndpoint := models.DataEntryIdentity{}
	for _, endp := range instance_data.EndpointsList {
		if endp.Name == endpointName {
			selectedEndpoint = endp
			break
		}
	}
	Db, err := InitMongoDB(config.WEBENV.PubDbName, db.ENDPOINTS)
	if err != nil {
		return endpointData, err
	}
	endpointFile := models.EndpointFile{}
	ctx := context.Background()
	identify := bson.M{"name": selectedEndpoint.Name, "uuid": selectedEndpoint.Id}
	err = Db.coll.FindOne(ctx, identify).Decode(&endpointFile)
	if err != nil {
		return endpointData, err
	}
	if endpointFile.Value[0].Get != "" || endpointFile.Value[0].Post != "" || endpointFile.Value[0].Delete != "" {
		endpointData.EndpointCode = endpointFile.Value[0]
		dataContext := GatherContext(instance_data, endpointFile.Value[0])
		endpointData.Context = dataContext
		fmt.Println(":dataContext:", dataContext)
		return endpointData, nil
	}
	return endpointData, errors.New("instance does not exist")
}

func GatherContext(instance_data models.InstanceCollection, code models.EndpointCode) string {
	selection := NewStringArray{}
	rex := regexp.MustCompile(`useContext\([\s]["'](.*)["'][\s]\)`)
	allItems := rex.FindAllString(code.Get, -1)
	for i := 0; i < len(allItems); i++ {
		match := rex.FindStringSubmatch(allItems[i])
		if len(match) != 0 {
			selection.elements = append(selection.elements, match[1])
		}
	}
	internalCTX := map[string]any{}

	ctx := context.Background()
	Db, err := InitMongoDB(config.WEBENV.PubDbName, db.FILES)
	if err != nil {
		return ""
	}
	for ii := 0; ii < len(instance_data.TextFilesList); ii++ {
		cur := instance_data.TextFilesList[ii]
		if key, ok := selection.ContainsKey(cur.RefId); ok {
			currSelFile := models.TextFileItem{}
			identify := bson.M{"name": cur.Name, "uuid": cur.Id}
			err = Db.coll.FindOne(ctx, identify).Decode(&currSelFile)
			if err != nil {
				return ""
			}
			internalCTX[key] = currSelFile.Value[0]
		}
	}
	res, err := json.Marshal(internalCTX)
	if err != nil {
		config.Err(fmt.Sprintf("Error utils.JSON.Marshal: %v", err))
		return "{\"error\": true}"
	}
	return string(res)
}

func PullInstanceCollection(instance_id string) (models.InstanceCollection, error) {
	dataInstance, err := db.GetInstanceInfo(instance_id, "")
	if err != nil {
		return models.InstanceCollection{}, err
	}
	return dataInstance, err
}
