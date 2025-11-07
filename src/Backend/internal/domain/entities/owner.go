package entities

import (
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"-" gorm:"not null"`
	Phone     string     `json:"phone" gorm:"not null"`
	CPF       string     `json:"cpf" gorm:"unique;not null"`
	BirthDate *time.Time `json:"birth_date,omitempty" gorm:"type:date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewOwner creates a new owner
func NewOwner(name, email, password, phone, cpf string, birthDate *time.Time) *Owner {
	return &Owner{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		Phone:     phone,
		CPF:       cpf,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateProfile updates owner profile
func (o *Owner) UpdateProfile(name, email, phone, cpf string, birthDate *time.Time) {
	o.Name = name
	o.Email = email
	o.Phone = phone
	o.CPF = cpf
	o.BirthDate = birthDate
	o.UpdatedAt = time.Now()
}