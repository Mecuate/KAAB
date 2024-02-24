package db

import (
	"context"
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
