package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

func UpdateVersions(versions []string, bump interface{}) []string {
	var resp []string
	var isBump = bump.(bool)

	if len(versions) == 0 || versions == nil || versions[0] == "" {
		return []string{"1.0"}
	} else {
		o := strings.Split(versions[0], ".")
		var xVal string
		if isBump {
			xN, _ := strconv.Atoi(o[0])
			xVal = fmt.Sprintf("%v.0", xN+1)
		} else {
			xN, _ := strconv.Atoi(o[1])
			xVal = fmt.Sprintf("%v.%v", o[0], xN+1)
		}

		if len(versions) >= 15 {
			resp = append([]string{xVal}, versions[0:14]...)
		} else {
			resp = append([]string{xVal}, versions...)
		}
	}
	return resp
}

func AppendValue(values []interface{}, newValue []interface{}) []interface{} {
	if len(values) == 0 || values == nil || values[0] == "" {
		return newValue
	} else {
		if len(values) >= 15 {
			values = append(newValue, values[0:14]...)
		} else {
			values = append(newValue, values...)
		}
	}
	return values
}
