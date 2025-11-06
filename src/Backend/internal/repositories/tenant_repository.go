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

func (r *tenantRepository) Delete(id string) *errors.AppError {
	// Verificar se existem contratos ativos primeiro
	var contractCount int64
	err := r.db.Model(&models.Contract{}).
		Where("tenant_id = ? AND status = ?", id, "active").
		Count(&contractCount).Error

	if err != nil {
		return errors.NewDatabaseError("erro ao verificar contratos do inquilino", err)
	}

	if contractCount > 0 {
		return errors.NewBusinessRuleError("não é possível deletar inquilino com contratos ativos")
	}

	// Se não houver contratos ativos, usar o Delete do BaseRepository
	return r.BaseRepository.Delete(id)
}

func (r *tenantRepository) FindByOwnerID(ownerID string) ([]models.Tenant, *errors.AppError) {
	var tenants []models.Tenant
	err := r.db.Where("owner_id = ?", ownerID).Preload("Contracts").Find(&tenants).Error
	if err != nil {
		return nil, errors.NewDatabaseError("erro ao buscar inquilinos por proprietário", err)
	}
	return tenants, nil
}

func (r *tenantRepository) FindByID(id string) (*models.Tenant, *errors.AppError) {
	// Usar o método com preloads
	return r.BaseRepository.FindByIDWithPreloads(id, "Contracts")
}

func (r *tenantRepository) FindAll() ([]models.Tenant, *errors.AppError) {
	// Usar o método com preloads
	return r.BaseRepository.FindAllWithPreloads("Contracts")
}
