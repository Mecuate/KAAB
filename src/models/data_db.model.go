package models

type EndpointItem struct {
	Uuid        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	File        string   `json:"_file"`
	MemFile     string   `json:"mem_file"`
	Versions    []string `json:"_file_versions"`
}

type SchemaItem struct {
	Uuid        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	File        string   `json:"_file"`
	Versions    []string `json:"_file_versions"`
}

type TextFileItem struct {
	Uuid        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	File        string   `json:"_file"`
	Versions    []string `json:"_file_versions"`
	RefId       string   `json:"ref_id"`
	Size        int16    `json:"size"`
	Schema      string   `json:"_schema"`
}

type DimentionsType struct {
	Width  int16 `json:"width"`
	Height int16 `json:"height"`
}

type MediaFileItem struct {
	Uuid             string         `json:"uuid"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	RefId            string         `json:"ref_id"`
	Size             int16          `json:"size"`
	File             string         `json:"_file"`
	Versions         []string       `json:"_file_versions"`
	Ttype            string         `json:"type"`
	Duration         int16          `json:"duration"`
	CreationDate     string         `json:"creationDate"`
	ModificationDate string         `json:"modificationDate"`
	Dimentions       DimentionsType `json:"dimentions"`
	Service          string         `json:"_service"`
	Thumb            string         `json:"_thumb"`
	Url              string         `json:"_url"`
	UriAddress       string         `json:"_uri"`
}

type MediaFilesCollectionList []MediaFileItem

type EndpointsCollectionList []EndpointItem

type SchemasCollectionList []SchemaItem

type TextFilesCollectionList []TextFileItem

type InstanceCollection struct {
	Name           string                   `json:"collection_name"`
	EndpointsList  EndpointsCollectionList  `json:"endpoints_collection_list"`
	SchemasList    SchemasCollectionList    `json:"schemas_collection_list"`
	TextFilesList  TextFilesCollectionList  `json:"text_files_collection_list"`
	MediaFilesList MediaFilesCollectionList `json:"media_files_collection_list"`
}

type EndpointInstance struct {
	EndpointCode EndpointCode `json:"endpoint_code"`
	Context      string       `json:"context"`
}

type EndpointCode struct {
	Generic string `json:"_generic" default:""`
	Get     string `json:"_get" default:""`
	Post    string `json:"_post" default:""`
	Delete  string `json:"_delete" default:""`
}

type ModificationRecord struct {
	Person string `json:"_person"`
	Date   string `json:"_date"`
	Index  int16  `json:"_index"`
}

type ModificationRecords []ModificationRecord

type EndpointFile struct {
	Value            EndpointCode        `json:"_value"`
	ModifiedBy       ModificationRecords `json:"modified_by"`
	Versions         []string            `json:"_versions" default:"[]"`
	CreatedBy        string              `json:"created_by"`
	CreationDate     string              `json:"creationDate"`
	ModificationDate string              `json:"modificationDate"`
}

type DBstorageFile struct {
	Value            interface{}         `json:"_value"`
	ModifiedBy       ModificationRecords `json:"modified_by"`
	Versions         []string            `json:"_versions" default:"[]"`
	CreatedBy        string              `json:"created_by"`
	CreationDate     string              `json:"creationDate"`
	ModificationDate string              `json:"modificationDate"`
}
