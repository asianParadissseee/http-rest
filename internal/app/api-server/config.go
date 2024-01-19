package api_server

import "github.com/asianParadissseee/http-rest/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"'` // адрес запуска сервера
	LogLever string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLever: "debug",
		Store:    store.NewConfig(),
	}
}
