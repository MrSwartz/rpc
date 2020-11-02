package config

type SQLConfig struct {
	Port     string
	User     string
	Driver   string
	Password string
	Database string
}

func New() *SQLConfig {
	return &SQLConfig{
		Port:     "3306",
		Database: "data",
		Password: "passwd",
		User:     "qulix",
		Driver:   "mysql",
	}
}

type HttpServerConfig struct {
	Port  string
	Proto string
}

func NewHttp() *HttpServerConfig {
	return &HttpServerConfig{
		Port:  ":1234",
		Proto: "tcp",
	}
}
