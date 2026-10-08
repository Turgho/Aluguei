package entities

import (
	"errors"
	"strings"
	"time"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`

	Street       string `gorm:"type:varchar(120);not null"`
	Number       string `gorm:"type:varchar(20);not null"`
	Complement   string `gorm:"type:varchar(100)"`
	Neighborhood string `gorm:"type:varchar(100);not null"`
	City         string `gorm:"type:varchar(100);not null;index"`
	State        string `gorm:"type:varchar(2);not null;index"`
	Country      string `gorm:"type:char(2);not null;default:'BR';index"`
	ZipCode      string `gorm:"type:varchar(10);not null"`

	Latitude  float64 `gorm:"type:decimal(10,8)"`
	Longitude float64 `gorm:"type:decimal(11,8)"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// NewAddress cria e valida uma nova instância de Address.
func NewAddress(
	street,
	number,
	complement,
	neighborhood,
	city,
	state,
	country,
	zipCode string,
	latitude,
	longitude float64,
) (*Address, error) {
	var errs []string

	// Normalização
	street = strings.TrimSpace(street)
	number = strings.TrimSpace(number)
	complement = strings.TrimSpace(complement)
	neighborhood = strings.TrimSpace(neighborhood)
	city = strings.TrimSpace(city)
	state = strings.ToUpper(strings.TrimSpace(state))
	country = strings.ToUpper(strings.TrimSpace(country))
	zipCode = addressValidators.NormalizeZipCode(country, zipCode)

	// ————— Rua —————
	if !addressValidators.ValidateStreet(street) {
		errs = append(errs, "rua inválida")
	}

	// ————— Número —————
	if !addressValidators.ValidateNumber(number) {
		errs = append(errs, "número inválido")
	}

	// ————— Complemento —————
	if !addressValidators.ValidateComplement(complement) {
		errs = append(errs, "complemento inválido")
	}

	// ————— Bairro —————
	if !addressValidators.ValidateNeighborhood(neighborhood) {
		errs = append(errs, "bairro inválido")
	}

	// ————— Cidade —————
	if !addressValidators.ValidateCity(city) {
		errs = append(errs, "cidade inválida")
	}

	// ————— Estado —————
	if !addressValidators.ValidateState(state) {
		errs = append(errs, "estado inválido")
	}

	if !addressValidators.ValidateCountry(country) {
		errs = append(errs, "country inválido")
	}

	// ————— CEP —————
	if !addressValidators.ValidateZipCode(country, zipCode) {
		errs = append(errs, "zipcode inválido")
	}

	// ————— Latitude —————
	if !addressValidators.ValidateLatitude(latitude) {
		errs = append(errs, "latitude inválida")
	}

	// ————— Longitude —————
	if !addressValidators.ValidateLongitude(longitude) {
		errs = append(errs, "longitude inválida")
	}

	// ————— Erros —————
	if len(errs) > 0 {
		return nil, errors.New(strings.Join(errs, "; "))
	}

	now := time.Now().UTC()

	return &Address{
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		Country:      country,
		ZipCode:      zipCode,
		Latitude:     latitude,
		Longitude:    longitude,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
