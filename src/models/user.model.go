package models

type RealmT struct {
	Apis    string `json:"apis"`
	Media   string `json:"media"`
	Mecuate string `json:"mecuate"`
}

type UserData struct {
	AccessToken       string      `json:"access_token"`
	Account           AccountType `json:"account"`
	CreationDate      int64       `json:"creation_date"`
	Email             string      `json:"email"`
	Id                string      `json:"id"`
	KnownHost         []string    `json:"known_host"`
	LastLogin         int64       `json:"last_login"`
	Modification_date int64       `json:"modification_date"`
	Monitored         bool        `json:"monitored"`
	Name              string      `json:"name"`
	LastName          string      `json:"last_name"`
	Nick              string      `json:"nick"`
	Password          string      `json:"password"`
	Realm             RealmT      `json:"realm"`
	Token             string      `json:"token"`
	UserRol           string      `json:"user_rol"`
}

type AccountType struct {
	ApprovedBy        []string `json:"approved_by"`
	ApproverOf        []string `json:"approver_of"`
	CreationDate      int64    `json:"creation_date"`
	ExpirationDate    int64    `json:"expiration_date"`
	Modification_date int64    `json:"modification_date"`
	Picture           string   `json:"picture"`
	PictureUrl        string   `json:"picture_url"`
}

type UserDataDB []UserData
