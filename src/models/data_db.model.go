package models

type DimentionsType struct {
	Width  int16 `json:"width" bson:"width"`
	Height int16 `json:"height" bson:"height"`
}

type ModificationRecord struct {
	Person string `json:"person" bson:"person"`
	Date   string `json:"date" bson:"date"`
	Index  int16  `json:"index" bson:"index"`
}

//	type ItemBasicData struct {
//		Uuid             string               `json:"uuid" bson:"uuid"`
//		Name             string               `json:"name" bson:"name"`
//		Description      string               `json:"description" bson:"description"`
//		Size             int16                `json:"size" bson:"size"`
//		Versions         []string             `json:"versions" bson:"versions" default:"[]"`
//		CreationDate     string               `json:"creation_date" bson:"creation_date"`
//		ModificationDate string               `json:"modification_date" bson:"modification_date"`
//		ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
//		CreatedBy        string               `json:"created_by" bson:"created_by"`
//		Value            []interface{}        `json:"value" bson:"value"`
//	}
type DBstorageFile struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            interface{}          `json:"_value" bson:"_value"`
}

/* SPECIFIC TYPES */
type EndpointItem struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            []interface{}        `json:"value" bson:"value"`
	File             string               `json:"file_data" bson:"file_data"`
	MemFile          string               `json:"mem_file" bson:"mem_file"`
}

type SchemaItem struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            []interface{}        `json:"value" bson:"value"`
	File             string               `json:"file_data" bson:"file_data"`
}

type TextFileItem struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            []interface{}        `json:"value" bson:"value"`
	File             string               `json:"file_data" bson:"file_data"`
	RefId            string               `json:"ref_id" bson:"ref_id"`
	Schema           string               `json:"schema_ref" bson:"schema_ref"`
}

type NodeFileItem struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            []interface{}        `json:"value" bson:"value"`
	File             interface{}          `json:"file_data" bson:"file_data"`
	RefId            string               `json:"ref_id" bson:"ref_id"`
	Schema           string               `json:"schema_ref" bson:"schema_ref"`
}
type MediaFileItem struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            []interface{}        `json:"value" bson:"value"`
	RefId            string               `json:"ref_id" bson:"ref_id"`
	Ttype            string               `json:"type" bson:"type"`
	Duration         int16                `json:"duration" bson:"duration"`
	Dimentions       DimentionsType       `json:"dimentions" bson:"dimentions"`
	Service          string               `json:"service" bson:"service"`
	Thumb            string               `json:"thumb" bson:"thumb"`
	Url              string               `json:"url" bson:"url"`
	UriAddress       string               `json:"uri" bson:"uri"`
	File             string               `json:"file_data" bson:"file_data"`
}

type MediaFilesCollectionList []MediaFileItem

type EndpointsCollectionList []EndpointItem

type SchemasCollectionList []SchemaItem

type TextFilesCollectionList []TextFileItem

type NodesFilesCollectionList []NodeFileItem

type InstanceCollection struct {
	Uuid             string                   `json:"uuid" bson:"uuid"`
	Name             string                   `json:"name" bson:"name"`
	Description      string                   `json:"description" bson:"description"`
	Size             int16                    `json:"size" bson:"size"`
	Versions         []string                 `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string                   `json:"creation_date" bson:"creation_date"`
	ModificationDate string                   `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord     `json:"modified_by" bson:"modified_by"`
	CreatedBy        string                   `json:"created_by" bson:"created_by"`
	Value            []interface{}            `json:"value" bson:"value"`
	Owner            string                   `json:"owner" bson:"owner"`
	Members          []string                 `json:"members" bson:"members"`
	Admin            []string                 `json:"admin" bson:"admin"`
	EndpointsList    EndpointsCollectionList  `json:"endpoints_collection_list" bson:"endpoints_collection_list"`
	SchemasList      SchemasCollectionList    `json:"schemas_collection_list" bson:"schemas_collection_list"`
	TextFilesList    TextFilesCollectionList  `json:"files_collection_list" bson:"files_collection_list"`
	NodesFilesList   NodesFilesCollectionList `json:"nodes_collection_list" bson:"nodes_collection_list"`
	MediaFilesList   MediaFilesCollectionList `json:"media_files_collection_list" bson:"media_files_collection_list"`
	Sys              map[string]string        `json:"sys" bson:"sys"`
}

type EndpointCode struct {
	Generic string `json:"generic" bson:"generic" default:""`
	Get     string `json:"get" bson:"get" default:""`
	Post    string `json:"post" bson:"post" default:""`
	Delete  string `json:"delete" bson:"delete" default:""`
}

type EndpointInstance struct {
	EndpointCode EndpointCode `json:"endpoint_code"`
	Context      string       `json:"context"`
}

type EndpointFile struct {
	Uuid             string               `json:"uuid" bson:"uuid"`
	Name             string               `json:"name" bson:"name"`
	Description      string               `json:"description" bson:"description"`
	Size             int16                `json:"size" bson:"size"`
	Versions         []string             `json:"versions" bson:"versions" default:"[]"`
	CreationDate     string               `json:"creation_date" bson:"creation_date"`
	ModificationDate string               `json:"modification_date" bson:"modification_date"`
	ModifiedBy       []ModificationRecord `json:"modified_by" bson:"modified_by"`
	CreatedBy        string               `json:"created_by" bson:"created_by"`
	Value            EndpointCode         `json:"value" bson:"value"`
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
	Name string `json:"name" bson:"name"`
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
