package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMediaItem(ref_id string, apiBase string) (models.MediaFileItem, error) {
	var res models.MediaFileItem
	var dims models.DimentionsType
	res.Dimensions = dims
	Db, err := InitMongoDB(config.WEBENV.PubDbName, MEDIA)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": ref_id, "api_base": apiBase}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreateMediaItem(data models.MediaFileItem, instData models.DataEntryIdentity, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, MEDIA)
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Media Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: data.Status,
		RefId:  data.RefId,
		Thumb:  data.Thumb,
	}
	err = AddNewMediaList(instData.Name, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Media List: %v", err))
	}
	return nil
}

func DeleteMediaItem(ref_id string) (models.Delition, error) {
	var R models.Delition
	var res models.MediaFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, MEDIA)
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

func UpdateMediaItem(data models.CreateMediaRequest, instData models.DataEntryIdentity, subjectId string, itemId string, ReqApi string, publishApiTarget string, mediaControlFiels models.InternalMediaCtrlFields) (interface{}, error) {
	var R models.Delition
	var recordDocument models.MediaFileItem
	Db, err := InitMongoDB(config.WEBENV.PubDbName, MEDIA)
	if err != nil {
		return R, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": itemId, "api_base": ReqApi}
	err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)
	if err != nil {
		return R, err
	}

	var mediaValueInfo models.MediaItemStorageValue
	update := ConformMediaUpdate(&mediaValueInfo, &data, &recordDocument, &mediaControlFiels, subjectId)
	fmt.Println("@mediaValueInfo", mediaValueInfo)
	fmt.Println("@mediaValueInfo", update)

	updateRes, err := Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return R, err
	}
	newRecord := models.DataEntryIdentity{
		Id:   itemId,
		Name: recordDocument.Name,
		Status: func() string {
			if val := data.Status; val != "" {
				return val
			}
			return recordDocument.Status
		}(),
		RefId: recordDocument.RefId,
	}
	err = UpdateMediaListItem(instData.Name, subjectId, newRecord, false)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Media List: %v", err))
	}

	if data.Bump {
		publishResponse, err := PublishTarget(MEDIA, instData, recordDocument.Name, recordDocument.RefId, subjectId, newRecord, publishApiTarget)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Media List: %v", err))
			return nil, fmt.Errorf("error publishing")
		}

		pubDocument := publishResponse.Meta.(models.MediaFileItem)

		updeateRecord := models.DataEntryIdentity{
			Id:     pubDocument.Uuid,
			Name:   pubDocument.Name,
			Status: pubDocument.Status,
			RefId:  pubDocument.RefId,
			Thumb:  pubDocument.Thumb,
		}
		err = UpdateMediaListItem(instData.RefId, subjectId, updeateRecord, data.Bump)
		if err != nil {
			config.Err(fmt.Sprintf("Error updating Endpoint List for published item: %v", err))
		}
	}

	return map[string]any{
		"id":        itemId,
		"ref":       recordDocument.RefId,
		"operation": updateRes != nil,
		"published": data.Bump,
	}, nil
}

func ConformMediaUpdate(mediaValueInfo *models.MediaItemStorageValue, data *models.CreateMediaRequest, recordDocument *models.MediaFileItem, mediaControlFiels *models.InternalMediaCtrlFields, subjectId string) bson.M {
	update := bson.M{
		"$set": bson.M{},
	}
	if val := data.Description; val != "" {
		update["$set"].(bson.M)["description"] = val
		mediaValueInfo.Description = val
	} else {
		mediaValueInfo.Description = recordDocument.Description
	}
	if val := data.Size; val >= 0 {
		update["$set"].(bson.M)["size"] = val
		mediaValueInfo.Size = val
	} else {
		mediaValueInfo.Size = recordDocument.Size
	}
	if val := data.Duration; val >= 0 {
		update["$set"].(bson.M)["duration"] = val
		mediaValueInfo.Duration = val
	} else {
		mediaValueInfo.Duration = recordDocument.Duration
	}
	if val := data.Dimensions; val == (models.DimentionsType{}) {
		update["$set"].(bson.M)["dimensions"] = val
		mediaValueInfo.Dimensions = val
	} else {
		mediaValueInfo.Dimensions = recordDocument.Dimensions
	}
	if val := data.Service; val != "" {
		update["$set"].(bson.M)["service"] = val
		mediaValueInfo.Service = val
	} else {
		mediaValueInfo.Service = recordDocument.Service
	}
	if val := data.Status; val != "" {
		update["$set"].(bson.M)["status"] = val
		mediaValueInfo.Status = val
	} else {
		mediaValueInfo.Status = recordDocument.Status
	}
	if val := data.Ttype; val != "" {
		update["$set"].(bson.M)["type"] = val
		mediaValueInfo.Ttype = val
	} else {
		mediaValueInfo.Ttype = recordDocument.Ttype
	}

	mediaValueInfo.Thumb = mediaControlFiels.Thumb
	mediaValueInfo.Url = mediaControlFiels.Url
	mediaValueInfo.UriAddress = mediaControlFiels.UriAddress
	mediaValueInfo.File = mediaControlFiels.File

	if chal := data.ChallengeID; chal != "" {
		update["$set"].(bson.M)["thumb"] = mediaControlFiels.Thumb
		update["$set"].(bson.M)["url"] = mediaControlFiels.Url
		update["$set"].(bson.M)["uri"] = mediaControlFiels.UriAddress
		update["$set"].(bson.M)["file_data"] = mediaControlFiels.File
	}

	valuesList := AppendMediaValue(recordDocument.Value, []interface{}{mediaValueInfo})
	update["$set"].(bson.M)["value"] = valuesList
	fmt.Println("@valuesList", valuesList)

	timeStamp := fmt.Sprintf("%v", time.Now().Unix())
	update["$set"].(bson.M)["versions"] = UpdateMediaVersions(recordDocument.Versions, data.Bump)
	update["$set"].(bson.M)["modified_by"] = AppendModificationRecord(recordDocument.ModifiedBy, subjectId, timeStamp)
	update["$set"].(bson.M)["modification_date"] = timeStamp

	return update
}
