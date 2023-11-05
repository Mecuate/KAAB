package utils

import (
	"fmt"
	"kaab/src/models"
	"net/http"
	"strings"
)

func ComposeEndpointJS(instance models.EndpointInstance, req *http.Request) (string, error) {
	context := strings.ReplaceAll(instance.Context, `'`, `\"`)
	return fmt.Sprintf(`{(context = JSON.parse('%s')), 
	(response = (function (context,useContext,useRequest) {
"use strict";
%s 
})(context,function(str){
	if (context[str] != undefined) {
		return context[str];
	}
	return null;
},
function(){
	return { 
		method: "%s",
		query: %s,
		body: %s,
		headers: %s,
		url: %s,
		params: %s,
		};
}
));}`, context, instance.EndpointCode, req.Method, JSON(req.URL.Query()), JSON(req.Body), JSON(req.Header), JSON(req.URL), JSON(req.URL.Path)), nil
}
