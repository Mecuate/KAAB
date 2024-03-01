package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetNodeItem(ref_id string) (models.NodeFileItem, error) {
	var res models.NodeFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, NODES)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": ref_id}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreateNodeItem(data models.NodeFileItem, instName string, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, NODES)
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Node Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: "active",
		RefId:  data.RefId,
	}
	err = UpdateNodeList(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Node List: %v", err))
	}
	return nil
}

func DeleteNodeItem(ref_id string) (models.Delition, error) {
	var R models.Delition
	var res models.NodeFileItem
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

func UpdateNodeItem(data models.CreateNodeRequest, instName string, subjectId string, itemId string) (interface{}, error) {
	var R models.Delition
	var recordDocument models.NodeFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, NODES)
	if err != nil {
		return R, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": itemId}
	err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)
	if err != nil {
		return R, err
	}

	update := bson.M{
		"$set": bson.M{},
	}
	if val := data.Name; val != "" {
		update["$set"].(bson.M)["name"] = val
	}
	if val := data.Description; val != "" {
		update["$set"].(bson.M)["description"] = val
	}
	if val := data.RefId; val != "" {
		update["$set"].(bson.M)["ref_id"] = val
	}
	if val := data.Schema; val != "" {
		update["$set"].(bson.M)["schema_ref"] = val
	}
	if val := data.Value; len(val) > 0 {
		update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, val)
	}
	bump := data.Bump
	update["$set"].(bson.M)["versions"] = UpdateVersions(recordDocument.Versions, bump)

	updateRes, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return R, err
	}
	return updateRes, nil
}
