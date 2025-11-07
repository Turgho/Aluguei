package persistence

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) repositories.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(ctx context.Context, tenant *entities.Tenant) error {
	return r.db.WithContext(ctx).Create(tenant).Error
}

func (r *tenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Tenant, error) {
	var tenant entities.Tenant
	err := r.db.WithContext(ctx).First(&tenant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Tenant, error) {
	var tenants []*entities.Tenant
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&tenants).Error
	return tenants, err
}

func (r *tenantRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Tenant, int64, error) {
	var tenants []*entities.Tenant
	var total int64

	r.db.WithContext(ctx).Model(&entities.Tenant{}).Count(&total)

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&tenants).Error
	return tenants, total, err
}

func (r *tenantRepository) Update(ctx context.Context, tenant *entities.Tenant) error {
	return r.db.WithContext(ctx).Save(tenant).Error
}

func (r *tenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Tenant{}, "id = ?", id).Error
}