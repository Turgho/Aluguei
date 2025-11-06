package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PropertyType string

const (
	PropertyTypeApartment  PropertyType = "apartment"
	PropertyTypeHouse      PropertyType = "house"
	PropertyTypeCommercial PropertyType = "commercial"
	PropertyTypeOther      PropertyType = "other"
)

type PropertyStatus string

const (
	PropertyStatusAvailable   PropertyStatus = "available"
	PropertyStatusRented      PropertyStatus = "rented"
	PropertyStatusMaintenance PropertyStatus = "maintenance"
)

type Property struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	OwnerID     uuid.UUID `json:"owner_id" gorm:"type:uuid;not null;index"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`

	// Endereço
	Address      string `json:"address" gorm:"type:varchar(255);not null"`
	Number       string `json:"number" gorm:"type:varchar(10);not null"`
	Complement   string `json:"complement" gorm:"type:varchar(100)"`
	Neighborhood string `json:"neighborhood" gorm:"type:varchar(100);not null"`
	City         string `json:"city" gorm:"type:varchar(100);not null"`
	State        string `json:"state" gorm:"type:varchar(2);not null"`
	ZipCode      string `json:"zip_code" gorm:"type:varchar(9);not null"`

	// Características
	Type       PropertyType `json:"type" gorm:"type:varchar(20);not null"`
	Bedrooms   int          `json:"bedrooms" gorm:"not null;default:0"`
	Bathrooms  int          `json:"bathrooms" gorm:"not null;default:1"`
	Area       float64      `json:"area" gorm:"type:decimal(10,2);not null"` // m²
	RentAmount float64      `json:"rent_amount" gorm:"type:decimal(10,2);not null"`

	// Status
	Status PropertyStatus `json:"status" gorm:"type:varchar(20);not null;default:available;index"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relacionamentos
	Owner     *Owner      `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Contracts []*Contract `json:"contracts,omitempty" gorm:"foreignKey:PropertyID"`
}

func (p *Property) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
