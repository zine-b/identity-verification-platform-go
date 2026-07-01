package config

import(
	"os"
)

type Config struct{
	HTTPAddr string
	Env      string
}

// new config
func Load() Config {
	return Config{
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
		Env:      getEnv("APP_ENV", "local"),
	}
}

// fallback default value
func getEnv(key, fallback string)string{
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}