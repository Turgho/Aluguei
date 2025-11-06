package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type PropertyRepository interface {
	Create(property *models.Property) *errors.AppError
	FindByID(id string) (*models.Property, *errors.AppError)
	FindByOwnerID(ownerID string) ([]models.Property, *errors.AppError)
	FindAll() ([]models.Property, *errors.AppError)
	Update(property *models.Property) *errors.AppError
	Delete(id string) *errors.AppError
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

func (r *propertyRepository) FindByOwnerID(ownerID string) ([]models.Property, *errors.AppError) {
	var properties []models.Property
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&properties).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar propriedades por proprietário", err)
	}
	return properties, nil
}

// Sobrescrever FindByID para incluir preloads
func (r *propertyRepository) FindByID(id string) (*models.Property, *errors.AppError) {
	var property models.Property
	err := r.db.Preload("Contracts").First(&property, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("propriedade", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar propriedade por ID", err)
	}
	return &property, nil
}

// Sobrescrever FindAll para incluir preloads
func (r *propertyRepository) FindAll() ([]models.Property, *errors.AppError) {
	var properties []models.Property
	err := r.db.Preload("Contracts").Find(&properties).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todas as propriedades", err)
	}
	return properties, nil
}
