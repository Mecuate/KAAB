package handlers

import (
	"fmt"
	"kaab/src/libs/config"
	"net/http"
	"strings"
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
