package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusOverdue PaymentStatus = "overdue"
)

type PaymentMethod string

const (
	PaymentMethodPIX          PaymentMethod = "pix"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodOther        PaymentMethod = "other"
)

type Payment struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	ContractID uuid.UUID `json:"contract_id" gorm:"type:uuid;not null;index"`

	// Datas
	DueDate        time.Time  `json:"due_date" gorm:"type:date;not null;index"`
	PaidDate       *time.Time `json:"paid_date" gorm:"type:date"`
	ReferenceMonth time.Time  `json:"reference_month" gorm:"type:date;not null"`

	// Valores
	Amount     float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	PaidAmount float64 `json:"paid_amount" gorm:"type:decimal(10,2);default:0"`
	LateFee    float64 `json:"late_fee" gorm:"type:decimal(10,2);default:0"`

	// Método e status
	Method PaymentMethod `json:"method" gorm:"type:varchar(20)"`
	Status PaymentStatus `json:"status" gorm:"type:varchar(20);not null;default:pending;index"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relacionamentos
	Contract *Contract `json:"contract,omitempty" gorm:"foreignKey:ContractID"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
