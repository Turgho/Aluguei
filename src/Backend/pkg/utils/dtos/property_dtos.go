package dtos

import "time"

type CreatePropertyRequest struct {
	Title        string  `json:"title" validate:"required,min=2,max=255"`
	Description  string  `json:"description"`
	Address      string  `json:"address" validate:"required,min=5,max=255"`
	Number       string  `json:"number" validate:"required,max=10"`
	Complement   string  `json:"complement" validate:"max=100"`
	Neighborhood string  `json:"neighborhood" validate:"required,max=100"`
	City         string  `json:"city" validate:"required,max=100"`
	State        string  `json:"state" validate:"required,len=2"`
	ZipCode      string  `json:"zip_code" validate:"required,len=9"`
	Type         string  `json:"type" validate:"required,oneof=apartment house commercial other"`
	Bedrooms     int     `json:"bedrooms" validate:"required,min=0"`
	Bathrooms    int     `json:"bathrooms" validate:"required,min=1"`
	Area         float64 `json:"area" validate:"required,min=0"`
	RentAmount   float64 `json:"rent_amount" validate:"required,min=0"`
}

type UpdatePropertyRequest struct {
	Title        string  `json:"title" validate:"omitempty,min=2,max=255"`
	Description  string  `json:"description"`
	Address      string  `json:"address" validate:"omitempty,min=5,max=255"`
	Number       string  `json:"number" validate:"omitempty,max=10"`
	Complement   string  `json:"complement" validate:"omitempty,max=100"`
	Neighborhood string  `json:"neighborhood" validate:"omitempty,max=100"`
	City         string  `json:"city" validate:"omitempty,max=100"`
	State        string  `json:"state" validate:"omitempty,len=2"`
	ZipCode      string  `json:"zip_code" validate:"omitempty,len=9"`
	Type         string  `json:"type" validate:"omitempty,oneof=apartment house commercial other"`
	Bedrooms     int     `json:"bedrooms" validate:"omitempty,min=0"`
	Bathrooms    int     `json:"bathrooms" validate:"omitempty,min=1"`
	Area         float64 `json:"area" validate:"omitempty,min=0"`
	RentAmount   float64 `json:"rent_amount" validate:"omitempty,min=0"`
}

type PropertyResponse struct {
	ID           string    `json:"id"`
	OwnerID      string    `json:"owner_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Address      string    `json:"address"`
	Number       string    `json:"number"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Type         string    `json:"type"`
	Bedrooms     int       `json:"bedrooms"`
	Bathrooms    int       `json:"bathrooms"`
	Area         float64   `json:"area"`
	RentAmount   float64   `json:"rent_amount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
