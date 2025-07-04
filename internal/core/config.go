// File: internal/config/config.go
// Purpose: Load and parse application-wide configuration from environment variables

package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	Debug           bool
	DatabaseURL     string
	RedisURL        string
	SecretKey       string
	EnableRateLimit bool
	Env             string
}

func Load() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		Port:            getEnv("PORT", "8080"),
		Debug:           getEnvBool("DEBUG", false),
		Env:             getEnv("ENV", "production"),
		RedisURL:        getEnv("REDIS_URL", ""),
		EnableRateLimit: getEnvBool("ENABLE_RATE_LIMIT", true),
		SecretKey:       getEnv("SECRET_KEY", "shorty-secret"),
		DatabaseURL:     getEnv("DATABASE_URL", "sqlite://data.db"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if val := os.Getenv(key); val != "" {
		b, err := strconv.ParseBool(val)
		if err != nil {
			log.Printf("invalid boolean for %s: %v", key, err)
			return fallback
		}
		return b
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Printf("invalid duration for %s: %v", key, err)
			return fallback
		}
		return d
	}
	return fallback
}
