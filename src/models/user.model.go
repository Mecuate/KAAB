package models

type RealmT struct {
	Apis    string `json:"apis" bson:"apis"`
	Media   string `json:"media" bson:"media"`
	Mecuate string `json:"mecuate" bson:"mecuate"`
}

type UserData struct {
	Account           AccountType      `json:"account" bson:"account"`
	Email             string           `json:"email" bson:"email"`
	Uuid              string           `json:"uuid" bson:"uuid"`
	KnownHost         []string         `json:"known_host" bson:"known_host"`
	LastLogin         string           `json:"last_login" bson:"last_login"`
	Monitored         bool             `json:"monitored" bson:"monitored"`
	Name              string           `json:"name" bson:"name"`
	LastName          string           `json:"last_name" bson:"last_name"`
	Nick              string           `json:"nick" bson:"nick"`
	Password          string           `json:"password" bson:"password"`
	Realm             RealmT           `json:"realm" bson:"realm"`
	Token             string           `json:"token" bson:"token"`
	UserRol           string           `json:"user_rol" bson:"user_rol"`
	CreationDate      string           `json:"creation_date" bson:"creation_date"`
	Modification_date string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy        ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy         string           `json:"created_by" bson:"created_by"`
}

type AccountType struct {
	ApprovedBy        []string         `json:"approved_by" bson:"approved_by"`
	ApproverOf        []string         `json:"approver_of" bson:"approver_of"`
	ExpirationDate    string           `json:"expiration_date" bson:"expiration_date"`
	Picture           string           `json:"picture" bson:"picture"`
	PictureUrl        string           `json:"picture_url" bson:"picture_url"`
	CreationDate      string           `json:"creation_date" bson:"creation_date"`
	Modification_date string           `json:"modification_date" bson:"modification_date"`
	ModifiedBy        ModificationList `json:"modified_by" bson:"modified_by"`
	CreatedBy         string           `json:"created_by" bson:"created_by"`
}

type UserDataDB []UserData

type CreateUserRequestBody struct {
	Email          string   `json:"email" bson:"email"`
	KnownHost      []string `json:"known_host" bson:"known_host"`
	Monitored      bool     `json:"monitored" bson:"monitored"`
	Name           string   `json:"name" bson:"name"`
	LastName       string   `json:"last_name" bson:"last_name"`
	Nick           string   `json:"nick" bson:"nick"`
	Password       string   `json:"password" bson:"password"`
	Token          string   `json:"token" bson:"token"`
	UserRol        string   `json:"user_rol" bson:"user_rol"`
	Apis           string   `json:"apis" bson:"apis"`
	Media          string   `json:"media" bson:"media"`
	Mecuate        string   `json:"mecuate" bson:"mecuate"`
	ApprovedBy     []string `json:"approved_by" bson:"approved_by"`
	ApproverOf     []string `json:"approver_of" bson:"approver_of"`
	ExpirationDate string   `json:"expiration_date" bson:"expiration_date"`
	Picture        string   `json:"picture" bson:"picture"`
	PictureUrl     string   `json:"picture_url" bson:"picture_url"`
}

type UpdateProfileRequestBody struct {
	Email          string   `json:"email" bson:"email"`
	KnownHost      []string `json:"known_host" bson:"known_host"`
	Monitored      bool     `json:"monitored" bson:"monitored"`
	Name           string   `json:"name" bson:"name"`
	LastName       string   `json:"last_name" bson:"last_name"`
	Nick           string   `json:"nick" bson:"nick"`
	Password       string   `json:"password" bson:"password"`
	Token          string   `json:"token" bson:"token"`
	UserRol        string   `json:"user_rol" bson:"user_rol"`
	Apis           string   `json:"apis" bson:"apis"`
	Media          string   `json:"media" bson:"media"`
	Mecuate        string   `json:"mecuate" bson:"mecuate"`
	ApprovedBy     []string `json:"approved_by" bson:"approved_by"`
	ApproverOf     []string `json:"approver_of" bson:"approver_of"`
	ExpirationDate string   `json:"expiration_date" bson:"expiration_date"`
	Picture        string   `json:"picture" bson:"picture"`
	PictureUrl     string   `json:"picture_url" bson:"picture_url"`
}
