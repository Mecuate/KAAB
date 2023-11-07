package handlers

import "kaab/src/models"

type APIVersion struct {
	Version string
}

var DynamicPath = models.DynamicPaths{
	EMULATED_API: "/web-engine/ep/{instance_id}/{file}.json",
	USER:         "/user/{id}/{action}",
	DATA_ENTRY:   "/data/{section}/{id}/{action}",
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

var Params = models.DynamicPathsParams{
	EMULATED_API: []string{"instance_id", "file"},
	USER:         []string{"id", "action"},
	DATA_ENTRY:   []string{"section", "id", "action"},
}
