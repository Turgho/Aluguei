package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(tenant *models.Tenant) *errors.AppError
	FindByID(id string) (*models.Tenant, *errors.AppError)
	FindByOwnerID(ownerID string) ([]models.Tenant, *errors.AppError)
	FindAll() ([]models.Tenant, *errors.AppError)
	Update(tenant *models.Tenant) *errors.AppError
	Delete(id string) *errors.AppError
}

type tenantRepository struct {
	*BaseRepository[models.Tenant]
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		BaseRepository: NewBaseRepository[models.Tenant](db),
		db:             db,
	}
}

func (r *tenantRepository) FindByOwnerID(ownerID string) ([]models.Tenant, *errors.AppError) {
	var tenants []models.Tenant
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&tenants).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar inquilinos por proprietário", err)
	}
	return tenants, nil
}

// Sobrescrever FindByID para incluir preloads se necessário
func (r *tenantRepository) FindByID(id string) (*models.Tenant, *errors.AppError) {
	var tenant models.Tenant
	err := r.db.Preload("Contracts").First(&tenant, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("inquilino", id)
		}
		return nil, errors.NewDatabaseError("erro ao buscar inquilino por ID", err)
	}
	return &tenant, nil
}

// Sobrescrever FindAll para incluir preloads se necessário
func (r *tenantRepository) FindAll() ([]models.Tenant, *errors.AppError) {
	var tenants []models.Tenant
	err := r.db.Preload("Contracts").Find(&tenants).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar todos os inquilinos", err)
	}
	return tenants, nil
}
