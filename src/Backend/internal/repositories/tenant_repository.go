package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(tenant *models.Tenant) error
	FindByID(id string) (*models.Tenant, error)
	FindByOwnerID(ownerID string) ([]models.Tenant, error)
	FindAll() ([]models.Tenant, error)
	Update(tenant *models.Tenant) error
	Delete(id string) error
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

func (r *tenantRepository) FindByOwnerID(ownerID string) ([]models.Tenant, error) {
	var tenants []models.Tenant
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&tenants).Error
	return tenants, err
}
