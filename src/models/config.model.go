package models

type WebConfigs struct {
	Port        string `required:"true" split_words:"true" default:"11112"`
	SSLCertPath string `required:"true" split_words:"true" default:"/home/mecuateConf/certs/kaab.crt"`
	SSLKeyPath  string `required:"true" split_words:"true" default:"/home/mecuateConf/certs/kaab.key"`
	BaseApiUrl  string `required:"true" split_words:"true" default:"kaab.mecuate.org/api"`
	CorsEnabled bool   `required:"true" split_words:"true"`
	EnabledAuth bool
	Enviroment  string `required:"true" split_words:"true"`
}

type LoggingConfig struct {
	logPath       string `required:"true" split_words:"true" default:"/home/devops/logs/kaab/kaab.log"`
	logFileName   string `required:"true" split_words:"true" default:"/home/devops/logs/kaab/server.log"`
	errorFileName string `required:"true" split_words:"true" default:"/home/devops/logs/kaab/server_error.log"`
}

type EnvConfs struct {
	ProcessName  string `required:"true" split_words:"true"`
	ProcessSufix string `required:"true" split_words:"true"`
	Hmac         string `required:"true" split_words:"true"`
	Secret       string `required:"true" split_words:"true"`
	Copyright    string `required:"true" split_words:"false"`
}

type ServiceConfig struct {
	WebServerConfig *WebConfigs
	LoggingConfig   *LoggingConfig
	EnvConfig       *EnvConfs
}

type CLIflags struct {
	PORT   string
	NAME   string
	LOGDIR string
}
