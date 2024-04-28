package db

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"kaab/src/models"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		resp[i].Index = int64(i)
	}
	return resp
}

func UpdateVersions(versions []string, bump interface{}) []string {
	var resp []string
	var isBump = bump.(bool)

	if len(versions) == 0 || versions == nil || versions[0] == "" {
		return []string{"0.0"}
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

func UpdateMediaVersions(versions []string, bump interface{}) []string {
	var resp []string
	var isBump = bump.(bool)

	if len(versions) == 0 || versions == nil || versions[0] == "" {
		return []string{"0.0"}
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

		if len(versions) >= 3 {
			resp = append([]string{xVal}, versions[0:2]...)
		} else {
			resp = append([]string{xVal}, versions...)
		}
	}
	return resp
}

func AppendMediaValue(values []interface{}, newValue []interface{}) []interface{} {
	if len(values) == 0 || values == nil || values[0] == "" {
		return newValue
	} else {
		if len(values) >= 3 {
			values = append(newValue, values[0:2]...)
		} else {
			values = append(newValue, values...)
		}
	}
	return values
}

func CreateCtrlFields(idnt string) models.InternalCtrlFields {
	t := fmt.Sprintf("%v", time.Now().Unix())
	list := []string{"0.0"}

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

func ModificationRecord(idnt string, ix int64) models.ModificationRecord {
	t := fmt.Sprintf("%v", time.Now().Unix())
	return models.ModificationRecord{
		Person: idnt,
		Date:   t,
		Index:  ix,
	}
}

type NewStringArray struct {
	elements []string
}

/* funcs */
func (s NewStringArray) Contains(target string) bool {
	for _, elem := range s.elements {
		if elem == target {
			return true
		}
	}
	return false
}

func (s NewStringArray) ContainsKey(target string) (string, bool) {
	for _, elem := range s.elements {
		if elem == target {
			return elem, true
		}
	}
	return "", false
}

func (s NewStringArray) Join(targets []string) {
	for _, elem := range targets {
		if !s.Contains(elem) {
			s.elements = append(s.elements, elem)
		}
	}
}

func MakeSHA1Hash(data string) string {
	bv := []byte(data)
	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

func MakeHash(data string) string {
	bv := []byte(data)
	hasher := sha1.New()
	hasher.Write(bv)

	return hex.EncodeToString(hasher.Sum(nil))
}

func RandomRefID() string {
	return uuid.New().String()
}

func ShortRefID() string {
	numBytes := (12 * 4) / 2
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return cleanString("f---------------")
	}
	res := base64.URLEncoding.EncodeToString(randomBytes)
	return cleanString(res[:16])
}

func cleanString(input string) string {
	mrand.Seed(time.Now().UnixNano())
	result := []rune(input)
	for i, char := range result {
		if char == '-' || char == '_' {
			randomChar := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdeghijklmnopqrstuvwxyz0123456789")[mrand.Intn(61)]
			result[i] = randomChar
		}
	}
	return string(result)
}

func IndexOf(slice []string, item string) int64 {
	if item == "" {
		return int64(0)
	}
	for i, v := range slice {
		if v == item {
			return int64(i)
		}
	}
	return int64(-1)
}

func convertToMapArray(arr primitive.A) []models.MAPDATA {
	result := make([]models.MAPDATA, len(arr))
	for i, v := range arr {
		switch elem := v.(type) {
		case primitive.M:
			result[i] = models.MAPDATA(elem)
			continue
		default:
			result[i] = models.MAPDATA(nil)
		}
	}
	return result
}

func CreateToMapArray(arr []interface{}) []models.MAPDATA {
	result := make([]models.MAPDATA, len(arr))
	for i, v := range arr {
		switch elem := v.(type) {
		case interface{}:
			result[i] = elem.(map[string]interface{})
			continue
		default:
			result[i] = elem.(map[string]interface{})
		}
	}
	return result
}

func DeleteItemFromArray(arr *[]models.MAPDATA, index int) {
	if index < 0 || index >= len(*arr) {
		return
	}
	start := (*arr)[:index]
	end := (*arr)[index+1:]
	*arr = append(start, end...)
}

func UpdateItemFromArray(arr *[]models.MAPDATA, index int, mod models.MAPDATA) {
	if index < 0 || index >= len(*arr) {
		return
	}
	start := append((*arr)[:index], mod)
	end := (*arr)[index+1:]
	*arr = append(start, end...)
}

func AddItemFromArray(arr *[]models.MAPDATA, index int, newItem models.MAPDATA) {
	if index < len(*arr) {
		return
	}
	newItem["$__i"] = MakeHash(fmt.Sprintf("%d:%v:v_:%v", index, time.Now().UnixNano(), newItem))
	*arr = append(*arr, newItem)
}

func HandleContentCreationItems(arr []models.MAPDATA) []interface{} {
	Total := len(arr)
	if Total <= 0 {
		return []interface{}{[]models.MAPDATA{}}
	}
	res := make([]models.MAPDATA, Total)
	for i, v := range arr {
		hash := MakeHash(fmt.Sprintf("%d:%v:v_:%v", i, time.Now().UnixNano(), v))[30:]
		newItem := v
		newItem["$__i"] = hash
		res[i] = newItem
	}
	return []interface{}{res}
}
