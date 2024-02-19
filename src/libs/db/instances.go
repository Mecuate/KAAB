package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveInstance(databaseName string, instanceName string) (string, error) {
	DB, err := InitMongoDB(databaseName, "instanceInfo")
	if err != nil {
		config.Err(fmt.Sprintf("Error building Instance [%s] Initial Data: %v", instanceName, err))
		return "", err
	}
	InternalRegistryData := bson.M{"collection_name": instanceName}

	exist := DB.FindOne(InternalRegistryData)

	if exist != nil {
		config.Err(fmt.Sprintf("Error building Instance [%s] already exist: %v", instanceName, err))
		return "", err
	}

	instanceID := uuid.New().String()
	newInstanceData := models.InstanceCollection{
		Name:           instanceName,
		Uuid:           instanceID,
		Owner:          "",
		Members:        []string{""},
		Admin:          []string{""},
		EndpointsList:  models.EndpointsCollectionList{},
		SchemasList:    models.SchemasCollectionList{},
		TextFilesList:  models.TextFilesCollectionList{},
		MediaFilesList: models.MediaFilesCollectionList{},
	}
	err = DB.InsertOne(newInstanceData)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceName, err))
		return "", err
	}
	return instanceID, err
}

func VerifyInstanceExist(instanceName string, apiName string) (string, error) {
	Db, err := InitMongoDB(config.WEBENV.IntDbName, apiName)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName}
	var res models.InstanceIdentData
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return "", err
	}
	return res.Id, nil
}

func GetInstanceInfo(instanceName string, subjectId string) (models.InstanceCollection, error) {
	var res models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"collection_name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// fmt.Println("@@@", res)
// instance, err := PullInstanceCollection(instance_id)
// members := NewStringArray{instance.Members}
// allow := members.Contains(user_id)
// if !allow {
// 	return false, errors.New("user not allowed")
// }
