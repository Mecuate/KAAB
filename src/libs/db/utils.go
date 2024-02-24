package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func EncodeSignature(instId string, usrId string) string {
	a := ([]byte)(instId)
	b := ([]byte)(usrId)
	enc := make([]byte, 3)

	for i := 0; i < len(b); i++ {
		enc = append(enc, a[i]+b[i])
	}
	data := []byte(string(enc))

	return base64.StdEncoding.EncodeToString(data)
}

func MarshalKeyValueObject(value []interface{}) ([]interface{}, error) {
	var r []map[string]interface{}
	var resp [][]KeyValue

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
			nested, ok := v.Value.([]interface{})
			if !ok {
				rx_[v.Key] = v.Value
			} else {
				nnested, err := MarshalKeyValueObjectItem(nested)
				if err != nil {
					val, err := MarshalKeyValueObject(v.Value.([]interface{}))
					if err != nil {
						rx_[v.Key] = v.Value
						continue
					}
					rx_[v.Key] = val
				} else {
					rx_[v.Key] = nnested
				}
			}
		}
		r = append(r, rx_)
	}
	var result []interface{}
	for _, kv := range r {
		result = append(result, kv)
	}
	return result, nil
}

func MarshalKeyValueObjectItem(value interface{}) (interface{}, error) {
	var resp []KeyValue
	res, err := json.Marshal(value)
	if err != nil {
		return value, err
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return value, err
	}
	if len(resp) > 0 {
		r := make(map[string]interface{})
		for _, v := range resp {
			r[v.Key] = v.Value
		}
		return r, nil
	}
	return value, fmt.Errorf("no nested object found")
}
