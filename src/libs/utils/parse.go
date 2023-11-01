package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBODY(r *http.Request, mo interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return ParseJSON(body, &mo)
}

func ParseVerify(r *http.Request, mo interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return ParseJSON(body, &mo)
}

func ParseJSON(data []byte, model interface{}) error {
	err := json.Unmarshal(data, &model)
	if err != nil {
		return err
	}
	return nil
}
