package models

type KeyValue struct {
	Key   string      `bson:"Key"`
	Value interface{} `bson:"Value"`
}

type InternalCtrlFields struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Size             int64                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
}

type NodeItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int64         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Status      string        `json:"status" bson:"status"`
}

type ManyNodeItemResponse = []NodeItemResponse

type ContentItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int64         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Status      string        `json:"status" bson:"status"`
}

type ManyContentItemResponse = []ContentItemResponse

type SchemaItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int64         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	Status      string        `json:"status" bson:"status"`
}

type ManySchemaItemResponse = []SchemaItemResponse

type MediaItemResponse struct {
	Uuid        string         `json:"uuid" bson:"uuid"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Size        int64          `json:"size" bson:"size"`
	Versions    []string       `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{}  `json:"value" bson:"value"`
	Ttype       string         `json:"type" bson:"type"`
	Duration    int64          `json:"duration" bson:"duration"`
	Dimensions  DimentionsType `json:"dimensions" bson:"dimensions"`
	Service     string         `json:"service" bson:"service"`
	Thumb       string         `json:"thumb" bson:"thumb"`
	Url         string         `json:"url" bson:"url"`
	File        string         `json:"file_data" bson:"file_data"`
	Status      string         `json:"status" bson:"status"`
	RefId       string         `json:"ref_id" bson:"ref_id"`
}

type ManyMediaItemResponse = []MediaItemResponse

type EndpointItemResponse struct {
	Uuid           string      `json:"uuid" bson:"uuid"`
	Name           string      `json:"name" bson:"name"`
	Description    string      `json:"description" bson:"description"`
	Size           int64       `json:"size" bson:"size"`
	Versions       []string    `json:"versions" bson:"versions" default:"0.0"`
	CurrentVersion []string    `json:"current_version" bson:"current_version" default:"0.0"`
	Value          interface{} `json:"value" bson:"value"`
	RefId          string      `json:"ref_id" bson:"ref_id"`
	MemFile        string      `json:"mem_file" bson:"mem_file"`
	Status         string      `json:"status" bson:"status"`
}

type InternalMediaCtrlFields struct {
	Thumb      string `json:"thumb" bson:"thumb"`
	Url        string `json:"url" bson:"url"`
	UriAddress string `json:"uri" bson:"uri"`
	File       string `json:"file_data" bson:"file_data"`
}

type SystemMediaAddress struct {
	UrlAddress   string
	ThumbAddres  string
	UriAddress   string
	PhysicalName string
}

type Deletion struct {
	Id string `json:"id" bson:"id"`
}

type URLFilterSearchParams struct {
	Version    string `json:"version" bson:"version"`
	Sorting    string `json:"sorting" bson:"sorting"`
	Pagination string `json:"pagination" bson:"pagination"`
	Limit      string `json:"limit" bson:"limit"`
}

type PublishResponse struct {
	Meta     any `json:"meta"`
	TargetID any `json:"target_id"`
	Status   any `json:"status"`
}
