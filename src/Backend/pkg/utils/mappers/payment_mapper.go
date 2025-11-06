package mappers

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/pkg/utils/dtos"
	"github.com/google/uuid"
)

func ToPaymentModel(dto *dtos.CreatePaymentRequest) *models.Payment {
	contractID, _ := uuid.Parse(dto.ContractID)

	return &models.Payment{
		ContractID: contractID,
		DueDate:    dto.DueDate,
		Amount:     dto.Amount,
	}
}

func ToPaymentResponse(payment *models.Payment) *dtos.PaymentResponse {
	resp := &dtos.PaymentResponse{
		ID:             payment.ID.String(),
		ContractID:     payment.ContractID.String(),
		DueDate:        payment.DueDate,
		PaidDate:       payment.PaidDate,
		Amount:         payment.Amount,
		PaidAmount:     payment.PaidAmount,
		LateFee:        payment.LateFee,
		Method:         string(payment.Method),
		Status:         string(payment.Status),
		ReferenceMonth: payment.ReferenceMonth,
		CreatedAt:      payment.CreatedAt,
		UpdatedAt:      payment.UpdatedAt,
	}

	// Incluir contrato se estiver carregado
	if payment.Contract != nil {
		contractResp := ToContractResponse(payment.Contract)
		resp.Contract = *contractResp
	}

	return resp
}

func ToPaymentModelFromUpdate(dto *dtos.UpdatePaymentRequest, payment *models.Payment) {
	if dto.PaidAmount > 0 {
		payment.PaidAmount = dto.PaidAmount
	}
	if dto.PaidDate != nil {
		payment.PaidDate = dto.PaidDate
	}
	if dto.Method != "" {
		payment.Method = models.PaymentMethod(dto.Method)
	}
}
