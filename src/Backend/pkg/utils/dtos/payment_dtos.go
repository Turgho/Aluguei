package dtos

import "time"

type CreatePaymentRequest struct {
	ContractID string    `json:"contract_id" validate:"required,uuid4"`
	DueDate    time.Time `json:"due_date" validate:"required"`
	Amount     float64   `json:"amount" validate:"required,min=0"`
}

type UpdatePaymentRequest struct {
	PaidAmount float64    `json:"paid_amount" validate:"omitempty,min=0"`
	PaidDate   *time.Time `json:"paid_date"`
	Method     string     `json:"method" validate:"omitempty,oneof=pix bank_transfer cash other"`
}

type ProcessPaymentRequest struct {
	PaidAmount float64   `json:"paid_amount" validate:"required,min=0"`
	Method     string    `json:"method" validate:"required,oneof=pix bank_transfer cash other"`
	PaidDate   time.Time `json:"paid_date" validate:"required"`
}

type PaymentResponse struct {
	ID             string           `json:"id"`
	ContractID     string           `json:"contract_id"`
	DueDate        time.Time        `json:"due_date"`
	PaidDate       *time.Time       `json:"paid_date"`
	Amount         float64          `json:"amount"`
	PaidAmount     float64          `json:"paid_amount"`
	LateFee        float64          `json:"late_fee"`
	Method         string           `json:"method"`
	Status         string           `json:"status"`
	ReferenceMonth time.Time        `json:"reference_month"`
	Contract       ContractResponse `json:"contract,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}
