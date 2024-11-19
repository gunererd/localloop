package config

import (
	"os"
)

type Config struct {
	Port        string
	PostgresURI string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8082"),
		PostgresURI: getEnv("POSTGRES_URI", "postgres://postgres:password@localhost:5432/catalog?sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
