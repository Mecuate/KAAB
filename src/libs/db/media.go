package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMediaItem(ref_id string) (models.MediaFileItem, error) {
	var res models.MediaFileItem
	var dims models.DimentionsType
	res.Dimensions = dims
	Db, err := InitMongoDB(config.WEBENV.PubDbName, MEDIA)
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

func CreateMediaItem(data models.MediaFileItem, instName string, subjectId string) error {
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
		Status: "active",
		RefId:  data.Thumb,
	}
	err = UpdateMediaList(instName, subjectId, newRecord)
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
