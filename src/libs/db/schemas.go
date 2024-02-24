package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

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
	err = UpdateSchemasList(instName, subjectId, newRecord)
	if err != nil {
		config.Err(fmt.Sprintf("Error updating Schema List: %v", err))
	}
	return nil
}
