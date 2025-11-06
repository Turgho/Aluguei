package mappers

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/utils/dtos"
	"github.com/google/uuid"
)

func ToContractModel(dto *dtos.CreateContractRequest) *models.Contract {
	propertyID, _ := uuid.Parse(dto.PropertyID)
	tenantID, _ := uuid.Parse(dto.TenantID)

	return &models.Contract{
		PropertyID:    propertyID,
		TenantID:      tenantID,
		StartDate:     dto.StartDate,
		EndDate:       dto.EndDate,
		MonthlyRent:   dto.MonthlyRent,
		DepositAmount: dto.DepositAmount,
		PaymentDueDay: dto.PaymentDueDay,
	}
}

func ToContractResponse(contract *models.Contract) *dtos.ContractResponse {
	resp := &dtos.ContractResponse{
		ID:            contract.ID.String(),
		PropertyID:    contract.PropertyID.String(),
		TenantID:      contract.TenantID.String(),
		StartDate:     contract.StartDate,
		EndDate:       contract.EndDate,
		MonthlyRent:   contract.MonthlyRent,
		DepositAmount: contract.DepositAmount,
		PaymentDueDay: contract.PaymentDueDay,
		Status:        string(contract.Status),
		CreatedAt:     contract.CreatedAt,
		UpdatedAt:     contract.UpdatedAt,
	}

	// Incluir propriedade e inquilino se estiverem carregados
	if contract.Property != nil {
		propertyResp := ToPropertyResponse(contract.Property)
		resp.Property = *propertyResp
	}

	if contract.Tenant != nil {
		tenantResp := ToTenantResponse(contract.Tenant)
		resp.Tenant = *tenantResp
	}

	return resp
}

func ToContractModelFromUpdate(dto *dtos.UpdateContractRequest, contract *models.Contract) {
	if !dto.EndDate.IsZero() {
		contract.EndDate = dto.EndDate
	}
	if dto.MonthlyRent > 0 {
		contract.MonthlyRent = dto.MonthlyRent
	}
	if dto.PaymentDueDay > 0 {
		contract.PaymentDueDay = dto.PaymentDueDay
	}
}
