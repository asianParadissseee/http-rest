package store

type Config struct {
	DB_URL string `toml:"DB_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
