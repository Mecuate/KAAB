package models

type DimentionsType struct {
	Width  int16 `json:"width" bson:"width"`
	Height int16 `json:"height" bson:"height"`
}

type ModificationRecord struct {
	Person string `json:"_person" bson:"_person"`
	Date   string `json:"_date" bson:"_date"`
	Index  int16  `json:"_index" bson:"_index"`
}

type ModificationList []ModificationRecord

type DBstorageFile struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            interface{}      `json:"_value" bson:"_value"`
}

/* SPECIFIC TYPES */
type EndpointItem struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []interface{}    `json:"value" bson:"value"`
	MemFile          string           `json:"mem_file" bson:"mem_file"`
}

type SchemaItem struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []interface{}    `json:"value" bson:"value"`
}

type TextFileItem struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []interface{}    `json:"value" bson:"value"`
	RefId            string           `json:"ref_id" bson:"ref_id"`
	Schema           string           `json:"schema_ref" bson:"schema_ref"`
}

type NodeFileItem struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []interface{}    `json:"value" bson:"value"`
	RefId            string           `json:"ref_id" bson:"ref_id"`
	Schema           string           `json:"schema_ref" bson:"schema_ref"`
}

type MediaFileItem struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []interface{}    `json:"value" bson:"value"`
	RefId            string           `json:"ref_id" bson:"ref_id"`
	Ttype            string           `json:"type" bson:"type"`
	Duration         int16            `json:"duration" bson:"duration"`
	Dimensions       DimentionsType   `json:"dimensions" bson:"dimensions"`
	Service          string           `json:"service" bson:"service"`
	Thumb            string           `json:"thumb" bson:"thumb"`
	Url              string           `json:"url" bson:"url"`
	UriAddress       string           `json:"uri" bson:"uri"`
	File             string           `json:"file_data" bson:"file_data"`
}

type MediaFilesCollectionList []DataEntryIdentity

type EndpointsCollectionList []DataEntryIdentity

type SchemasCollectionList []DataEntryIdentity

type TextFilesCollectionList []DataEntryIdentity

type NodesFilesCollectionList []DataEntryIdentity

type InstanceCollection struct {
	Name           string                   `json:"name" bson:"name"`
	Versions       []string                 `json:"versions" bson:"versions" default:"[]"`
	Owner          string                   `json:"owner" bson:"owner"`
	Admin          []string                 `json:"admin" bson:"admin"`
	Members        []string                 `json:"members" bson:"members"`
	MediaFilesList MediaFilesCollectionList `json:"media_files_collection_list" bson:"media_files_collection_list"`
	EndpointsList  EndpointsCollectionList  `json:"endpoints_collection_list" bson:"endpoints_collection_list"`
	SchemasList    SchemasCollectionList    `json:"schemas_collection_list" bson:"schemas_collection_list"`
	TextFilesList  TextFilesCollectionList  `json:"files_collection_list" bson:"files_collection_list"`
	NodesFilesList NodesFilesCollectionList `json:"nodes_collection_list" bson:"nodes_collection_list"`
	Sys            SysData                  `json:"sys" bson:"sys"`
}

type SysData struct {
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	Status           string           `json:"status" bson:"status"`
}

type EndpointCode struct {
	Get    string `json:"get" bson:"get" default:""`
	Post   string `json:"post" bson:"post" default:""`
	Delete string `json:"delete" bson:"delete" default:""`
}

type EndpointInstance struct {
	EndpointCode EndpointCode `json:"endpoint_code"`
	Context      string       `json:"context"`
}

type EndpointFile struct {
	Uuid             string           `json:"uuid" bson:"uuid"`
	Name             string           `json:"name" bson:"name"`
	Description      string           `json:"description" bson:"description"`
	Size             int16            `json:"size" bson:"size"`
	Versions         []string         `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string           `json:"creation_date" bson:"creation_date"`
	ModificationDate string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy       ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy        string           `json:"created_by" bson:"created_by"`
	Value            []EndpointCode   `json:"value" bson:"value"`
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
type DataEntryIdentity struct {
	Name   string `json:"name" bson:"name"`
	Id     string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
	RefId  string `json:"ref_id" bson:"ref_id"`
}

type APICollections struct {
	Size      int16               `json:"size" bson:"size"`
	Instances []DataEntryIdentity `json:"instances" bson:"instances"`
}

type CollectionBasis struct {
	Uuid    string `json:"uuid" bson:"uuid"`
	Name    string `json:"_name" bson:"_name"`
	Created string `json:"created" bson:"created"`
}
