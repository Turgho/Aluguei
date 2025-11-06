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

func (r *propertyRepository) Delete(id string) *errors.AppError {
	// Verificar se existem contratos ativos primeiro
	var contractCount int64
	err := r.db.Model(&models.Contract{}).
		Where("property_id = ? AND status = ?", id, "active").
		Count(&contractCount).Error

	if err != nil {
		return errors.NewDatabaseError("erro ao verificar contratos da propriedade", err)
	}

	if contractCount > 0 {
		return errors.NewBusinessRuleError("não é possível deletar propriedade com contratos ativos")
	}

	// Se não houver contratos ativos, usar o Delete do BaseRepository
	return r.BaseRepository.Delete(id)
}

func (r *propertyRepository) FindByOwnerID(ownerID string) ([]models.Property, *errors.AppError) {
	var properties []models.Property
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&properties).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar propriedades por proprietário", err)
	}
	return properties, nil
}

func (r *propertyRepository) FindByID(id string) (*models.Property, *errors.AppError) {
	// Usar o método com preloads
	return r.BaseRepository.FindByIDWithPreloads(id, "Contracts")
}

func (r *propertyRepository) FindAll() ([]models.Property, *errors.AppError) {
	// Usar o método com preloads
	return r.BaseRepository.FindAllWithPreloads("Contracts")
}
