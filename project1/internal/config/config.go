package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr    string `yaml:"addr" env:"addr" env-required:"true" env-default:"localhost:8080"`
	Timeout int    `yaml:"timeout" env:"http_timeout" env-required:"true" env-default:"40"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"db_host" env-default:"localhost"`
	Port     int    `yaml:"port" env:"db_port" env-default:"5432"`
	User     string `yaml:"user" env:"db_user" env-default:"postgres"`
	Password string `yaml:"password" env:"db_password" env-default:"password"`
	URL      string `yaml:"url" env:"db_url" env-default:"postgresql://postgres:password@localhost/postgres"`
}

type Config struct {
	Env            string         `yaml:"env" env:"env" env-required:"true" env-default:"production"`
	HTTPServer     HTTPServer     `yaml:"http"`
	StoragePath    string         `yaml:"storage_path" env:"storage_path" env-required:"true" env-default:"./storage"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

// MustLoad reads configuration from YAML file and environment variables.
// It returns a Config pointer and an error if the configuration cannot be loaded.
func MustLoad(configPath string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
