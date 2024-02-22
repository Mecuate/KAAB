package db

import (
	"context"
	"kaab/src/libs/config"
	"kaab/src/models"

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
