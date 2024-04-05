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

var FailedPublishing = map[string]string{
	"status": "failed",
}

var SuccessfulPublishing = map[string]string{
	"status": "failed",
}

func PublishTarget(args ...any) (interface{}, error) {
	/* parameters collection , istance:name.refid.status.id , current.refId, current.id*/
	COLLECTION := args[0].(string)
	INST_DATA := args[1].(models.DataEntryIdentity)
	refid := args[2].(string)
	name := args[3].(string)
	// subjectId := args[4].(string)
	// newRecord := args[5].(models.DataEntryIdentity)

	if _, ok := CollectionType[COLLECTION]; ok {
		// fmt.Println("subject_id: ", subjectId, "newRecord", newRecord)
		target, err := InjectTarget(COLLECTION, name, refid)
		if err != nil {
			return FailedPublishing, err
		}

		return map[string]any{
			"meta":   SuccessfulPublishing,
			"target": target,
		}, nil
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

func InjectTarget(COLLECTION string, refid string, name string) (any, error) {
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
		identify := bson.M{"name": refid, "ref_id": name}

		var recordDocument = assignToType(source, COLLECTION, refid, name)
		updateResult, err := Db.coll.ReplaceOne(ctx, identify, recordDocument, opts)
		if err != nil {
			return FailedPublishing, err
		}
		return map[string]any{
			"meta":   SuccessfulPublishing,
			"target": updateResult.UpsertedID,
		}, nil
	}
	return SuccessfulPublishing, nil
}

func assignToType(source interface{}, COLLECTION string, refid string, name string) interface{} {
	switch COLLECTION {
	case "nodes":
		var customStruct models.NodeFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Name = refid
		customStruct.RefId = name
		return customStruct
	case "schemas":
		var customStruct models.SchemaItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Name = refid
		customStruct.RefId = name
		return customStruct
	case "endpoints":
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
		return customStruct
	case "content":
		var customStruct models.TextFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Name = refid
		customStruct.RefId = name
		return customStruct
	case "media":
		var customStruct models.MediaFileItem
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Name = refid
		customStruct.RefId = name
		return customStruct
	case "instance":
		var customStruct models.InstanceCollection
		data, err := bson.Marshal(source)
		if err != nil {
			return source
		}
		err = bson.Unmarshal(data, &customStruct)
		if err != nil {
			return source
		}
		customStruct.Name = refid
		customStruct.RefId = name
		return customStruct
	}
	return source
}

/* 660cd8518394427830a593b8 */
