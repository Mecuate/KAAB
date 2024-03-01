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
	InternalRegistryData := bson.M{"name": instanceName}

	exist := DB.FindOne(InternalRegistryData)

	if exist != nil {
		config.Err(fmt.Sprintf("Error building Instance [%s] already exist: %v", instanceName, err))
		return "", err
	}

	instanceID := uuid.New().String()
	newInstanceData := models.InstanceCollection{
		Name:           instanceName,
		Owner:          "",
		Members:        []string{""},
		Admin:          []string{""},
		MediaFilesList: models.MediaFilesCollectionList{},
		EndpointsList:  models.EndpointsCollectionList{},
		SchemasList:    models.SchemasCollectionList{},
		TextFilesList:  models.TextFilesCollectionList{},
		NodesFilesList: models.NodesFilesCollectionList{},
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
	var res models.DataEntryIdentity
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
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func UpdateNodeList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	res, err := Db.coll.UpdateOne(ctx, identify, bson.M{"$push": bson.M{"nodes_collection_list": data}})
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Node Items UPDATED: %v", res))
	return nil
}

func UnsetNodeList(instanceName string, subjectId string, itemId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{
		"$pull": bson.M{
			"nodes_collection_list": bson.M{"id": itemId},
		},
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Node Items UPDATED: %v", res))
	return nil
}

func UpdateMediaList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	res, err := Db.coll.UpdateOne(ctx, identify, bson.M{"$push": bson.M{"media_files_collection_list": data}})
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Media Items UPDATED: %v", res))
	return nil
}

func UnsetMediaList(instanceName string, subjectId string, itemId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{
		"$pull": bson.M{
			"media_files_collection_list": bson.M{"id": itemId},
		},
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of MEdia Items UPDATED: %v", res))
	return nil
}

func UpdateEndpointsList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	res, err := Db.coll.UpdateOne(ctx, identify, bson.M{"$push": bson.M{"endpoints_collection_list": data}})
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Endpoints Items UPDATED: %v", res))
	return nil
}

func UnsetEndpointList(instanceName string, subjectId string, itemId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{
		"$pull": bson.M{
			"endpoints_collection_list": bson.M{"id": itemId},
		},
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Endpoints Items UPDATED: %v", res))
	return nil
}

func UpdateSchemasList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	res, err := Db.coll.UpdateOne(ctx, identify, bson.M{"$push": bson.M{"schemas_collection_list": data}})
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Schemas Items UPDATED: %v", res))
	return nil
}

func UnsetSchemasList(instanceName string, subjectId string, itemId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{
		"$pull": bson.M{
			"schemas_collection_list": bson.M{"id": itemId},
		},
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Schemas Items UPDATED: %v", res))
	return nil
}

func UpdateContentList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	res, err := Db.coll.UpdateOne(ctx, identify, bson.M{"$push": bson.M{"files_collection_list": data}})
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Files Items UPDATED: %v", res))
	return nil
}

func UnsetContentList(instanceName string, subjectId string, itemId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{
		"$pull": bson.M{
			"files_collection_list": bson.M{"id": itemId},
		},
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("List of Content Items UPDATED: %v", res))
	return nil
}

func DeleteInstanceItem(ref_id string) (models.Delition, error) {
	var R models.Delition
	var res models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, NODES)
	if err != nil {
		return R, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": ref_id}
	err = Db.coll.FindOneAndDelete(ctx, identify).Decode(&res)
	if err != nil {
		return R, err
	}
	R.Id = ref_id
	return R, nil
}
