package models

type DynamicPaths struct {
	EMULATED_API string
	USER         string
	DATA_ENTRY   string
}

type DynamicPathsParams struct {
	EMULATED_API []string
	USER         []string
	DATA_ENTRY   []string
}

type StaticPaths struct {
	LOGIN     string
	LOGOUT    string
	VERSION   string
	ANALYTICS string
}
