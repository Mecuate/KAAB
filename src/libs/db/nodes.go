package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

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

func CreateNodeItem(data models.NodeFileItem, instData models.DataEntryIdentity, subjectId string) error {
	instName := instData.Name
	Db, err := InitMongoDB(config.WEBENV.PubDbName, NODES)
	if err != nil {
		return err
	}
	ctx := context.Background()
	data.Value = []interface{}{data.Value}
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Node Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:    data.Name,
		Id:      data.Uuid,
		Status:  STATUS.Activate(),
		RefId:   data.RefId,
		RefName: data.RefName,
		Thumb:   data.Thumb,
	}
	err = AddNewNodeToList(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Node List: %v", err))
	}
	return nil
}

func DeleteNodeItem(ref_id string) (models.SingleIDMap, error) {
	var R models.SingleIDMap
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

func UpdateNodeItem(data models.CreateNodeRequest, instData models.DataEntryIdentity, subjectId string, itemId string, ReqApi string, publishApiTarget string) (interface{}, error) {
	var R models.SingleIDMap
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
	if val := data.Description; val != "" {
		update["$set"].(bson.M)["description"] = val
	}
	if val := data.Schema; val != "" {
		update["$set"].(bson.M)["schema_ref"] = val
	}
	if val := data.Status; val != "" && STATUS.Contains(val) {
		update["$set"].(bson.M)["status"] = val
	}
	if val := data.Value; len(val) > 0 {
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
		Id:      itemId,
		Name:    recordDocument.Name,
		RefId:   recordDocument.RefId,
		RefName: recordDocument.RefName,
		Thumb: func() string {
			if val := data.Thumb; val != "" {
				return val
			}
			return recordDocument.Thumb
		}(),
		Status: func() string {
			if val := data.Status; val != "" && STATUS.Contains(val) {
				return val
			}
			return recordDocument.Status
		}(),
	}
	err = UpdateNodeListItem(instData.Name, subjectId, newRecord, false)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Node List: %v", err))
	}

	if data.Bump {
		publishResponse, err := PublishTarget(NODES, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord, publishApiTarget)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Nodes List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}

		pubDocument := publishResponse.Meta.(models.NodeFileItem)

		updeateRecord := models.DataEntryIdentity{
			Id:      pubDocument.Uuid,
			Name:    pubDocument.Name,
			Status:  pubDocument.Status,
			RefId:   pubDocument.RefId,
			RefName: pubDocument.RefName,
			Thumb:   pubDocument.Thumb,
		}
		err = UpdateNodeListItem(instData.RefId, subjectId, updeateRecord, data.Bump)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Nodes List for published item: %v", err))
		}
	}

	return map[string]any{
		"id":        itemId,
		"ref":       recordDocument.RefId,
		"operation": updateRes != nil,
		"published": data.Bump,
	}, nil
}
