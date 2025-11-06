package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusLate     PaymentStatus = "late"
	PaymentStatusCanceled PaymentStatus = "canceled"
)

type PaymentMethod string

const (
	PaymentMethodPIX    PaymentMethod = "pix"
	PaymentMethodBoleto PaymentMethod = "boleto"
	PaymentMethodCredit PaymentMethod = "credit_card"
	PaymentMethodDebit  PaymentMethod = "debit_card"
)

type Payment struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	ContractID uuid.UUID `json:"contract_id" gorm:"type:uuid;not null"`

	// Datas
	DueDate  time.Time  `json:"due_date" gorm:"not null"`
	PaidDate *time.Time `json:"paid_date"`

	// Valores
	Amount     float64 `json:"amount" gorm:"not null"`
	PaidAmount float64 `json:"paid_amount" gorm:"default:0"`
	LateFee    float64 `json:"late_fee" gorm:"default:0"`

	// Método e status
	Method PaymentMethod `json:"method" gorm:"type:varchar(20)"`
	Status PaymentStatus `json:"status" gorm:"type:varchar(20);not null;default:pending"`

	// Referências externas
	ReferenceID string `json:"reference_id"` // ID do gateway de pagamento

	// Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relacionamentos
	Contract *Contract `json:"contract,omitempty" gorm:"foreignKey:ContractID"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
