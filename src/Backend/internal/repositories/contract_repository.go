package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type ContractRepository interface {
	Create(contract *models.Contract) *errors.AppError
	FindByID(id string) (*models.Contract, *errors.AppError)
	FindByPropertyID(propertyID string) ([]models.Contract, *errors.AppError)
	FindByTenantID(tenantID string) ([]models.Contract, *errors.AppError)
	FindActiveByPropertyID(propertyID string) (*models.Contract, *errors.AppError)
	FindAll() ([]models.Contract, *errors.AppError)
	Update(contract *models.Contract) *errors.AppError
	Delete(id string) *errors.AppError
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

func (r *contractRepository) FindByPropertyID(propertyID string) ([]models.Contract, *errors.AppError) {
	var contracts []models.Contract
	err := r.db.Where("property_id = ?", propertyID).
		Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		Find(&contracts).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar contratos por propriedade", err)
	}
	return contracts, nil
}

func (r *contractRepository) FindByTenantID(tenantID string) ([]models.Contract, *errors.AppError) {
	var contracts []models.Contract
	err := r.db.Where("tenant_id = ?", tenantID).
		Preload("Property").
		Preload("Payments").
		Find(&contracts).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar contratos por inquilino", err)
	}
	return contracts, nil
}

func (r *contractRepository) FindActiveByPropertyID(propertyID string) (*models.Contract, *errors.AppError) {
	var contract models.Contract
	err := r.db.Where("property_id = ? AND status = ?", propertyID, models.ContractStatusActive).
		Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		First(&contract).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("contrato ativo para propriedade", propertyID)
		}
		return nil, errors.NewDatabaseError("erro ao buscar contrato ativo por propriedade", err)
	}
	return &contract, nil
}

// Sobrescrever FindByID para incluir preloads
func (r *contractRepository) FindByID(id string) (*models.Contract, *errors.AppError) {
	var contract models.Contract
	err := r.db.Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		First(&contract, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("contrato", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar contrato por ID", err)
	}
	return &contract, nil
}

// Sobrescrever FindAll para incluir preloads
func (r *contractRepository) FindAll() ([]models.Contract, *errors.AppError) {
	var contracts []models.Contract
	err := r.db.Preload("Property").
		Preload("Tenant").
		Preload("Payments").
		Find(&contracts).Error

	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todos os contratos", err)
	}
	return contracts, nil
}
