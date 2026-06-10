package config

import "os"

type Config struct {
	Addr string
}

func Load() Config {
	return Config{
		Addr: getenv("APP_ADDR", ":8080"),
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
