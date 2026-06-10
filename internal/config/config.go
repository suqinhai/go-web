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
	JWT    JWT    `yaml:"jwt"`
	Casbin Casbin `yaml:"casbin"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type Log struct {
	Path string `yaml:"path"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
}

type Casbin struct {
	ModelPath  string `yaml:"model_path"`
	PolicyPath string `yaml:"policy_path"`
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
		JWT: JWT{
			Secret: "change-me",
			Issuer: "go-web",
		},
		Casbin: Casbin{
			ModelPath:  "configs/rbac_model.conf",
			PolicyPath: "configs/rbac_policy.csv",
		},
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
