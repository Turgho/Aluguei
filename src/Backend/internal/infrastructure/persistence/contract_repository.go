package persistence

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type contractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) repositories.ContractRepository {
	return &contractRepository{db: db}
}

func (r *contractRepository) Create(ctx context.Context, contract *entities.Contract) error {
	return r.db.WithContext(ctx).Create(contract).Error
}

func (r *contractRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Contract, error) {
	var contract entities.Contract
	err := r.db.WithContext(ctx).First(&contract, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (r *contractRepository) GetByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*entities.Contract, error) {
	var contracts []*entities.Contract
	err := r.db.WithContext(ctx).Where("property_id = ?", propertyID).Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) GetByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*entities.Contract, error) {
	var contracts []*entities.Contract
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) GetActiveByPropertyID(ctx context.Context, propertyID uuid.UUID) (*entities.Contract, error) {
	var contract entities.Contract
	err := r.db.WithContext(ctx).Where("property_id = ? AND status = ?", propertyID, entities.ContractStatusActive).First(&contract).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (r *contractRepository) GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Contract, int64, error) {
	var contracts []*entities.Contract
	var total int64

	query := r.db.WithContext(ctx).Model(&entities.Contract{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&contracts).Error
	return contracts, total, err
}

func (r *contractRepository) Update(ctx context.Context, contract *entities.Contract) error {
	return r.db.WithContext(ctx).Save(contract).Error
}

func (r *contractRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Contract{}, "id = ?", id).Error
}