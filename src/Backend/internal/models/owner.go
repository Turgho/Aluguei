package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Owner representa o proprietário
type Owner struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     string     `gorm:"type:varchar(20)" json:"phone"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	CPF       string     `gorm:"type:varchar(14);unique" json:"cpf"`
	BirthDate *time.Time `gorm:"type:date" json:"birth_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Relations
	Properties []Property `gorm:"foreignKey:OwnerID" json:"properties,omitempty"`
	Tenants    []Tenant   `gorm:"foreignKey:OwnerID" json:"tenants,omitempty"`
}

func (o *Owner) BeforeCreate(tx *gorm.DB) error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}
