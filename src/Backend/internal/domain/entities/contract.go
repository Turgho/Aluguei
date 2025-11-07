package entities

import (
	"time"

	"github.com/google/uuid"
)

type ContractStatus string

const (
	ContractStatusActive    ContractStatus = "active"
	ContractStatusExpired   ContractStatus = "expired"
	ContractStatusCancelled ContractStatus = "cancelled"
	ContractStatusPending   ContractStatus = "pending"
)

type Contract struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	PropertyID   uuid.UUID      `json:"property_id" gorm:"type:uuid;not null;index"`
	TenantID     uuid.UUID      `json:"tenant_id" gorm:"type:uuid;not null;index"`
	StartDate    time.Time      `json:"start_date" gorm:"type:date;not null"`
	EndDate      time.Time      `json:"end_date" gorm:"type:date;not null"`
	MonthlyRent  float64        `json:"monthly_rent" gorm:"not null"`
	PaymentDueDay int           `json:"payment_due_day" gorm:"not null;check:payment_due_day >= 1 AND payment_due_day <= 28"`
	Status       ContractStatus `json:"status" gorm:"default:active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// NewContract creates a new contract
func NewContract(propertyID, tenantID uuid.UUID, startDate, endDate time.Time, monthlyRent float64, paymentDueDay int, status ContractStatus) *Contract {
	return &Contract{
		ID:            uuid.New(),
		PropertyID:    propertyID,
		TenantID:      tenantID,
		StartDate:     startDate,
		EndDate:       endDate,
		MonthlyRent:   monthlyRent,
		PaymentDueDay: paymentDueDay,
		Status:        status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

// Update updates contract information
func (c *Contract) Update(startDate, endDate time.Time, monthlyRent float64, paymentDueDay int, status ContractStatus) {
	c.StartDate = startDate
	c.EndDate = endDate
	c.MonthlyRent = monthlyRent
	c.PaymentDueDay = paymentDueDay
	c.Status = status
	c.UpdatedAt = time.Now()
}

// Cancel cancels the contract
func (c *Contract) Cancel() {
	c.Status = ContractStatusCancelled
	c.UpdatedAt = time.Now()
}

// IsActive checks if contract is active
func (c *Contract) IsActive() bool {
	return c.Status == ContractStatusActive && time.Now().Before(c.EndDate)
}

// IsExpired checks if contract is expired
func (c *Contract) IsExpired() bool {
	return time.Now().After(c.EndDate)
}