package config

import (
	"os"
	"strings"
)

type Config struct {
	ServiceName string
	Port        string
	PostgresDSN string
	RedisAddr   string
}

func Load(serviceName, defaultPort string) Config {
	// Prefer POSTGRES_DSN, fall back to POSTGRES_URL for compatibility with docker-compose.
	dsn := getEnv("POSTGRES_DSN", "")
	if dsn == "" {
		dsn = getEnv("POSTGRES_URL", "")
	}
	if dsn == "" {
		dsn = "postgres://kodrapay:kodrapay_password@postgres:5432/kodrapay?sslmode=disable"
	}

	// Ensure sslmode=disable is present for local/docker Postgres
	if !strings.Contains(strings.ToLower(dsn), "sslmode=") {
		if strings.Contains(dsn, "?") {
			dsn += "&sslmode=disable"
		} else {
			dsn += "?sslmode=disable"
		}
	}

	return Config{
		ServiceName: serviceName,
		Port:        getEnv("PORT", defaultPort),
		PostgresDSN: dsn,
		RedisAddr:   getEnv("REDIS_ADDR", "redis:6379"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
