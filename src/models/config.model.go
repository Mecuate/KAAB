package models

type WebConfigs struct {
	Port                 string `required:"true" split_words:"true" default:"11112"`
	CorsEnabled          string `required:"true" split_words:"true" default:"true"`
	Environment          string `required:"true" split_words:"true"`
	ApiVersions          string `required:"true" split_words:"true"`
	ApiPublishedVersions string `required:"true" split_words:"true"`
	PubDbName            string `required:"true" split_words:"true"`
	IntDbName            string `required:"true" split_words:"true"`
	Mongodburi           string `required:"true" split_words:"true"`
	UrlAddress           string `required:"true" split_words:"true"`
	UriAddress           string `required:"true" split_words:"true"`
	Thumbs               string `required:"true" split_words:"true"`
	PhysicalName         string `required:"true" split_words:"true"`
}

type LoggingConfig struct {
	LogPath       string `required:"true" split_words:"true" default:"/home/devops/logs/kaab"`
	LogFileName   string `required:"true" split_words:"true" default:"server.log"`
	ErrorFileName string `required:"true" split_words:"true" default:"server_error.log"`
}

type AppConfig struct {
	ProcessName string `required:"true" split_words:"true"`
	Copyright   string `required:"true" split_words:"false"`
}

type EnvConfigs struct {
	WebServerConfig *WebConfigs
	LoggingConfig   *LoggingConfig
	EnvConfig       *AppConfig
}
