package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractStatus string

const (
	ContractStatusActive   ContractStatus = "active"
	ContractStatusExpired  ContractStatus = "expired"
	ContractStatusCanceled ContractStatus = "canceled"
)

type Contract struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	PropertyID uuid.UUID `json:"property_id" gorm:"type:uuid;not null"`
	TenantID   uuid.UUID `json:"tenant_id" gorm:"type:uuid;not null"`

	// Período
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`

	// Valores
	MonthlyRent   float64 `json:"monthly_rent" gorm:"not null"`
	DepositAmount float64 `json:"deposit_amount" gorm:"not null"`
	PaymentDueDay int     `json:"payment_due_day" gorm:"not null;default:5"`

	// Status
	Status ContractStatus `json:"status" gorm:"type:varchar(20);not null;default:active"`

	// Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relacionamentos
	Property *Property  `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Tenant   *User      `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Payments []*Payment `json:"payments,omitempty" gorm:"foreignKey:ContractID"`
}

func (c *Contract) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
