package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func PullUserData(userId string, instanceId string) (models.UserData, error) {
	var res models.UserData
	signature := EncodeSignature(instanceId, userId)
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

func CreateUser(userData models.UserData, payload models.CreateUserRequestBody) (interface{}, error) {
	var F interface{}
	Db, err := InitMongoDB(config.WEBENV.PubDbName, USERS)
	if err != nil {
		return F, err
	}
	ctx := context.Background()
	userUUID := uuid.New().String()
	currentTime := fmt.Sprintf("%v", time.Now().Unix())
	ctrlFields := CreateCtrlFields(userData.Uuid)
	userAccount := models.AccountType{
		ApprovedBy:        payload.ApprovedBy,
		ApproverOf:        payload.ApproverOf,
		ExpirationDate:    payload.ExpirationDate,
		Picture:           payload.Picture,
		PictureUrl:        payload.PictureUrl,
		CreationDate:      ctrlFields.CreationDate,
		Modification_date: ctrlFields.ModificationDate,
		ModifiedBy:        ctrlFields.ModifiedBy,
		CreatedBy:         ctrlFields.CreatedBy,
	}
	RealmData := models.RealmT{
		Apis:    payload.Apis,
		Media:   payload.Media,
		Mecuate: payload.Mecuate,
	}
	newUser := models.UserData{
		AccessToken:       userData.AccessToken,
		Account:           userAccount,
		Email:             payload.Email,
		Uuid:              userUUID,
		KnownHost:         payload.KnownHost,
		LastLogin:         currentTime,
		Monitored:         payload.Monitored,
		Name:              payload.Name,
		LastName:          payload.LastName,
		Nick:              payload.Nick,
		Password:          payload.Password,
		Realm:             RealmData,
		Token:             userData.Token,
		UserRol:           payload.UserRol,
		CreationDate:      ctrlFields.CreationDate,
		Modification_date: ctrlFields.ModificationDate,
		ModifiedBy:        ctrlFields.ModifiedBy,
		CreatedBy:         ctrlFields.CreatedBy,
	}

	F, err = Db.coll.InsertOne(ctx, newUser)
	if err != nil {
		return F, err
	}

	return F, nil
}

func UpdateUserProfile(userData models.UserData, payload models.UpdateProfileRequestBody) (interface{}, error) {
	var F interface{}
	Db, err := InitMongoDB(config.WEBENV.PubDbName, USERS)
	if err != nil {
		return F, err
	}
	ctx := context.Background()
	userUUID := uuid.New().String()
	currentTime := fmt.Sprintf("%v", time.Now().Unix())
	ctrlFields := CreateCtrlFields(userData.Uuid)
	userAccount := models.AccountType{
		ApprovedBy:        payload.ApprovedBy,
		ApproverOf:        payload.ApproverOf,
		ExpirationDate:    payload.ExpirationDate,
		Picture:           payload.Picture,
		PictureUrl:        payload.PictureUrl,
		CreationDate:      ctrlFields.CreationDate,
		Modification_date: ctrlFields.ModificationDate,
		ModifiedBy:        ctrlFields.ModifiedBy,
		CreatedBy:         ctrlFields.CreatedBy,
	}
	RealmData := models.RealmT{
		Apis:    payload.Apis,
		Media:   payload.Media,
		Mecuate: payload.Mecuate,
	}
	newUser := models.UserData{
		AccessToken:       userData.AccessToken,
		Account:           userAccount,
		Email:             payload.Email,
		Uuid:              userUUID,
		KnownHost:         payload.KnownHost,
		LastLogin:         currentTime,
		Monitored:         payload.Monitored,
		Name:              payload.Name,
		LastName:          payload.LastName,
		Nick:              payload.Nick,
		Password:          payload.Password,
		Realm:             RealmData,
		Token:             userData.Token,
		UserRol:           payload.UserRol,
		CreationDate:      ctrlFields.CreationDate,
		Modification_date: ctrlFields.ModificationDate,
		ModifiedBy:        ctrlFields.ModifiedBy,
		CreatedBy:         ctrlFields.CreatedBy,
	}

	F, err = Db.coll.InsertOne(ctx, newUser)
	if err != nil {
		return F, err
	}

	return F, nil
}
