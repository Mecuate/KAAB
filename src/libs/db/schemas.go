package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSchemaItem(uuid string) (models.SchemaItem, error) {
	var res models.SchemaItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, SCHEMAS)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": uuid}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreateSchemaItem(data models.SchemaItem, instData models.DataEntryIdentity, subjectId string) error {
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
		Status: STATUS.Activate(),
		RefId:  data.RefId,
	}
	err = AddNewSchemasList(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Schema List: %v", err))
	}
	return nil
}

func DeleteSchemaItem(ref_id string) (models.SingleIDMap, error) {
	var R models.SingleIDMap
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

func UpdateSchemaItem(data models.CreateSchemaRequest, instData models.DataEntryIdentity, subjectId string, itemId string, ReqApi string, publishApiTarget string) (interface{}, error) {
	var R models.SingleIDMap
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
	if val := data.Description; val != "" {
		update["$set"].(bson.M)["description"] = val
	}
	if val := data.Status; val != "" && STATUS.Contains(val) {
		update["$set"].(bson.M)["status"] = val
	}
	if val := data.Value; len(val) == 1 {
		update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, []interface{}{val})
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
		Id:    itemId,
		Name:  recordDocument.Name,
		RefId: recordDocument.RefId,
		Status: func() string {
			if val := data.Status; val != "" && STATUS.Contains(val) {
				return val
			}
			return recordDocument.Status
		}(),
	}
	err = UpdateSchemaListItem(instData.Name, subjectId, newRecord, false)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Node List: %v", err))
	}

	if data.Bump {
		publishResponse, err := PublishTarget(SCHEMAS, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord, publishApiTarget)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Schema List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}

		pubDocument := publishResponse.Meta.(models.SchemaItem)

		updeateRecord := models.DataEntryIdentity{
			Id:     pubDocument.Uuid,
			Name:   pubDocument.Name,
			Status: pubDocument.Status,
			RefId:  pubDocument.RefId,
			Thumb:  pubDocument.Thumb,
		}
		err = UpdateSchemaListItem(instData.RefId, subjectId, updeateRecord, data.Bump)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Media List for published item: %v", err))
		}
	}

	return map[string]any{
		"id":        itemId,
		"ref":       recordDocument.RefId,
		"operation": updateRes != nil,
		"published": data.Bump,
	}, nil
}
