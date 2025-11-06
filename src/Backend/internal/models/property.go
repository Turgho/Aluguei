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
)

type PropertyStatus string

const (
	PropertyStatusAvailable   PropertyStatus = "available"
	PropertyStatusRented      PropertyStatus = "rented"
	PropertyStatusMaintenance PropertyStatus = "maintenance"
)

type Property struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	OwnerID     uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`

	// Endereço
	Address      string  `json:"address" gorm:"not null"`
	Number       string  `json:"number" gorm:"not null"`
	Complement   string  `json:"complement"`
	Neighborhood string  `json:"neighborhood" gorm:"not null"`
	City         string  `json:"city" gorm:"not null"`
	State        string  `json:"state" gorm:"not null;size:2"`
	ZipCode      string  `json:"zip_code" gorm:"not null;size:9"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`

	// Características
	Type       PropertyType `json:"type" gorm:"type:varchar(20);not null"`
	Bedrooms   int          `json:"bedrooms" gorm:"not null;default:0"`
	Bathrooms  int          `json:"bathrooms" gorm:"not null;default:1"`
	Area       float64      `json:"area" gorm:"not null"` // m²
	RentAmount float64      `json:"rent_amount" gorm:"not null"`
	CondoFee   float64      `json:"condo_fee" gorm:"default:0"`
	IPTU       float64      `json:"iptu" gorm:"default:0"`

	// Status
	Status PropertyStatus `json:"status" gorm:"type:varchar(20);not null;default:available"`

	// Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relacionamentos
	Owner     *User       `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Contracts []*Contract `json:"contracts,omitempty" gorm:"foreignKey:PropertyID"`
}

func (p *Property) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
