package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateInstanceItem(data models.InstanceCollection, instanceData models.DataEntryIdentity, subjectId string, apiName string, publishTarget string, bump bool) (any, error) {
	refID := RandomRefID()
	data.RefId = refID
	err := CreatePublisedInstanceItem(data, instanceData, subjectId, publishTarget, refID, bump)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Published Instance [%s] Initial Data: %v", instanceData.Name, err))
		return "", err
	}

	DB, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		config.Err(fmt.Sprintf("Error initializing bd conn [%s] err: %v", instanceData.Name, err))
		return "", err
	}
	existingId, _ := VerifyInstanceExist(data.Name, apiName)

	if existingId.Id != "" {
		config.Err(fmt.Sprintf("Error Instance already Exist: %s", existingId))
		return "", fmt.Errorf("instance name in use: %s", data.Name)
	}
	RefId := data.Name
	Name := data.RefId
	data.Name = Name
	data.RefId = RefId

	err = DB.InsertOne(data)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceData.Name, err))
		return "", err
	}
	newID, err := AppendInstanceToRegistry(data, apiName, instanceData.Name)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceData.Name, err))
		return "", err
	}

	config.Log(fmt.Sprintf("Instance origin: [%s] Initial Data Saved: %v", instanceData.Name, newID))
	return refID, nil
}

func CreatePublisedInstanceItem(data models.InstanceCollection, instanceData models.DataEntryIdentity, subjectId string, apiName string, refID string, bump bool) error {
	DB, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		config.Err(fmt.Sprintf("Error initializing bd conn [%s] err: %v", instanceData.Name, err))
		return err
	}
	existingId, _ := VerifyInstanceExist(data.Name, apiName)

	if existingId.Id != "" {
		config.Err(fmt.Sprintf("Error Instance already Exist: %s", existingId))
		return fmt.Errorf("instance name in use: %s", data.Name)
	}
	data.ApiBase = apiName
	err = DB.InsertOne(data)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceData.Name, err))
		return err
	}
	newID, err := AppendInstanceToRegistry(data, apiName, refID)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceData.Name, err))
		return err
	}
	config.Log(fmt.Sprintf("Instance origin: [%s] Initial Data Saved: %v", instanceData.Name, newID))
	return nil
}

func AppendInstanceToRegistry(data models.InstanceCollection, apiName string, FKrefId string) (string, error) {
	instanceName := data.Name
	Db, err := InitMongoDB(config.WEBENV.IntDbName, apiName)
	if err != nil {
		return "", err
	}
	var refID string
	if FKrefId != "" {
		refID = FKrefId
	} else {
		refID = RandomRefID()
	}
	ctx := context.Background()
	item := models.DataEntryIdentity{
		Name:   instanceName,
		Id:     uuid.New().String(),
		Status: STATUS.Activate(),
		RefId:  refID,
	}
	_, err = Db.coll.InsertOne(ctx, item)
	if err != nil {
		return "", err
	}
	return item.Id, nil
}

func VerifyInstanceExist(instanceName string, apiName string) (models.DataEntryIdentity, error) {
	var res models.DataEntryIdentity
	Db, err := InitMongoDB(config.WEBENV.IntDbName, apiName)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GetInstanceInfo(instanceName string, subjectId string, ReqApi string) (models.InstanceCollection, error) {
	var res models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "api_base": ReqApi, "members": bson.M{"$in": []string{subjectId}}}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func PullInstanceInfo(instanceName string, ReqApi string) (models.InstanceCollection, error) {
	var res models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "api_base": ReqApi}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func AddNewNodeToList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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

func UpdateNodeListItem(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}, "nodes_collection_list.id": data.Id}
	update := bson.M{"$set": bson.M{
		"nodes_collection_list.$.name":   data.Name,
		"nodes_collection_list.$.status": data.Status,
		"nodes_collection_list.$.ref_id": data.RefId,
	}}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		config.Err(fmt.Sprintf("List of Node Items UPDATE_ERROR: %v", err))
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

