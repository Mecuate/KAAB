package models

type CreateNodeRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Bump        bool          `json:"bump" bson:"bump"`
	Status      string        `json:"status" bson:"status"`
}

type CreateContentRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	Deletes     []MAPDATA     `json:"deletes" bson:"deletes" default:"[]"`
	Appends     []MAPDATA     `json:"appends" bson:"appends" default:"[]"`
	Patches     []MAPDATA     `json:"patches" bson:"patches" default:"[]"`
	Version     string        `json:"version" bson:"version"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Schema      string        `json:"schema_ref" bson:"schema_ref"`
	Bump        bool          `json:"bump" bson:"bump"`
	Status      string        `json:"status" bson:"status"`
}

type CreateEndpointRequest struct {
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description" bson:"description"`
	Value       EndpointCode `json:"value" bson:"value"`
	Schema      string       `json:"schema_ref" bson:"schema_ref"`
	Bump        bool         `json:"bump" bson:"bump"`
	Status      string       `json:"status" bson:"status"`
}

type CreateSchemaRequest struct {
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Value       []interface{} `json:"value" bson:"value"`
	RefId       string        `json:"ref_id" bson:"ref_id"`
	Bump        bool          `json:"bump" bson:"bump"`
	Status      string        `json:"status" bson:"status"`
}

type CreateMediaRequest struct {
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Size        int64          `json:"size" bson:"size"`
	Ttype       string         `json:"type" bson:"type"`
	Duration    int64          `json:"duration" bson:"duration"`
	Dimensions  DimentionsType `json:"dimensions" bson:"dimensions"`
	Service     string         `json:"service" bson:"service"`
	Bump        bool           `json:"bump" bson:"bump"`
	Status      string         `json:"status" bson:"status"`
	ChallengeID string         `json:"challenge" bson:"challenge"`
}

type CreateInstanceRequest struct {
	Name    string   `json:"name" bson:"name"`
	Owner   string   `json:"owner" bson:"owner"`
	Admin   []string `json:"admin" bson:"admin"`
	Members []string `json:"members" bson:"members"`
	Status  string   `json:"status" bson:"status"`
	Bump    bool     `json:"bump" bson:"bump"`
}
