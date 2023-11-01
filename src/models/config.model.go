package models

type WebConfigs struct {
	Port        string `required:"true" split_words:"true" default:"11112"`
	CorsEnabled string `required:"true" split_words:"true" default:"true"`
	Environment string `required:"true" split_words:"true"`
	ApiVersions string `required:"true" split_words:"true"`
}

type LoggingConfig struct {
	LogPath       string `required:"true" split_words:"true" default:"/home/devops/logs/kaab"`
	LogFileName   string `required:"true" split_words:"true" default:"server.log"`
	ErrorFileName string `required:"true" split_words:"true" default:"server_error.log"`
}

type AppConfig struct {
	ProcessName string `required:"true" split_words:"true"`
	DbDir       string `required:"true" split_words:"true"`
	Copyright   string `required:"true" split_words:"false"`
}

type EnvConfigs struct {
	WebServerConfig *WebConfigs
	LoggingConfig   *LoggingConfig
	EnvConfig       *AppConfig
}
