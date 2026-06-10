package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

const DefaultPath = "configs/config.yaml"

type Config struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type Log struct {
	Path string `yaml:"path"`
}

func Load() (Config, error) {
	path := getenv("APP_CONFIG", DefaultPath)
	cfg := defaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config file %q: %w", path, err)
	}

	if addr := os.Getenv("APP_ADDR"); addr != "" {
		cfg.Server.Addr = addr
	}

	return cfg, nil
}

func defaultConfig() Config {
	return Config{
		Server: Server{
			Addr: ":8080",
		},
		Log: Log{
			Path: "logs/app.log",
		},
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