func AddNewMediaList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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

func UpdateMediaListItem(instanceName string, subjectId string, data models.DataEntryIdentity, publish bool) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{"$set": bson.M{
		"media_files_collection_list.$.name":   data.Name,
		"media_files_collection_list.$.status": data.Status,
		"media_files_collection_list.$.ref_id": data.RefId,
	}}
	if publish {
		update["$set"].(bson.M)["media_files_collection_list.$.id"] = data.Id
	} else {
		identify["media_files_collection_list.id"] = data.Id
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		config.Err(fmt.Sprintf("List of MEdia Items UPDATE_ERROR: %v", err))
		return err
	}
	config.Log(fmt.Sprintf("List of MEdia Items UPDATED: %v", res))
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

func AddNewEndpointsList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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

func UpdateEndpointListItem(instanceName string, subjectId string, data models.DataEntryIdentity, publish bool) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}

	update := bson.M{"$set": bson.M{
		"endpoints_collection_list.$.name":   data.Name,
		"endpoints_collection_list.$.status": data.Status,
		"endpoints_collection_list.$.ref_id": data.RefId,
	}}
	if publish {
		update["$set"].(bson.M)["endpoints_collection_list.$.id"] = data.Id
	} else {
		identify["endpoints_collection_list.id"] = data.Id
	}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		config.Err(fmt.Sprintf("List of Endpoint Items UPDATE_ERROR: %v", err))
		return err
	}
	config.Log(fmt.Sprintf("List of Endpoint Items UPDATED: %v", res))
	return nil
}

func PublishEndpointListItem(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}, "endpoints_collection_list.id": data.Id}
	update := bson.M{"$set": bson.M{
		"endpoints_collection_list.$.name":   data.Name,
		"endpoints_collection_list.$.status": data.Status,
		"endpoints_collection_list.$.ref_id": data.RefId,
		"endpoints_collection_list.$.id":     data.Id,
	}}
	opts := options.Update().SetUpsert(true)
	res, err := Db.coll.UpdateOne(ctx, identify, update, opts)
	if err != nil {
		config.Err(fmt.Sprintf("List of Endpoint Items UPDATE_ERROR: %v", err))
		return err
	}
	config.Log(fmt.Sprintf("List of Endpoint Items UPDATED: %v", res))
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

func AddNewSchemasList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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

func UpdateSchemaListItem(instanceName string, subjectId string, data models.DataEntryIdentity) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}, "schemas_collection_list.id": data.Id}
	update := bson.M{"$set": bson.M{
		"schemas_collection_list.$.name":   data.Name,
		"schemas_collection_list.$.status": data.Status,
		"schemas_collection_list.$.ref_id": data.RefId,
	}}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		config.Err(fmt.Sprintf("List of Schema Items UPDATE_ERROR: %v", err))
		return err
	}
	config.Log(fmt.Sprintf("List of Schema Items UPDATED: %v", res))
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

func AddNewContentList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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

func UpdateContentListItem(instanceName string, subjectId string, data models.DataEntryIdentity, publish bool) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return err
	}
	ctx := context.Background()
	identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}}
	update := bson.M{"$set": bson.M{
		"files_collection_list.$.name":   data.Name,
		"files_collection_list.$.status": data.Status,
		"files_collection_list.$.ref_id": data.RefId,
	}}
	if publish {
		update["$set"].(bson.M)["files_collection_list.$.id"] = data.Id
	} else {
		identify["files_collection_list.id"] = data.Id
	}

	// identify := bson.M{"name": instanceName, "members": bson.M{"$in": []string{subjectId}}, "files_collection_list.id": data.Id}
	// update := bson.M{"$set": bson.M{
	// 	"files_collection_list.$.name":   data.Name,
	// 	"files_collection_list.$.status": data.Status,
	// 	"files_collection_list.$.ref_id": data.RefId,
	// }}
	res, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		config.Err(fmt.Sprintf("List of Content Items UPDATE_ERROR: %v", err))
		return err
	}
	config.Log(fmt.Sprintf("List of Content Items UPDATED: %v", res))
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

