package fixtures

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ComposeEndpointJS(data map[string]string, userCode string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	context := strings.ReplaceAll(string(jsonData), `"`, `\"`)
	return fmt.Sprintf(`{(context = JSON.parse("%s")), (response = (function (context) {
"use strict";
%s
})(context));}
console.log(response);`, context, userCode), nil
}
