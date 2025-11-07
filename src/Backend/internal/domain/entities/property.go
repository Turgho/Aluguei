package entities

import (
	"time"

	"github.com/google/uuid"
)

type PropertyStatus string

const (
	PropertyStatusAvailable   PropertyStatus = "available"
	PropertyStatusRented      PropertyStatus = "rented"
	PropertyStatusMaintenance PropertyStatus = "maintenance"
	PropertyStatusInactive    PropertyStatus = "inactive"
)

type Property struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	OwnerID     uuid.UUID      `json:"owner_id" gorm:"type:uuid;not null;index"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Address     string         `json:"address" gorm:"not null"`
	City        string         `json:"city" gorm:"not null"`
	State       string         `json:"state" gorm:"not null"`
	ZipCode     string         `json:"zip_code"`
	Bedrooms    int            `json:"bedrooms"`
	Bathrooms   int            `json:"bathrooms"`
	Area        int            `json:"area"`
	RentAmount  float64        `json:"rent_amount" gorm:"not null"`
	Status      PropertyStatus `json:"status" gorm:"default:available"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// NewProperty creates a new property
func NewProperty(ownerID uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64) *Property {
	return &Property{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Title:       title,
		Description: description,
		Address:     address,
		City:        city,
		State:       state,
		ZipCode:     zipCode,
		Bedrooms:    bedrooms,
		Bathrooms:   bathrooms,
		Area:        area,
		RentAmount:  rentAmount,
		Status:      PropertyStatusAvailable,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// UpdateStatus updates the property status
func (p *Property) UpdateStatus(status PropertyStatus) {
	p.Status = status
	p.UpdatedAt = time.Now()
}

// Update updates property information
func (p *Property) Update(title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64, status PropertyStatus) {
	p.Title = title
	p.Description = description
	p.Address = address
	p.City = city
	p.State = state
	p.ZipCode = zipCode
	p.Bedrooms = bedrooms
	p.Bathrooms = bathrooms
	p.Area = area
	p.RentAmount = rentAmount
	p.Status = status
	p.UpdatedAt = time.Now()
}

// IsAvailable checks if property is available for rent
func (p *Property) IsAvailable() bool {
	return p.Status == PropertyStatusAvailable
}