func DeleteInstanceItem(ref_id string) (models.SingleIDMap, error) {
	var R models.SingleIDMap
	var res models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
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

func AddNewEndpointToList(instanceName string, subjectId string, data models.DataEntryIdentity) error {
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
	config.Log(fmt.Sprintf("List of Endpoint Items UPDATED: %v", res))
	return nil
}

func UpdateInstanceItem(data models.CreateInstanceRequest, instData models.DataEntryIdentity, subjectId string, itemId string, apiName string) (interface{}, error) {
	var R models.SingleIDMap
	var recordDocument models.InstanceCollection
	Db, err := InitMongoDB(config.WEBENV.PubDbName, INSTANCE_INFO)
	if err != nil {
		return R, err
	}
	ctx := context.Background()
	identify := bson.M{"name": instData.Name}
	err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)
	if err != nil {
		return R, err
	}
	timeStamp := fmt.Sprintf("%v", time.Now().Unix())
	update := bson.M{
		"$set": bson.M{},
	}
	if val := data.Name; val != "" {
		return R, fmt.Errorf(": Name cannot be updated")
	}
	if val := data.Owner; val != "" {
		return R, fmt.Errorf(": Owner cannot be updated")
	}
	if val := data.Admin; len(val) > 0 {
		update["$set"].(bson.M)["admin"] = append(val, recordDocument.Admin...)
	}
	if val := data.Members; len(val) > 0 {
		// TODO: [` add and remove needs to be added `]-{2024-03-04}
		update["$set"].(bson.M)["members"] = append(val, recordDocument.Members...)
	}
	if val := data.Status; val != "" && STATUS.Contains(val) {
		sysData := models.SysData{
			CreationDate:     recordDocument.Sys.CreationDate,
			ModificationDate: timeStamp,
			ModifiedBy:       AppendModificationRecord(recordDocument.Sys.ModifiedBy, subjectId, timeStamp),
			CreatedBy:        recordDocument.Sys.CreatedBy,
			Status:           val,
		}
		update["$set"].(bson.M)["sys"] = sysData
		newRecord := models.DataEntryIdentity{
			Id:   itemId,
			Name: instData.Name,
			Status: func() string {
				if val := data.Status; val != "" && STATUS.Contains(val) {
					return val
				}
				return recordDocument.Sys.Status
			}(),
		}
		UpdateInstanceInRegistry(newRecord, instData.Name, apiName)
	}
	update["$set"].(bson.M)["versions"] = UpdateVersions(recordDocument.Versions, data.Bump)

	updateRes, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return R, err
	}

	return updateRes, nil
}

func UpdateInstanceInRegistry(newRecord models.DataEntryIdentity, instanceName string, apiName string) {
	Db, err := InitMongoDB(config.WEBENV.IntDbName, apiName)
	if err != nil {
		config.Err(fmt.Sprintf("Error initializing bd conn [%s] err: %v", instanceName, err))
	}
	ctx := context.Background()
	recordDocument := models.DataEntryIdentity{}
	identify := bson.M{"name": instanceName}
	err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)
	if err != nil {
		config.Err(fmt.Sprintf("Error verifying Instance Exist: %v", err))
	}
	item := models.DataEntryIdentity{
		Name:   instanceName,
		Id:     recordDocument.Id,
		RefId:  recordDocument.RefId,
		Status: newRecord.Status,
	}
	updateRes, err := Db.coll.UpdateOne(ctx, identify, item)
	if err != nil {
		config.Err(fmt.Sprintf("Error saving Instance [%s] Initial Data: %v", instanceName, err))
	}
	config.Log(fmt.Sprintf("Instance origin: [%v] Initial Data Saved: %v", updateRes, apiName))
}
