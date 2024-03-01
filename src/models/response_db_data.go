package models

type InternalCtrlFields struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Size             int16                `json:"size" bson:"size"`
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
	Size        int16         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
}

type ManyNodeItemResponse = []NodeItemResponse
type KeyValue struct {
	Key   string      `bson:"Key"`
	Value interface{} `bson:"Value"`
}
type ContentItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int16         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
}

type ManyContentItemResponse = []ContentItemResponse

type SchemaItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int16         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
}

type ManySchemaItemResponse = []SchemaItemResponse

type MediaItemResponse struct {
	Uuid        string         `json:"uuid" bson:"uuid"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Size        int16          `json:"size" bson:"size"`
	Versions    []string       `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{}  `json:"value" bson:"value"`
	RefId       string         `json:"ref_id" bson:"ref_id"`
	Ttype       string         `json:"type" bson:"type"`
	Duration    int16          `json:"duration" bson:"duration"`
	Dimensions  DimentionsType `json:"dimensions" bson:"dimensions"`
	Service     string         `json:"service" bson:"service"`
	Thumb       string         `json:"thumb" bson:"thumb"`
	Url         string         `json:"url" bson:"url"`
	UriAddress  string         `json:"uri" bson:"uri"`
	File        string         `json:"file_data" bson:"file_data"`
}

type ManyMediaItemResponse = []MediaItemResponse

type CreateNodeRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Bump        bool          `json:"bump" bson:"bump"`
}

type CreateContentRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Bump        bool          `json:"bump" bson:"bump"`
}

type CreateSchemaRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	Bump        bool          `json:"bump" bson:"bump"`
}

type CreateMediaRequest struct {
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Size        int16          `json:"size" bson:"size"`
	Value       []interface{}  `json:"value" bson:"value"`
	RefId       string         `json:"ref_id" bson:"ref_id"`
	Ttype       string         `json:"type" bson:"type"`
	Duration    int16          `json:"duration" bson:"duration"`
	Dimensions  DimentionsType `json:"dimensions" bson:"dimensions"`
	Service     string         `json:"service" bson:"service"`
	Bump        bool           `json:"bump" bson:"bump"`
}

type InternalMediaCtrlFields struct {
	Thumb      string `json:"thumb" bson:"thumb"`
	Url        string `json:"url" bson:"url"`
	UriAddress string `json:"uri" bson:"uri"`
	File       string `json:"file_data" bson:"file_data"`
}

type SystemMediaAddress struct {
	UrlAddress      string
	ThumbAddres     string
	UriAddress      string
	PhysicalAddress string
}

type Delition struct {
	Id string `json:"id" bson:"id"`
}

type URLFilterSearchParams struct {
	Version    string `json:"version" bson:"version"`
	Sorting    string `json:"sorting" bson:"sorting"`
	Pagination string `json:"pagination" bson:"pagination"`
	Limit      string `json:"limit" bson:"limit"`
}
