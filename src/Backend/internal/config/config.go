package config

import (
	"os"
	"strconv"
)

type Config struct {
	// Database
	DatabaseURL string

	// App
	Port        string
	Environment string

	// Logging
	LogLevel    string
	LogFilePath string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://aluguei_user:password@localhost:5432/aluguei_db"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		LogFilePath: getEnv("LOG_FILE_PATH", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DatabaseConfig retorna configurações específicas do banco
func (c *Config) DatabaseConfig() struct {
	MaxIdleConns int
	MaxOpenConns int
} {
	return struct {
		MaxIdleConns int
		MaxOpenConns int
	}{
		MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
