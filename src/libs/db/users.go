package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func PullUserData(userId string, instanceId string) (models.UserData, error) {
	var res models.UserData
	signature := EncodeSignature(instanceId, userId)

	fmt.Println("signature: ", signature)

	Db, err := InitMongoDB(config.WEBENV.PubDbName, USERS)
	if err != nil {
		return res, err
	}
	ctx := context.Background()
	identify := bson.M{"id": userId, "access_token": signature}
	err = Db.coll.FindOne(ctx, identify).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

/*
res:  73a18b3f-c0de-45fb-b8f4-3d5e5cb5b74f  ::  AAAAm2TGlmuSZZ5axpHKnlplZsvHWsRxy5ZaZpZqlW6ZxGibnWid
res:  c675b18b-3057-4ce5-b653-58b9e55bc673  ::  AAAAx2ecmpVhappalpGbcFpllMqaWsRvmpVaaGqXaZ5rl5WcnGtq
res:  cb538194-8fa1-4dad-8ff7-cbbceb6c42f8  ::  AAAAx5OamGtha2xam8fHalpllcbJWpqfy5lalpSXk56YmJZtmJpv
*/
