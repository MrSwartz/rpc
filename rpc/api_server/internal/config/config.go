package config

type APIServerConfig struct {
	Host string
	Path string
	Port string
}

func New() APIServerConfig {
	return APIServerConfig{
		Host: "127.0.0.1",
		Path: "/hostname",
		Port: "8080",
	}
}
