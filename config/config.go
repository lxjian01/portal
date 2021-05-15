package config

type AppConfig struct {
	Version    string        `yaml:"version"`
	Env        string        `yaml:"env"`
	Httpd      *HttpdConfig   `yaml:"httpd"`
	RecordDir  string        `yaml:"recordDir"`
	LinuxUser  string        `yaml:"linuxUser"`
	PoolNum    int           `yaml:"poolNum"`
	Log        *LogConfig     `yaml:"log"`
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