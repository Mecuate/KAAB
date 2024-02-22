package db

import (
	"context"
	"encoding/json"
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
	marshalled, _ := MarshalContentObject(res.Value)
	if err != nil {
		return res, err
	}
	res.Value = marshalled
	return res, nil
}

func MarshalContentObject(value []interface{}) ([]interface{}, error) {
	var resp [][]KeyValue
	var r []map[string]interface{}
	res, err := json.Marshal(value)
	if err != nil {
		return value, err
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return value, err
	}
	for i := 0; i < len(resp); i++ {
		rx_ := obj{}
		for _, v := range resp[i] {
			rx_[v.Key] = v.Value
		}
		r = append(r, rx_)
	}
	var result []interface{}
	for _, kv := range r {
		result = append(result, kv)
	}
	return result, nil
}
