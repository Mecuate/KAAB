package db

import (
	"encoding/base64"
	"fmt"
	"kaab/src/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

func AppendModificationRecord(modifiedBy models.ModificationList, subjectId string, timeStamp string) models.ModificationList {
	var resp models.ModificationList
	xVal := models.ModificationRecord{
		Person: subjectId,
		Date:   timeStamp,
		Index:  0,
	}
	if len(modifiedBy) == 0 || modifiedBy == nil {
		resp = append(modifiedBy, xVal)
		return resp
	} else {
		if len(modifiedBy) >= 15 {
			resp = append(models.ModificationList{xVal}, modifiedBy[0:14]...)
		} else {
			resp = append(models.ModificationList{xVal}, modifiedBy...)
		}
	}
	for i := 1; i < len(resp); i++ {
		resp[i].Index = int16(i)
	}
	return resp
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

func CreateCtrlFields(idnt string) models.InternalCtrlFields {
	t := fmt.Sprintf("%v", time.Now().Unix())
	list := []string{"1.0"}

	res := models.InternalCtrlFields{
		Uuid:             uuid.New().String(),
		Size:             0,
		Versions:         list,
		CreationDate:     t,
		ModificationDate: t,
		ModifiedBy:       models.ModificationList{ModificationRecord(idnt, 0)},
		CreatedBy:        idnt,
	}
	return res
}

func ModificationRecord(idnt string, ix int16) models.ModificationRecord {
	t := fmt.Sprintf("%v", time.Now().Unix())
	return models.ModificationRecord{
		Person: idnt,
		Date:   t,
		Index:  ix,
	}
}
