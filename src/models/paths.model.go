package models

type InstancePaths struct {
	EMULATED_API string
	USER         string
	DATA_ACTION  string
}

type InstancePathsParams struct {
	EMULATED_API []string
	USER         []string
	DATA_ACTION  []string
}
