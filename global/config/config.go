package config

type AppConfig struct {
	Version    string         `yaml:"version"`
	Env        string         `yaml:"env"`
	Httpd      *HttpdConfig   `yaml:"web"`
	Log        *LogConfig     `yaml:"log"`
	Mysql      *MysqlConfig   `yaml:"mysql"`
	Consul *ConsulConfig   `yaml:"consul"`
}

type HttpdConfig struct {
	Host string
	Port int
}

type LogConfig struct {
	Dir       string
	Name      string
	Format    string
	RetainDay int8
	Level     string
}

type MysqlConfig struct {
	Host        string
	Port        int
	Db          string
	User        string
	Password    string
	Charset     string
}

type ConsulConfig struct {
	Host        string
	Port        int
}

var conf *AppConfig

func SetAppConfig(c *AppConfig){
	conf = c
}

func GetAppConfig() *AppConfig {
	return conf
}