package usecases

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
)

type PropertyUseCase struct {
	propertyRepo repositories.PropertyRepository
}

func NewPropertyUseCase(propertyRepo repositories.PropertyRepository) *PropertyUseCase {
	return &PropertyUseCase{
		propertyRepo: propertyRepo,
	}
}

func (uc *PropertyUseCase) CreateProperty(ctx context.Context, ownerID uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64) (*entities.Property, error) {
	property := entities.NewProperty(ownerID, title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount)
	
	if err := uc.propertyRepo.Create(ctx, property); err != nil {
		return nil, err
	}
	
	return property, nil
}

func (uc *PropertyUseCase) GetProperty(ctx context.Context, id uuid.UUID) (*entities.Property, error) {
	return uc.propertyRepo.GetByID(ctx, id)
}

func (uc *PropertyUseCase) GetPropertiesByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entities.Property, error) {
	return uc.propertyRepo.GetByOwnerID(ctx, ownerID)
}

func (uc *PropertyUseCase) GetAllProperties(ctx context.Context, page, limit int, status string) ([]*entities.Property, int64, error) {
	return uc.propertyRepo.GetAll(ctx, page, limit, status)
}

func (uc *PropertyUseCase) UpdateProperty(ctx context.Context, id uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64, status entities.PropertyStatus) error {
	property, err := uc.propertyRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	property.Update(title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount, status)
	
	return uc.propertyRepo.Update(ctx, property)
}

func (uc *PropertyUseCase) DeleteProperty(ctx context.Context, id uuid.UUID) error {
	return uc.propertyRepo.Delete(ctx, id)
}