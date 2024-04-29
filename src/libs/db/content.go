package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyValue struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

func GetContentItem(ref_id string) (models.ContentItemResponse, error) {
	var res models.ContentItemResponse
	Db, err := InitMongoDB(config.WEBENV.PubDbName, FILES)
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

func CreateContentItem(data models.TextFileItem, instData models.DataEntryIdentity, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, FILES)
	if err != nil {
		return err
	}
	ctx := context.Background()

	data.Value = HandleContentCreationItems(CreateToMapArray(data.Value))
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Content Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: STATUS.Activate(),
		RefId:  data.RefId,
	}
	err = AddNewContentList(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Content List: %v", err))
	}
	return nil
}

func UpdateContentItem(data models.CreateContentRequest, instData models.DataEntryIdentity, subjectId string, itemId string, ReqApi string, publishApiTarget string) (interface{}, error) {
	var R models.SingleIDMap
	var recordDocument models.TextFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, FILES)
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
	if val := data.RefId; val != "" {
		update["$set"].(bson.M)["ref_id"] = val
	}
	if val := data.Schema; val != "" {
		update["$set"].(bson.M)["schema_ref"] = val
	}
	if val := data.Status; val != "" && STATUS.Contains(val) {
		update["$set"].(bson.M)["status"] = val
	}

	shouldUpdateVal := true
	currentIndex := func() int64 {
		r := IndexOf(recordDocument.Versions, data.Version)
		if r == -1 {
			return 0
		}
		return r
	}()
	valueItems := convertToMapArray(recordDocument.Value[currentIndex].(primitive.A))
	deletes := data.Deletes
	patches := data.Patches
	appends := data.Appends
	if len(deletes) > 0 {
		shouldUpdateVal = false
		for _, d := range deletes {
			for k, v := range valueItems {
				if v["$__i"] == d["$__i"] {
					DeleteItemFromArray(&valueItems, k)
				}
			}
		}
	}
	if len(patches) > 0 {
		shouldUpdateVal = false
		for _, p := range patches {
			for k, v := range valueItems {
				if v["$__i"] == p["$__i"] {
					UpdateItemFromArray(&valueItems, k, p)
				}
			}

		}
	}
	if len(appends) > 0 {
		shouldUpdateVal = false
		tl := len(valueItems)
		for p, v := range appends {
			AddItemFromArray(&valueItems, tl+p, v)
		}
	}
	if shouldUpdateVal {
		if val := data.Value; len(val) > 0 {
			updatedValues := HandleContentCreationItems(CreateToMapArray(val))
			update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, updatedValues)
		}
	} else {
		update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, []interface{}{valueItems})
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
		Thumb: recordDocument.Thumb,
		Status: func() string {
			if val := data.Status; val != "" && STATUS.Contains(val) {
				return val
			}
			return recordDocument.Status
		}(),
	}
	err = UpdateContentListItem(instData.Name, subjectId, newRecord, false)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Content List: %v", err))
	}

	if data.Bump {
		publishResponse, err := PublishTarget(FILES, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord, publishApiTarget)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Publish Content List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}

		pubDocument := publishResponse.Meta.(models.TextFileItem)

		updeateRecord := models.DataEntryIdentity{
			Id:     pubDocument.Uuid,
			Name:   pubDocument.Name,
			Status: pubDocument.Status,
			RefId:  pubDocument.RefId,
			Thumb:  pubDocument.Thumb,
		}
		err = UpdateContentListItem(instData.RefId, subjectId, updeateRecord, data.Bump)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Content List for published item: %v", err))
		}
	}

	return map[string]any{
		"id":        itemId,
		"ref":       recordDocument.RefId,
		"operation": updateRes != nil,
		"published": data.Bump,
	}, nil
}

func DeleteContentItem(ref_id string) (models.SingleIDMap, error) {
	var R models.SingleIDMap
	var res models.NodeFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, FILES)
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
