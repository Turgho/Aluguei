package database

import (
	"fmt"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Connect(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate
	if err := db.AutoMigrate(
		&entities.Owner{},
		&entities.Tenant{},
		&entities.Property{},
		&entities.Contract{},
		&entities.Payment{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
