package config

import (
	"os"
	"strconv"
)

type Config struct {
	// Server configs
	Port string

	// Database configs
	MongoURI string

	// JWT configs
	JWTSecret            string
	JWTExpirationMinutes int

	// Security configs
	SaltLength int
}

func Load() *Config {
	return &Config{
		Port:                 getEnv("PORT", "8080"),
		MongoURI:             getEnv("MONGO_URI", "mongodb://admin:password@localhost:27017/user-management?authSource=admin"),
		JWTSecret:            getEnv("JWT_SECRET", "your-default-secret-key"),
		JWTExpirationMinutes: getEnvAsInt("JWT_EXPIRATION_MINUTES", 60*24), // 24 hours
		SaltLength:           getEnvAsInt("SALT_LENGTH", 25),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
