package server

import "kaab/src/models"

// returned registred path param endpoints
var DynamicPath = models.DynamicPaths{
	EMULATED_API: "/web-engine/ep/{id}/{file}.json",
	USER:         "/user/{id}/{action}",
	DATA_ENTRY:   "/data/{id}/{file}",
}

type APIVersion struct {
	Version string
}

func (v *APIVersion) emulatedAPIPath() string {
	return "/" + v.Version + DynamicPath.EMULATED_API
}
func (v *APIVersion) userPath() string {
	return "/" + v.Version + DynamicPath.USER
}
func (v *APIVersion) dataEntryPath() string {
	return "/" + v.Version + DynamicPath.DATA_ENTRY
}
