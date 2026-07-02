package config

import (
	"os"
)

type Config struct {
	HTTPAddr    string
	Env         string
	DatabaseURL string
	JWTSecret 	string
}

// new config
func Load() Config {
	return Config{
		HTTPAddr:    getEnv("HTTP_ADDR", ":8080"),
		Env:         getEnv("APP_ENV", "local"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret: 	 getEnv("JWT_SECRET", "dev-secret-change-me"),
	}
}

// fallback are default value
// getEnv get env values from the systemes (terminal/docker/...)
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
