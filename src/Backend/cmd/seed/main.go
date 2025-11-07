package main

import (
	"log"
	"os"

	"github.com/Turgho/Aluguei/internal/infrastructure/database"
	"github.com/Turgho/Aluguei/internal/infrastructure/seeds"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Database configuration
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", ""),
		Port:     getEnv("DB_PORT", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Connect to database
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully")

	// Run seeds
	seeder := seeds.NewSeeder(db)

	log.Println("Starting database seeding...")

	if err := seeder.SeedAll(); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	log.Println("Database seeded successfully!")
	log.Println("Sample data created:")
	log.Println("- 3 Owners (password: 123456)")
	log.Println("- 3 Tenants")
	log.Println("- 4 Properties")
	log.Println("- 2 Contracts")
	log.Println("- 4 Payments")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
