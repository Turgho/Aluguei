package persistence

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) repositories.PropertyRepository {
	return &propertyRepository{db: db}
}

func (r *propertyRepository) Create(ctx context.Context, property *entities.Property) error {
	return r.db.WithContext(ctx).Create(property).Error
}

func (r *propertyRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Property, error) {
	var property entities.Property
	err := r.db.WithContext(ctx).First(&property, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &property, nil
}

func (r *propertyRepository) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Property, error) {
	var properties []*entities.Property
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&properties).Error
	return properties, err
}

func (r *propertyRepository) GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Property, int64, error) {
	var properties []*entities.Property
	var total int64

	query := r.db.WithContext(ctx).Model(&entities.Property{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&properties).Error
	return properties, total, err
}

func (r *propertyRepository) Update(ctx context.Context, property *entities.Property) error {
	return r.db.WithContext(ctx).Save(property).Error
}

func (r *propertyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Property{}, "id = ?", id).Error
}