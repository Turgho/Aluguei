package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Phone     string     `json:"phone" gorm:"not null"`
	CPF       string     `json:"cpf" gorm:"unique;not null"`
	BirthDate *time.Time `json:"birth_date,omitempty" gorm:"type:date"`
	OwnerID   uuid.UUID  `json:"owner_id" gorm:"type:uuid;not null;index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewTenant creates a new tenant
func NewTenant(name, email, phone, cpf string, ownerID uuid.UUID, birthDate *time.Time) *Tenant {
	return &Tenant{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		CPF:       cpf,
		OwnerID:   ownerID,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateProfile updates tenant profile
func (t *Tenant) UpdateProfile(name, email, phone, cpf string, birthDate *time.Time) {
	t.Name = name
	t.Email = email
	t.Phone = phone
	t.CPF = cpf
	t.BirthDate = birthDate
	t.UpdatedAt = time.Now()
}