package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CollectionType = map[string]interface{}{
	"nodes":     models.NodeFileItem{},
	"instance":  models.InstanceCollection{},
	"content":   models.TextFileItem{},
	"media":     models.MediaFileItem{},
	"schemas":   models.SchemaItem{},
	"endpoints": models.EndpointItem{},
}

var FailedPublishing = models.PublishResponse{
	Status: "failed",
}

var SuccessfulPublishing = models.PublishResponse{
	Status: "failed",
}

func PublishTarget(args ...any) (models.PublishResponse, error) {
	COLLECTION := args[0].(string)
	INST_DATA := args[1].(models.DataEntryIdentity)
	refid := args[2].(string)
	name := args[3].(string)
	publishApiTarget := args[6].(string)

	if _, ok := CollectionType[COLLECTION]; ok {
		target, err := InjectTarget(COLLECTION, name, refid, publishApiTarget)
		if err != nil {
			return FailedPublishing, err
		}

		return target, nil
	}
	return FailedPublishing, fmt.Errorf("publishing Target: %v:%v", COLLECTION, INST_DATA.Name)
}

func PullSource(COLLECTION string, name string, refid string) (any, error) {
	if recordType, ok := CollectionType[COLLECTION]; ok {
		var recordDocument = recordType
		Db, err := InitMongoDB(config.WEBENV.PubDbName, COLLECTION)
		if err != nil {
			return FailedPublishing, err
		}
		ctx := context.Background()
		identify := bson.M{"name": name, "ref_id": refid}
		err = Db.coll.FindOne(ctx, identify).Decode(&recordDocument)
		if err != nil {
			return FailedPublishing, err
		}
		return recordDocument, nil
	}
	return FailedPublishing, fmt.Errorf("pulling Source: %v:%v", COLLECTION, refid)
}

func InjectTarget(COLLECTION string, refid string, name string, publishApiTarget string) (models.PublishResponse, error) {
	if _, ok := CollectionType[COLLECTION]; ok {
		source, err := PullSource(COLLECTION, name, refid)
		if err != nil {
			return FailedPublishing, err
		}

		Db, err := InitMongoDB(config.WEBENV.PubDbName, COLLECTION)
		if err != nil {
			return FailedPublishing, err
		}
		ctx := context.Background()
		opts := options.Replace().SetUpsert(true)
		identify := bson.M{"name": name, "uuid": refid, "api_base": publishApiTarget}

		var recordDocument = assignToType(source, COLLECTION, refid, publishApiTarget)
		updateResult, err := Db.coll.ReplaceOne(ctx, identify, recordDocument, opts)
		if err != nil {
			return FailedPublishing, err
		}
		return models.PublishResponse{
			Meta:     recordDocument,
			TargetID: updateResult.UpsertedID,
		}, nil
	}
	return SuccessfulPublishing, nil
}

func assignToType(source interface{}, COLLECTION string, refid string, publishApiTarget string) interface{} {
	switch COLLECTION {
	case NODES:
		var customStruct models.NodeFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Value = customStruct.Value[0:1]
		customStruct.Versions = customStruct.Versions[0:1]
		customStruct.ModifiedBy = customStruct.ModifiedBy[0:1]
		customStruct.RefId = customStruct.Uuid
		customStruct.Uuid = refid
		customStruct.ApiBase = publishApiTarget

		return customStruct
	case SCHEMAS:
		var customStruct models.SchemaItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Value = customStruct.Value[0:1]
		customStruct.Versions = customStruct.Versions[0:1]
		customStruct.ModifiedBy = customStruct.ModifiedBy[0:1]
		customStruct.RefId = customStruct.Uuid
		customStruct.Uuid = refid
		customStruct.ApiBase = publishApiTarget

		return customStruct
	case ENDPOINTS:
		var customStruct models.EndpointItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Value = customStruct.Value[0:1]
		customStruct.Versions = customStruct.Versions[0:1]
		customStruct.ModifiedBy = customStruct.ModifiedBy[0:1]
		customStruct.RefId = customStruct.Uuid
		customStruct.Uuid = refid
		customStruct.ApiBase = publishApiTarget

		return customStruct
	case FILES:
		var customStruct models.TextFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Value = customStruct.Value[0:1]
		customStruct.Versions = customStruct.Versions[0:1]
		customStruct.ModifiedBy = customStruct.ModifiedBy[0:1]
		customStruct.RefId = customStruct.Uuid
		customStruct.Uuid = refid
		customStruct.ApiBase = publishApiTarget

		return customStruct
	case MEDIA:
		var customStruct models.MediaFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Value = customStruct.Value[0:1]
		customStruct.Versions = customStruct.Versions[0:1]
		customStruct.ModifiedBy = customStruct.ModifiedBy[0:1]
		customStruct.RefId = customStruct.Uuid
		customStruct.Uuid = refid
		customStruct.ApiBase = publishApiTarget

		return customStruct
	case INSTANCE_INFO:
		var customStruct models.InstanceCollection
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}

		return customStruct
	}
	return source
}
