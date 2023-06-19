package config

type Config struct {
	DSN string
}

func NewConfig() *Config {
	return &Config{
		DSN: "abc",
	}
}
