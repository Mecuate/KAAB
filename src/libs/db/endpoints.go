package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetEndpointItem(ref_id string) (models.EndpointItem, error) {
	var res models.EndpointItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
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

func CreateEndpointItem(data models.EndpointItem, instData models.DataEntryIdentity, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
	if err != nil {
		return err
	}
	ctx := context.Background()

	var objectExist models.EndpointItem
	_ = Db.coll.FindOne(ctx, bson.M{"name": data.Name}).Decode(&objectExist)
	if objectExist.Name != "" {
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
		Status: data.Status,
		RefId:  newReferenceID,
	}
	err = AddNewEndpointToList(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
	}
	return nil
}

func DeleteEndpointItem(ref_id string) (models.Delition, error) {
	var R models.Delition
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

func UpdateEndpointItem(data models.CreateEndpointRequest, instData models.DataEntryIdentity, subjectId string, itemId string) (interface{}, error) {
	var R models.Delition
	var recordDocument models.EndpointItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, ENDPOINTS)
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
	if val := data.Status; val != "" {
		update["$set"].(bson.M)["status"] = val
	}
	if val := data.Value; len(val) > 0 {
		update["$set"].(bson.M)["value"] = AppendValue(recordDocument.Value, val)
		update["$set"].(bson.M)["size"] = int16(len(fmt.Sprintf("%v", data.Value)))

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
		RefId: recordDocument.RefId,
	}
	err = UpdateEndpointListItem(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
	}
	var respublish interface{}
	if data.Bump {
		respublish, err = PublishTarget(ENDPOINTS, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Endpoint List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}
	}

	return map[string]interface{}{
		"ref":       recordDocument.RefId,
		"id":        itemId,
		"operation": updateRes != nil,
		"published": respublish != nil,
	}, nil
}
