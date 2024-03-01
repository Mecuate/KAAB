package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

type obj map[string]interface{}
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

func CreateContentItem(data models.TextFileItem, instName string, subjectId string) error {
	Db, err := InitMongoDB(config.WEBENV.PubDbName, FILES)
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := Db.coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	config.Log(fmt.Sprintf("Content Item Created: %v", res))
	newRecord := models.DataEntryIdentity{
		Name:   data.Name,
		Id:     data.Uuid,
		Status: "active",
		RefId:  data.RefId,
	}
	err = UpdateContentList(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Content List: %v", err))
	}
	return nil
}

func DeleteContentItem(ref_id string) (models.Delition, error) {
	var R models.Delition
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
