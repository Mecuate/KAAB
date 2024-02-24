package db

import (
	"context"
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
