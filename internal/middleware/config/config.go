package config

import (
	"os"
)

type Config struct {
	ServerAddr       string
	MetricsAddr      string
	FinamBase        string
	FinamAccessToken string
	RateLimit        int
}

func LoadConfig() *Config {
	// Запиши дефолты либо забирай из ENV
	return &Config{
		ServerAddr:       env("SERVER_ADDR", ":8080"),
		MetricsAddr:      env("METRICS_ADDR", ":9090"),
		FinamBase:        env("FINAM_API_BASE_URL", "https://api.finam.ru"),
		FinamAccessToken: os.Getenv("FINAM_ACCESS_TOKEN"),
		RateLimit:        100,
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
