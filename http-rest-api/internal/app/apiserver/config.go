package apiserver

import "github.com/eighthGnom/http-rest-api/storage"

type Config struct {
	BindAddr      string          `toml:"bind_addr"`
	LoggerLevel   string          `toml:"logger_level"`
	StorageConfig *storage.Config `toml:"storage_config"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8080",
		LoggerLevel:   "debug",
		StorageConfig: storage.NewConfig(),
	}
}
