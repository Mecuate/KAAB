package models

type DynamicPaths struct {
	EMULATED_API string
	USER         string
	DATA_ACTION  string
}

type DynamicPathsParams struct {
	EMULATED_API []string
	USER         []string
	DATA_ACTION  []string
}
