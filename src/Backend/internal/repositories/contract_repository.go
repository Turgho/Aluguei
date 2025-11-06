package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type ContractRepository interface {
	Create(contract *models.Contract) error
	FindByID(id string) (*models.Contract, error)
	FindByPropertyID(propertyID string) ([]models.Contract, error)
	FindByTenantID(tenantID string) ([]models.Contract, error)
	FindActiveByPropertyID(propertyID string) (*models.Contract, error)
	FindAll() ([]models.Contract, error)
	Update(contract *models.Contract) error
	Delete(id string) error
}

type contractRepository struct {
	*BaseRepository[models.Contract]
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) ContractRepository {
	return &contractRepository{
		BaseRepository: NewBaseRepository[models.Contract](db),
		db:             db,
	}
}

func (r *contractRepository) FindByPropertyID(propertyID string) ([]models.Contract, error) {
	var contracts []models.Contract
	err := r.db.Where("property_id = ?", propertyID).
		Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) FindByTenantID(tenantID string) ([]models.Contract, error) {
	var contracts []models.Contract
	err := r.db.Where("tenant_id = ?", tenantID).
		Preload("Property").
		Preload("Payments").
		Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) FindActiveByPropertyID(propertyID string) (*models.Contract, error) {
	var contract models.Contract
	err := r.db.Where("property_id = ? AND status = ?", propertyID, models.ContractStatusActive).
		Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		First(&contract).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
