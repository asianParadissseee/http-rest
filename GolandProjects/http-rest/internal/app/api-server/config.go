package api_server

type Config struct {
	BindAddr string `toml:"bind_addr"'` // адрес запуска сервера
	LogLever string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLever: "debug",
	}
}
