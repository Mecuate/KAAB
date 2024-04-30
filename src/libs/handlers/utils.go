package handlers

import (
	"encoding/json"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type NewStringArray struct {
	elements []string
}

type KeyValue struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
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
	http.Error(w, "", http.StatusUnauthorized)
}

func getReqApi(r *http.Request) (string, error) {
	internal := strings.Split(config.WEBENV.ApiVersions, ",")
	pub := strings.Split(config.WEBENV.ApiPublishedVersions, ",")
	availApis := NewStringArray{append(internal, pub...)}
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
				}
			}
		}
	}

	return Res, nil
}

type KV = map[string]interface{}
type By func(p1, p2 *KV) bool
type pSorter struct {
	items []KV
	by    func(p1, p2 *KV) bool
}

func (by By) Sort(items []KV) {
	ps := &pSorter{
		items: items,
		by:    by,
	}
	sort.Sort(ps)
}

func (s *pSorter) Len() int {
	return len(s.items)
}

func (s *pSorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *pSorter) Less(i, j int) bool {
	return s.by(&s.items[i], &s.items[j])
}

func SortData(itemValues []KV, sortVal string) []map[string]interface{} {
	str := strings.Split(sortVal, ":")
	if len(str) < 3 {
		return itemValues
	}
	sortKey := str[0]
	sortDir := str[1] == "desc"
	dataT := str[2]

	strvals := func(p1, p2 *KV) bool {
		if a, ok := (*p1)[sortKey].(string); ok {
			if b, ook := (*p2)[sortKey].(string); ook {
				if sortDir {
					return a > b
				}
				return a < b
			}
		}
		return false
	}
	intvals := func(p1, p2 *KV) bool {
		if a, ok := (*p1)[sortKey].(float64); ok {
			if b, ook := (*p2)[sortKey].(float64); ook {
				if sortDir {
					return a > b
				}
				return a < b
			}
		}
		return false
	}
	bolvals := func(p1, p2 *KV) bool {
		if ar, ok := (*p1)[sortKey].(bool); ok {
			if br, ook := (*p2)[sortKey].(bool); ook {
				a := fmt.Sprintf("%v", ar)
				b := fmt.Sprintf("%v", br)
				if sortDir {
					return a > b
				}
				return a < b
			}
		}
		return false
	}

	switch dataT {
	case "str":
		By(strvals).Sort(itemValues)
	case "num":
		By(intvals).Sort(itemValues)
	case "bol":
		By(bolvals).Sort(itemValues)
	}

	return itemValues
}

type AssortedData struct {
	DataSelected    []interface{}
	VersionSelected string
}

func (A *AssortedData) data() []interface{} {
	return A.DataSelected
}
func (A *AssortedData) version() []string {
	return []string{A.VersionSelected}
}
func AssortData(itemValues []interface{}, ReqSearch models.URLFilterSearchParams, versions []string, flag bool) AssortedData {
	var Res []interface{}
	var Result []interface{}
	var selItem int
	if len(itemValues) == 0 {
		return AssortedData{
			DataSelected:    itemValues,
			VersionSelected: versions[0],
		}
	}
	if ReqSearch.Version != "" {
		selItem = IndexOf(versions, ReqSearch.Version)
		if selItem > -1 {
			Res = append(Res, itemValues[selItem])
		}
	} else {
		selItem = 0
		Res = append(Res, itemValues[selItem])
	}
	var selectedItem []map[string]interface{}
	res, _ := json.Marshal(Res[0])
	json.Unmarshal(res, &selectedItem)

	if ReqSearch.Sorting != "" && flag {
		selectedItem = SortData(selectedItem, ReqSearch.Sorting)
	}
	if ReqSearch.Limit != "" && flag {
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
	if Result == nil {
		Result = append(Result, selectedItem)
	}

	return AssortedData{
		DataSelected:    Result,
		VersionSelected: versions[selItem],
	}
}

func AssortEndpointData(itemValues []interface{}, ReqSearch models.URLFilterSearchParams, versions []string) AssortedData {
	var Res []interface{}
	var Result []interface{}
	var selItem int
	availVersions := NewStringArray{versions}
	if ReqSearch.Version != "" && availVersions.Contains(ReqSearch.Version) {
		selItem = IndexOf(versions, ReqSearch.Version)
		if selItem > -1 && selItem < len(itemValues) {
			Res = append(Res, itemValues[selItem])
		}
	} else {
		selItem = 0
		Res = append(Res, itemValues[selItem])
	}
	var selectedItem map[string]string
	res, _ := json.Marshal(Res[0])
	json.Unmarshal(res, &selectedItem)

	Result = append(Result, selectedItem)

	return AssortedData{
		DataSelected:    Result,
		VersionSelected: versions[selItem],
	}
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

func CreateMediaCtrlFields(challengeID string, instanceID string) models.InternalMediaCtrlFields {
	sys := ObtainSystemMedia(challengeID, instanceID)
	res := models.InternalMediaCtrlFields{
		Thumb:      sys.ThumbAddres,
		Url:        sys.UrlAddress,
		UriAddress: sys.UriAddress,
		File:       sys.PhysicalName,
	}
	return res
}

func ObtainSystemMedia(challengeID string, instanceID string) models.SystemMediaAddress {
	urlAddres := config.WEBENV.UrlAddress
	thumbs := config.WEBENV.Thumbs
	uriAddress := fmt.Sprintf(config.WEBENV.UriAddress, instanceID)
	physicalName := config.WEBENV.PhysicalName

	res := models.SystemMediaAddress{
		UrlAddress:   fmt.Sprintf(urlAddres, challengeID),
		ThumbAddres:  fmt.Sprintf(thumbs, challengeID),
		UriAddress:   fmt.Sprintf("%s/%s", uriAddress, challengeID),
		PhysicalName: fmt.Sprintf(physicalName, challengeID),
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

func MaskURIAddress(values any) []interface{} {
	Result := []interface{}{}
	address := []models.MediaItemStorageValue{}

	jps, err := json.Marshal(values)
	if err != nil {
		return Result
	}
	err = json.Unmarshal(jps, &address)
	if err != nil {
		return Result
	}

	for _, a := range address {
		a.UriAddress = "*"
		Result = append(Result, a)
	}
	return Result
}

func ExtractObjectItem(arr []interface{}) interface{} {
	if len(arr) == 0 {
		return models.MAPDATA{}
	}
	res := arr[0].([]map[string]interface{})
	return res[0]
}
