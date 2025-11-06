package mappers

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/utils/dtos"
	"github.com/google/uuid"
)

func ToTenantModel(dto *dtos.CreateTenantRequest, ownerID uuid.UUID) *models.Tenant {
	return &models.Tenant{
		OwnerID: ownerID,
		Name:    dto.Name,
		Email:   dto.Email,
		Phone:   dto.Phone,
		CPF:     dto.CPF,
	}
}

func ToTenantResponse(tenant *models.Tenant) *dtos.TenantResponse {
	return &dtos.TenantResponse{
		ID:        tenant.ID.String(),
		OwnerID:   tenant.OwnerID.String(),
		Name:      tenant.Name,
		Email:     tenant.Email,
		Phone:     tenant.Phone,
		CPF:       tenant.CPF,
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
	}
}

func ToTenantModelFromUpdate(dto *dtos.UpdateTenantRequest, tenant *models.Tenant) {
	if dto.Name != "" {
		tenant.Name = dto.Name
	}
	if dto.Email != "" {
		tenant.Email = dto.Email
	}
	if dto.Phone != "" {
		tenant.Phone = dto.Phone
	}
	if dto.CPF != "" {
		tenant.CPF = dto.CPF
	}
}
