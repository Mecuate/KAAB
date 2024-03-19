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
		Token:             payload.Token,
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

func UpdateUser(userData models.UserData, payload models.UpdateProfileRequestBody, subject string) (interface{}, error) {
	timeStamp := fmt.Sprintf("%v", time.Now().Unix())
	var Res interface{}
	var recordDocument models.UserData

	Db, err := InitMongoDB(config.WEBENV.PubDbName, USERS)
	if err != nil {
		return Res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": subject}
	err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)

	if err != nil {
		return Res, err
	}

	updatedCheck := false
	update := bson.M{
		"$set": bson.M{},
	}
	khosts := NewStringArray{recordDocument.KnownHost}
	if khosts.Contains(payload.KnownHost[0]) {
		updatedCheck = true
		update["$set"].(bson.M)["known_host"] = append(khosts.elements, payload.KnownHost[0])
	}
	if val := payload.Email; val != recordDocument.Email {
		updatedCheck = true
		update["$set"].(bson.M)["email"] = val
	}
	if val := payload.Monitored; val != recordDocument.Monitored {
		updatedCheck = true
		update["$set"].(bson.M)["monitored"] = val
	}
	if val := payload.Name; val != "" {
		updatedCheck = true
		update["$set"].(bson.M)["name"] = val
	}
	if val := payload.LastName; val != "" {
		updatedCheck = true
		update["$set"].(bson.M)["last_name"] = val
	}
	if val := payload.Nick; val != "" {
		updatedCheck = true
		update["$set"].(bson.M)["nick"] = val
	}
	if val := payload.Password; val != "" {
		updatedCheck = true
		update["$set"].(bson.M)["nick"] = MakeSHA1Hash(val)
	}
	if val := payload.Token; val != "" && val != recordDocument.Token {
		updatedCheck = true
		update["$set"].(bson.M)["token"] = val
	}
	if val := payload.UserRol; val != "" && val != recordDocument.UserRol {
		updatedCheck = true
		update["$set"].(bson.M)["user_rol"] = val
	}
	if val := payload.UserRol; val != "" && val != recordDocument.UserRol {
		updatedCheck = true
		update["$set"].(bson.M)["user_rol"] = val
	}
	var realms models.RealmT
	realmsCheck := false
	if val := payload.Apis; val != "" && val != recordDocument.Realm.Apis {
		realmsCheck = true
		realms.Apis = val
	} else {
		realms.Apis = recordDocument.Realm.Apis
	}
	if val := payload.Media; val != "" && val != recordDocument.Realm.Media {
		realmsCheck = true
		realms.Media = val
	} else {
		realms.Media = recordDocument.Realm.Media
	}
	if val := payload.Mecuate; val != "" && val != recordDocument.Realm.Mecuate {
		realmsCheck = true
		realms.Mecuate = val
	} else {
		realms.Mecuate = recordDocument.Realm.Mecuate
	}
	if realmsCheck {
		updatedCheck = true
		update["$set"].(bson.M)["realm"] = realms
	}

	var account models.AccountType
	accountCheck := false
	if val := payload.Picture; val != "" && val != recordDocument.Account.Picture {
		accountCheck = true
		account.Picture = val
	} else {
		account.Picture = recordDocument.Account.Picture
	}
	if val := payload.PictureUrl; val != "" && val != recordDocument.Account.PictureUrl {
		accountCheck = true
		account.PictureUrl = val
	} else {
		account.PictureUrl = recordDocument.Account.PictureUrl
	}
	if val := payload.ExpirationDate; val != "" && val != recordDocument.Account.ExpirationDate {
		accountCheck = true
		account.ExpirationDate = val
	} else {
		account.ExpirationDate = recordDocument.Account.ExpirationDate
	}
	ApprovedBy := NewStringArray{recordDocument.Account.ApprovedBy}
	if len(payload.ApprovedBy) > 0 {
		accountCheck = true
		ApprovedBy.Join(payload.ApprovedBy)
		account.ApprovedBy = ApprovedBy.elements
	} else {
		account.ApprovedBy = ApprovedBy.elements
	}
	ApproverOf := NewStringArray{recordDocument.Account.ApproverOf}
	if len(payload.ApproverOf) > 0 {
		accountCheck = true
		ApproverOf.Join(payload.ApproverOf)
		account.ApproverOf = ApproverOf.elements
	} else {
		account.ApproverOf = ApproverOf.elements
	}

	if accountCheck {
		account.CreationDate = recordDocument.Account.CreationDate
		account.Modification_date = timeStamp
		account.ModifiedBy = AppendModificationRecord(recordDocument.Account.ModifiedBy, userData.Uuid, timeStamp)
		account.CreatedBy = recordDocument.Account.CreatedBy
		update["$set"].(bson.M)["account"] = account
	}

	if updatedCheck {
		update["$set"].(bson.M)["modified_by"] = AppendModificationRecord(recordDocument.ModifiedBy, userData.Uuid, timeStamp)
		update["$set"].(bson.M)["modification_date"] = timeStamp
	}

	Res, err = Db.coll.UpdateOne(ctx, identify, update)
	if err != nil {
		return Res, err
	}

	return Res, nil
}

func DeleteUser(userData models.UserData, instanceId string) (interface{}, error) {
	userId := userData.Uuid
	var Res interface{}
	Db, err := InitMongoDB(config.WEBENV.PubDbName, USERS)
	if err != nil {
		return Res, err
	}
	ctx := context.Background()
	identify := bson.M{"uuid": userId}
	Res, err = Db.coll.DeleteOne(ctx, identify)
	if err != nil {
		return Res, err
	}

	return Res, nil
}
