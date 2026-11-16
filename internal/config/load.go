package config

import (
	"log/slog"
	"os"

	"github.com/go-yaml/yaml"
)

func Load() *Config {
	var config Config

	configFilePath := getEnvOrDefault("CONFIG_FILE_PATH", "/etc/uptime-phoenix/config.yaml")
	configFile, err := os.Open(configFilePath)
	if err != nil {
		slog.Warn("Error opening config file", "path", configFilePath, "error", err)
	} else {
		defer configFile.Close()
		if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
			slog.Warn("Error decoding config file", "path", configFilePath, "error", err)
		}
	}

	return &config
}

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
