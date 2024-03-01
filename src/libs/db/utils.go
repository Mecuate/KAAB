package db

import (
	"encoding/base64"
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
