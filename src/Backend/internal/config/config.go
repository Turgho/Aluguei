package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	fmt.Println("🔍 Iniciando carregamento de configurações...")

	// Tentar carregar .env
	if err := godotenv.Load(); err != nil {
		fmt.Printf("❌ Erro ao carregar .env: %v\n", err)
	} else {
		fmt.Println("✅ .env carregado com sucesso")
	}

	// Debug: mostrar variáveis carregadas
	fmt.Printf("📋 DATABASE_URL: %s\n", os.Getenv("DATABASE_URL"))
	fmt.Printf("📋 PORT: %s\n", os.Getenv("PORT"))
	fmt.Printf("📋 ENVIRONMENT: %s\n", os.Getenv("ENVIRONMENT"))

	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", ""),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		LogFilePath: getEnv("LOG_FILE_PATH", ""),
	}

	fmt.Printf("🎯 Config carregada - DatabaseURL: %s\n", cfg.DatabaseURL)
	return cfg
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
	if value := getEnv(key, ""); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// Validate verifica se as configurações essenciais estão presentes
func (c *Config) Validate() error {
	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL não configurada. Verifique o arquivo .env")
	}
	return nil
}
