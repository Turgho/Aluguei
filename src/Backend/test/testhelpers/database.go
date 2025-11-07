package testhelpers

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB creates an in-memory SQLite database for testing
func SetupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate all entities
	err = db.AutoMigrate(
		&entities.Owner{},
		&entities.Tenant{},
		&entities.Property{},
		&entities.Contract{},
		&entities.Payment{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CleanupTestDB cleans all tables in the test database
func CleanupTestDB(db *gorm.DB) error {
	tables := []string{"payments", "contracts", "properties", "tenants", "owners"}
	
	for _, table := range tables {
		if err := db.Exec("DELETE FROM " + table).Error; err != nil {
			return err
		}
	}
	
	return nil
}