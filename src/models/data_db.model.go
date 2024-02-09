package models

type EndpointItem struct {
	Uuid        string   `json:"uuid" bson:"uuid"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	File        string   `json:"_file" bson:"_file"`
	MemFile     string   `json:"mem_file" bson:"mem_file"`
	Versions    []string `json:"_file_versions" bson:"_file_versions"`
}

type SchemaItem struct {
	Uuid        string   `json:"uuid" bson:"uuid"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	File        string   `json:"_file" bson:"_file"`
	Versions    []string `json:"_file_versions" bson:"_file_versions"`
}

type TextFileItem struct {
	Uuid        string   `json:"uuid" bson:"uuid"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	File        string   `json:"_file" bson:"_file"`
	Versions    []string `json:"_file_versions" bson:"_file_versions"`
	RefId       string   `json:"ref_id" bson:"ref_id"`
	Size        int16    `json:"size" bson:"size"`
	Schema      string   `json:"_schema" bson:"_schema"`
}

type DimentionsType struct {
	Width  int16 `json:"width" bson:"width"`
	Height int16 `json:"height" bson:"height"`
}

type MediaFileItem struct {
	Uuid             string         `json:"uuid" bson:"uuid"`
	Name             string         `json:"name" bson:"name"`
	Description      string         `json:"description" bson:"description"`
	RefId            string         `json:"ref_id" bson:"ref_id"`
	Size             int16          `json:"size" bson:"size"`
	File             string         `json:"_file" bson:"_file" default:""`
	Versions         []string       `json:"_file_versions" bson:"_file_versions"`
	Ttype            string         `json:"type" bson:"type"`
	Duration         int16          `json:"duration" bson:"duration"`
	CreationDate     string         `json:"creationDate" bson:"creationDate"`
	ModificationDate string         `json:"modificationDate" bson:"modificationDate"`
	Dimentions       DimentionsType `json:"dimentions" bson:"dimentions"`
	Service          string         `json:"_service" bson:"_service"`
	Thumb            string         `json:"_thumb" bson:"_thumb"`
	Url              string         `json:"_url" bson:"_url"`
	UriAddress       string         `json:"_uri" bson:"_uri"`
}

type MediaFilesCollectionList []MediaFileItem

type EndpointsCollectionList []EndpointItem

type SchemasCollectionList []SchemaItem

type TextFilesCollectionList []TextFileItem

type InstanceCollection struct {
	Name           string                   `json:"collection_name" bson:"collection_name"`
	Uuid           string                   `json:"uuid" bson:"uuid"`
	Owner          string                   `json:"owner" bson:"owner"`
	Members        []string                 `json:"members" bson:"members"`
	Admin          []string                 `json:"admin" bson:"admin"`
	EndpointsList  EndpointsCollectionList  `json:"endpoints_collection_list" bson:"endpoints_collection_list"`
	SchemasList    SchemasCollectionList    `json:"schemas_collection_list" bson:"schemas_collection_list"`
	TextFilesList  TextFilesCollectionList  `json:"text_files_collection_list" bson:"text_files_collection_list"`
	MediaFilesList MediaFilesCollectionList `json:"media_files_collection_list" bson:"media_files_collection_list"`
}

type EndpointInstance struct {
	EndpointCode EndpointCode `json:"endpoint_code"`
	Context      string       `json:"context"`
}

type EndpointCode struct {
	Generic string `json:"_generic" bson:"_generic" default:""`
	Get     string `json:"_get" bson:"_get" default:""`
	Post    string `json:"_post" bson:"_post" default:""`
	Delete  string `json:"_delete" bson:"_delete" default:""`
}

type ModificationRecord struct {
	Person string `json:"_person" bson:"_person"`
	Date   string `json:"_date" bson:"_date"`
	Index  int16  `json:"_index" bson:"_index"`
}

type ModificationRecords []ModificationRecord

type EndpointFile struct {
	Uuid             string              `json:"uuid" bson:"uuid"`
	Value            EndpointCode        `json:"_value" bson:"_value"`
	ModifiedBy       ModificationRecords `json:"modified_by" bson:"modified_by"`
	Versions         []string            `json:"_versions" bson:"_versions" default:"[]"`
	CreatedBy        string              `json:"created_by" bson:"created_by"`
	CreationDate     string              `json:"creationDate" bson:"creationDate"`
	ModificationDate string              `json:"modificationDate" bson:"modificationDate"`
}

type DBstorageFile struct {
	Value            interface{}         `json:"_value" bson:"_value"`
	ModifiedBy       ModificationRecords `json:"modified_by" bson:"modified_by"`
	Versions         []string            `json:"_versions" bson:"_versions" default:"[]"`
	CreatedBy        string              `json:"created_by" bson:"created_by"`
	CreationDate     string              `json:"creationDate" bson:"creationDate"`
	ModificationDate string              `json:"modificationDate" bson:"modificationDate"`
}

// response for conformation files

type AccountConform struct {
	Account     AccountType `json:"account" bson:"account"`
	Email       string      `json:"email" bson:"email"`
	Id          string      `json:"id" bson:"id"`
	AccessToken string      `json:"access_token" bson:"access_token"`
	Uuid        string      `json:"uuid" bson:"uuid"`
}

type PermissionsConform struct {
	Uuid        string `json:"uuid" bson:"uuid"`
	Permissions RealmT `json:"realm" bson:"realm"`
	UserRol     string `json:"user_rol" bson:"user_rol"`
	Token       string `json:"token" bson:"token"`
}

type SecurityConform struct {
	Uuid      string   `json:"uuid" bson:"uuid"`
	Password  string   `json:"password" bson:"password"`
	Monitored bool     `json:"monitored" bson:"monitored"`
	KnownHost []string `json:"known_host" bson:"known_host"`
}

type ProfileConform struct {
	Uuid                 string `json:"uuid" bson:"uuid"`
	Name                 string `json:"name" bson:"name"`
	LastName             string `json:"last_name" bson:"last_name"`
	Nick                 string `json:"nick" bson:"nick"`
	UserRol              string `json:"user_rol" bson:"user_rol"`
	LastLogin            int64  `json:"last_login" bson:"last_login"`
	Modification_date    int64  `json:"modification_date" bson:"modification_date"`
	Picture              string `json:"picture" bson:"picture" default:""`
	PictureUrl           string `json:"picture_url" bson:"picture_url" default:""`
	PicModification_date int64  `json:"pic_modification_date" bson:"pic_modification_date" default:""`
	ExpirationDate       int64  `json:"expiration_date" bson:"expiration_date"`
}

type ReportConform struct {
	ReportFrame string `json:"report_frame" bson:"report_frame"`
}
type InstanceIdentData struct {
	Name string `json:"name" bson:"instances"`
	Id   string `json:"id" bson:"id"`
}

type APICollections struct {
	Size      int16               `json:"size" bson:"size"`
	Instances []InstanceIdentData `json:"instances" bson:"instances"`
}

type CollectionBasis struct {
	Uuid    string `json:"uuid" bson:"uuid"`
	Name    string `json:"_name" bson:"_name"`
	Created string `json:"created" bson:"created"`
}
