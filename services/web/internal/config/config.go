package config

import (
	"os"
)

type Config struct {
	Port              string
	UserServiceURL    string
	CatalogServiceURL string
}

func Load() *Config {
	return &Config{
		Port:              getEnv("PORT", "3000"),
		UserServiceURL:    getEnv("USER_SERVICE_URL", "http://localhost:8080"),
		CatalogServiceURL: getEnv("CATALOG_SERVICE_URL", "http://localhost:8081"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
