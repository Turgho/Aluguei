package persistence

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) repositories.OwnerRepository {
	return &ownerRepository{db: db}
}

func (r *ownerRepository) Create(ctx context.Context, owner *entities.Owner) error {
	return r.db.WithContext(ctx).Create(owner).Error
}

func (r *ownerRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Owner, error) {
	var owner entities.Owner
	err := r.db.WithContext(ctx).First(&owner, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &owner, nil
}

func (r *ownerRepository) GetByEmail(ctx context.Context, email string) (*entities.Owner, error) {
	var owner entities.Owner
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&owner).Error
	if err != nil {
		return nil, err
	}
	return &owner, nil
}

func (r *ownerRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Owner, int64, error) {
	var owners []*entities.Owner
	var total int64

	r.db.WithContext(ctx).Model(&entities.Owner{}).Count(&total)

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&owners).Error
	return owners, total, err
}

func (r *ownerRepository) Update(ctx context.Context, owner *entities.Owner) error {
	return r.db.WithContext(ctx).Save(owner).Error
}

func (r *ownerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Owner{}, "id = ?", id).Error
}