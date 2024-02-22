package models

type NodeFileItemResponse struct {
	Uuid        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Size        int16         `json:"size" bson:"size"`
	Versions    []string      `json:"versions" bson:"versions" default:"[]"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
}
