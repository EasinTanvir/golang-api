package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type HTTPServer struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Env         string     `yaml:"env" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

// MustLoad loads the config file and returns a Config instance.
// It will stop the application if any error occurs.
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	return &cfg
}