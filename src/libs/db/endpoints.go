package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetEndpointItem(ref_id string, ReqApi string) (models.EndpointItem, error) {
	var res models.EndpointItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": ref_id, "api_base": ReqApi}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreateEndpointItem(data models.EndpointItem, instData models.DataEntryIdentity, subjectId string, ReqApi string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
	if err != nil {
		return err
	}
	ctx := context.Background()

	var endpointFileExist models.EndpointItem
	_ = Db.coll.FindOne(ctx, bson.M{"name": data.Name, "api_base": ReqApi}).Decode(&endpointFileExist)
	if endpointFileExist.Name != "" {
		return fmt.Errorf("endpoint already in use")
	}

	newReferenceID := RandomRefID()
	data.RefId = newReferenceID
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Endpoint Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: STATUS.Activate(),
		RefId:  newReferenceID,
		Thumb:  data.Thumb,
	}
	err = AddNewEndpointToList(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
	}
	return nil
}

func DeleteEndpointItem(ref_id string) (models.SingleIDMap, error) {
	var R models.SingleIDMap
	var res models.EndpointItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
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

func UpdateEndpointItem(data models.CreateEndpointRequest, instData models.DataEntryIdentity, subjectId string, itemId string, ReqApi string, publishApiTarget string) (interface{}, error) {
	var R models.SingleIDMap
	var recordDocument models.EndpointItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
	if err != nil {
		return R, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": itemId, "api_base": ReqApi}
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
	if val := data.Value; val.Get != "" || val.Post != "" || val.Delete != "" {
		values := make([]interface{}, len(recordDocument.Value))
		copy(values, recordDocument.Value)
		value := []interface{}{models.EndpointCode{Get: val.Get, Post: val.Post, Delete: val.Delete}}
		update["$set"].(bson.M)["value"] = AppendValue(values, value)
		update["$set"].(bson.M)["size"] = int64(len(fmt.Sprintf("%v", data.Value)))
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
	err = UpdateEndpointListItem(instData.Name, subjectId, newRecord, false)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
	}

	if data.Bump {
		publishResponse, err := PublishTarget(ENDPOINTS, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord, publishApiTarget)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}

		pubDocument := publishResponse.Meta.(models.EndpointItem)

		updeateRecord := models.DataEntryIdentity{
			Id:     pubDocument.Uuid,
			Name:   pubDocument.Name,
			Status: pubDocument.Status,
			RefId:  pubDocument.RefId,
			Thumb:  pubDocument.Thumb,
		}
		err = UpdateEndpointListItem(instData.RefId, subjectId, updeateRecord, data.Bump)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Endpoint List for published item: %v", err))
		}
	}

	return map[string]interface{}{
		"id":        itemId,
		"ref":       recordDocument.RefId,
		"operation": updateRes != nil,
		"published": data.Bump,
	}, nil
}
