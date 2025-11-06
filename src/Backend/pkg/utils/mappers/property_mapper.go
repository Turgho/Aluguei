package mappers

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/utils/dtos"
	"github.com/google/uuid"
)

func ToPropertyModel(dto *dtos.CreatePropertyRequest, ownerID uuid.UUID) *models.Property {
	return &models.Property{
		OwnerID:      ownerID,
		Title:        dto.Title,
		Description:  dto.Description,
		Address:      dto.Address,
		Number:       dto.Number,
		Complement:   dto.Complement,
		Neighborhood: dto.Neighborhood,
		City:         dto.City,
		State:        dto.State,
		ZipCode:      dto.ZipCode,
		Type:         models.PropertyType(dto.Type),
		Bedrooms:     dto.Bedrooms,
		Bathrooms:    dto.Bathrooms,
		Area:         dto.Area,
		RentAmount:   dto.RentAmount,
	}
}

func ToPropertyResponse(property *models.Property) *dtos.PropertyResponse {
	return &dtos.PropertyResponse{
		ID:           property.ID.String(),
		OwnerID:      property.OwnerID.String(),
		Title:        property.Title,
		Description:  property.Description,
		Address:      property.Address,
		Number:       property.Number,
		Complement:   property.Complement,
		Neighborhood: property.Neighborhood,
		City:         property.City,
		State:        property.State,
		ZipCode:      property.ZipCode,
		Type:         string(property.Type),
		Bedrooms:     property.Bedrooms,
		Bathrooms:    property.Bathrooms,
		Area:         property.Area,
		RentAmount:   property.RentAmount,
		Status:       string(property.Status),
		CreatedAt:    property.CreatedAt,
		UpdatedAt:    property.UpdatedAt,
	}
}

func ToPropertyModelFromUpdate(dto *dtos.UpdatePropertyRequest, property *models.Property) {
	if dto.Title != "" {
		property.Title = dto.Title
	}
	if dto.Description != "" {
		property.Description = dto.Description
	}
	if dto.Address != "" {
		property.Address = dto.Address
	}
	if dto.Number != "" {
		property.Number = dto.Number
	}
	if dto.Complement != "" {
		property.Complement = dto.Complement
	}
	if dto.Neighborhood != "" {
		property.Neighborhood = dto.Neighborhood
	}
	if dto.City != "" {
		property.City = dto.City
	}
	if dto.State != "" {
		property.State = dto.State
	}
	if dto.ZipCode != "" {
		property.ZipCode = dto.ZipCode
	}
	if dto.Type != "" {
		property.Type = models.PropertyType(dto.Type)
	}
	if dto.Bedrooms > 0 {
		property.Bedrooms = dto.Bedrooms
	}
	if dto.Bathrooms > 0 {
		property.Bathrooms = dto.Bathrooms
	}
	if dto.Area > 0 {
		property.Area = dto.Area
	}
	if dto.RentAmount > 0 {
		property.RentAmount = dto.RentAmount
	}
}
