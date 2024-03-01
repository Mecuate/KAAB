package handlers

import (
	"encoding/json"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type NewStringArray struct {
	elements []string
}

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

func (s NewStringArray) hasAllOf(targets []string) bool {
	val := 1
	for _, elem := range targets {
		if s.Contains(elem) {
			val *= 1
		} else {
			val *= 0
		}
	}
	return val == 1
}

func MapToStringSlice(m map[string]string) []string {
	var result []string
	for key := range m {
		result = append(result, key)
	}
	return result
}

func RequestAuth(w http.ResponseWriter) {
	http.Header.Add(w.Header(), "WWW-Authenticate", `JWT realm="Restricted"`)
	http.Header.Add(w.Header(), "User-Token", `SESSION`)
	http.Error(w, "88", http.StatusUnauthorized)
}

func getReqApi(r *http.Request) (string, error) {
	availApis := NewStringArray{strings.Split(config.WEBENV.ApiVersions, ",")}
	curr := strings.Split(r.RequestURI, "/")[1]

	if availApis.Contains(curr) {
		return curr, nil
	}
	errmsg := fmt.Sprintf("Invalid API version rerquested: %s", curr)
	config.Err(errmsg)

	return "", fmt.Errorf(errmsg)
}

func GetRequestFSS(requestURI string) (models.URLFilterSearchParams, error) {
	var Res models.URLFilterSearchParams
	s := strings.Split(requestURI, "?")
	if len(s) > 1 {
		fss := strings.Split(s[1], "&")
		for _, v := range fss {
			c := strings.Split(v, "=")
			if len(c) > 1 {
				switch c[0] {
				case "v":
					Res.Version = c[1]
				case "p":
					Res.Pagination = c[1]
				case "l":
					Res.Limit = c[1]
				case "s":
					Res.Sorting = c[1]
				case "f":
					Res.Filters = c[1]
				}
			}
		}
	}

	return Res, nil
}

type KeyValue struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}
type obj map[string]interface{}

func MarshalKeyValueObject(value interface{}) ([]interface{}, error) {
	var r []map[string]interface{}
	var resp [][]KeyValue

	res, err := json.Marshal(value)
	if err != nil {
		return []interface{}{value}, err
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return []interface{}{value}, err
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

func AssortData(nodeItemValues []interface{}, ReqSearch models.URLFilterSearchParams, versions []string) []interface{} {
	var Res []interface{}
	var Result []interface{}
	var selItem int
	if ReqSearch.Version != "" {
		selItem = IndexOf(versions, ReqSearch.Version)
		if selItem > -1 {
			// compound, _ := MarshalKeyValueObject(nodeItemValues[selItem])
			Res = append(Res, nodeItemValues[selItem])
		}
	} else {
		selItem = 0
		// compound, _ := MarshalKeyValueObject(nodeItemValues[selItem])
		Res = append(Res, nodeItemValues[selItem])
	}

	var selectedItem []map[string]interface{}
	res, _ := json.Marshal(Res[0])
	json.Unmarshal(res, &selectedItem)

	if ReqSearch.Filters != "" {
		fmt.Println("ReqSearch.Filters: ", ReqSearch.Filters)
		// for _, v := range nodeItemValues {
		// 	if strings.Contains(fmt.Sprintf("%v", v), ReqSearch.Search) {
		// 		Res = append(Res, v)
		// 	}
		// }
	}
	if ReqSearch.Sorting != "" {
		fmt.Println("ReqSearch.Sorting: ", ReqSearch.Sorting)
		// for _, v := range nodeItemValues {
		// 	if strings.Contains(fmt.Sprintf("%v", v), ReqSearch.Search) {
		// 		Res = append(Res, v)
		// 	}
		// }
	}
	if ReqSearch.Limit != "" {
		fin, err := strconv.Atoi(ReqSearch.Limit)
		if err != nil {
			fin = len(selectedItem)
		}
		if ReqSearch.Pagination != "" {
			factor, err := strconv.Atoi(ReqSearch.Pagination)
			if err != nil {
				factor = 0
			}
			strt := factor * fin
			finn := strt + fin
			max := len(selectedItem)
			if finn > max {
				finn = max
			}
			if strt > max {
				strt = max
			}
			Result = append(Result, selectedItem[strt:finn])
		} else {
			max := len(selectedItem)
			if fin > max {
				fin = max
			}
			Result = append(Result, selectedItem[0:fin])
		}
	}

	return Result
}

func IndexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
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

func CreateMediaCtrlFields(ref_id string) models.InternalMediaCtrlFields {
	sys := ObtainSystemMedia(ref_id)
	res := models.InternalMediaCtrlFields{
		Thumb:      sys.ThumbAddres,
		Url:        sys.UrlAddress,
		UriAddress: sys.UriAddress,
		File:       sys.PhysicalAddress,
	}
	return res
}

func ObtainSystemMedia(ref_id string) models.SystemMediaAddress {
	res := models.SystemMediaAddress{
		UrlAddress:      fmt.Sprintf("https://kaab.mecuate.org/film/%s", ref_id),
		ThumbAddres:     fmt.Sprintf("https://kaab.mecuate.org/pub/%s/thumbs", ref_id),
		UriAddress:      fmt.Sprintf("home/kaab/mecuate/org/%s", ref_id),
		PhysicalAddress: fmt.Sprintf("ffmpeg-%s.m3u8", ref_id),
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
