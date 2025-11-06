package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tenant representa um inquilino (apenas dados de contato)
type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;index" json:"owner_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	CPF       string    `gorm:"type:varchar(14)" json:"cpf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Owner     Owner      `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Contracts []Contract `gorm:"foreignKey:TenantID" json:"contracts,omitempty"`
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
