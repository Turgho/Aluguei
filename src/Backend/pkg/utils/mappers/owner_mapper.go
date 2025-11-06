package mappers

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/utils/dtos"
)

func ToOwnerModel(dto *dtos.CreateOwnerRequest) *models.Owner {
	return &models.Owner{
		Name:     dto.Name,
		Email:    dto.Email,
		Phone:    dto.Phone,
		Password: dto.Password,
	}
}

func ToOwnerResponse(owner *models.Owner) *dtos.OwnerResponse {
	return &dtos.OwnerResponse{
		ID:        owner.ID.String(),
		Name:      owner.Name,
		Email:     owner.Email,
		Phone:     owner.Phone,
		CreatedAt: owner.CreatedAt,
		UpdatedAt: owner.UpdatedAt,
	}
}

func ToOwnerModelFromUpdate(dto *dtos.UpdateOwnerRequest, owner *models.Owner) {
	if dto.Name != "" {
		owner.Name = dto.Name
	}
	if dto.Phone != "" {
		owner.Phone = dto.Phone
	}
	if dto.Password != "" {
		owner.Password = dto.Password
	}
}
