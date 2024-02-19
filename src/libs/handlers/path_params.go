package handlers

import "kaab/src/models"

type APIVersion struct {
	Version string
}

var DynamicPath = models.DynamicPaths{
	EMULATED_API: "/open-service/web/{instance_id}/{file}.json",
	USER:         "/user/{instance_id}/{action}",
	DATA_ACTION:  "/data/{instance_id}/{section}/{action}/{ref_id}",
}

func (v *APIVersion) emulatedAPIPath() string {
	return "/" + v.Version + DynamicPath.EMULATED_API
}
func (v *APIVersion) userPath() string {
	return "/" + v.Version + DynamicPath.USER
}
func (v *APIVersion) dataEntryPath() string {
	return "/" + v.Version + DynamicPath.DATA_ACTION
}

var Params = models.DynamicPathsParams{
	EMULATED_API: []string{"instance_id", "file"},
	USER:         []string{"instance_id", "action"},
	DATA_ACTION:  []string{"instance_id", "section", "action", "ref_id"},
}
