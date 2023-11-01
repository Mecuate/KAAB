package utils

import (
	"fmt"
	"kaab/src/models"
	"strings"
)

func ComposeEndpointJS(instance models.EndpointInstance) (string, error) {
	context := strings.ReplaceAll(instance.Context, `"`, `\"`)
	return fmt.Sprintf(`{(context = JSON.parse("%s")), (response = (function (context) {
"use strict";
%s })(context));}`, context, instance.EndpointCode), nil
}
