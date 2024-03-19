package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSchemaItem(ref_id string) (models.SchemaItemResponse, error) {
	var res models.SchemaItemResponse
	Db, err := InitMongoDB(config.WEBENV.PubDbName, SCHEMAS)
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

func CreateSchemaItem(data models.SchemaItem, instName string, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, SCHEMAS)
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Schema Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: "active",
		RefId:  "",
	}
	err = AddNewSchemasList(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Schema List: %v", err))
	}
	return nil
}

func DeleteSchemaItem(ref_id string) (models.Delition, error) {
	var R models.Delition
	var res models.SchemaItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, SCHEMAS)
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

func UpdateSchemaItem(data models.CreateSchemaRequest, instName string, subjectId string, itemId string) (interface{}, error) {
	var R models.Delition
	var recordDocument models.SchemaItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, SCHEMAS)
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
	if val := data.Status; val != "" {
		update["$set"].(bson.M)["status"] = val
	}
	if val := data.RefId; val != "" {
		update["$set"].(bson.M)["ref_id"] = val
	}
	if val := data.Value; len(val) > 0 {
		update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, val)
	}
	update["$set"].(bson.M)["versions"] = UpdateVersions(recordDocument.Versions, data.Bump)
	timeStamp := fmt.Sprintf("%v", time.Now().Unix())
	update["$set"].(bson.M)["modified_by"] = AppendModificationRecord(recordDocument.ModifiedBy, subjectId, timeStamp)
	update["$set"].(bson.M)["modification_date"] = timeStamp
	updateRes, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return R, err
	}
	newRecord := models.DataEntryIdentity{
		Id: itemId,
		Name: func() string {
			if val := data.Name; val != "" {
				return val
			}
			return recordDocument.Name
		}(),
		Status: func() string {
			if val := data.Status; val != "" {
				return val
			}
			return recordDocument.Status
		}(),
		RefId: func() string {
			if val := data.RefId; val != "" {
				return val
			}
			return recordDocument.RefId
		}(),
	}
	err = UpdateSchemaListItem(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Node List: %v", err))
	}

	return updateRes, nil
}
