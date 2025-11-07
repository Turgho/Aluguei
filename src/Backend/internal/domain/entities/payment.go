package entities

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusPaid      PaymentStatus = "paid"
	PaymentStatusOverdue   PaymentStatus = "overdue"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

type Payment struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	ContractID uuid.UUID      `json:"contract_id" gorm:"type:uuid;not null;index"`
	Amount     float64        `json:"amount" gorm:"not null"`
	DueDate    time.Time      `json:"due_date" gorm:"type:date;not null"`
	PaidAmount *float64       `json:"paid_amount,omitempty"`
	PaidDate   *time.Time     `json:"paid_date,omitempty"`
	Status     PaymentStatus  `json:"status" gorm:"default:pending"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

// NewPayment creates a new payment
func NewPayment(contractID uuid.UUID, amount float64, dueDate time.Time, status PaymentStatus) *Payment {
	return &Payment{
		ID:         uuid.New(),
		ContractID: contractID,
		Amount:     amount,
		DueDate:    dueDate,
		Status:     status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// Update updates payment information
func (p *Payment) Update(amount float64, dueDate time.Time, paidAmount *float64, paidDate *time.Time, status PaymentStatus) {
	p.Amount = amount
	p.DueDate = dueDate
	p.PaidAmount = paidAmount
	p.PaidDate = paidDate
	p.Status = status
	p.UpdatedAt = time.Now()
}

// MarkAsPaid marks the payment as paid
func (p *Payment) MarkAsPaid(paidAmount float64, paidDate time.Time) {
	p.PaidAmount = &paidAmount
	p.PaidDate = &paidDate
	p.Status = PaymentStatusPaid
	p.UpdatedAt = time.Now()
}

// IsOverdue checks if payment is overdue
func (p *Payment) IsOverdue() bool {
	return p.Status == PaymentStatusPending && time.Now().After(p.DueDate)
}