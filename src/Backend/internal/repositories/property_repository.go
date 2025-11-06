package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type PropertyRepository interface {
	Create(property *models.Property) error
	FindByID(id string) (*models.Property, error)
	FindByOwnerID(ownerID string) ([]models.Property, error)
	FindAll() ([]models.Property, error)
	Update(property *models.Property) error
	Delete(id string) error
}

type propertyRepository struct {
	*BaseRepository[models.Property]
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &propertyRepository{
		BaseRepository: NewBaseRepository[models.Property](db),
		db:             db,
	}
}

func (r *propertyRepository) FindByOwnerID(ownerID string) ([]models.Property, error) {
	var properties []models.Property
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&properties).Error
	return properties, err
}
