package db

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveIntanceRecord(intId string, instanceName string, apiName string) (string, error) {
	DB, err := InitMongoDB("apis_internal_registry", apiName)
	if err != nil {
		config.Err(fmt.Sprintf("Error with internal db records [%s]:%v", instanceName, apiName, err))
		return "", err
	}
	InstData := bson.M{"instances": bson.M{"name": instanceName}}

	exist := DB.FindOne(InstData)
	fmt.Println("internal record exist", exist)

	if exist != nil {
		fmt.Println("failed to create; instance already exist.", exist)
		config.Err(fmt.Sprintf("Error building Instance [%s] already exist: %v", instanceName, err))
		return "", err
	}
	instanceItem := []models.InstanceIdentData{
		{
			Name: instanceName,
			Id:   intId,
		},
	}
	newInstanceData := models.APICollections{
		Size:      1,
		Instances: instanceItem,
	}

	DB.InsertOne(newInstanceData)
	return "SUCCESFUL", nil
}
