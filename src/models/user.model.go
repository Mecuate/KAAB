package models

type RealmT struct {
	Apis    string `json:"apis" bson:"apis"`
	Media   string `json:"media" bson:"media"`
	Mecuate string `json:"mecuate" bson:"mecuate"`
}

type UserData struct {
	AccessToken       string      `json:"access_token" bson:"access_token"`
	Account           AccountType `json:"account" bson:"account"`
	CreationDate      int64       `json:"creation_date" bson:"creation_date"`
	Email             string      `json:"email" bson:"email"`
	Id                string      `json:"id" bson:"id"`
	KnownHost         []string    `json:"known_host" bson:"known_host"`
	LastLogin         int64       `json:"last_login" bson:"last_login"`
	Modification_date int64       `json:"modification_date" bson:"modification_date"`
	Monitored         bool        `json:"monitored" bson:"monitored"`
	Name              string      `json:"name" bson:"name"`
	LastName          string      `json:"last_name" bson:"last_name"`
	Nick              string      `json:"nick" bson:"nick"`
	Password          string      `json:"password" bson:"password"`
	Realm             RealmT      `json:"realm" bson:"realm"`
	Token             string      `json:"token" bson:"token"`
	UserRol           string      `json:"user_rol" bson:"user_rol"`
}

type AccountType struct {
	ApprovedBy        []string `json:"approved_by" bson:"approved_by"`
	ApproverOf        []string `json:"approver_of" bson:"approver_of"`
	CreationDate      int64    `json:"creation_date" bson:"creation_date"`
	ExpirationDate    int64    `json:"expiration_date" bson:"expiration_date"`
	Modification_date int64    `json:"modification_date" bson:"modification_date"`
	Picture           string   `json:"picture" bson:"picture"`
	PictureUrl        string   `json:"picture_url" bson:"picture_url"`
}

type UserDataDB []UserData
