package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"net/http"
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
