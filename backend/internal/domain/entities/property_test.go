// internal/domain/entities/property_test.go
package entities_test

import (
	"strings"
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type propertyTestCase struct {
	name         string
	ownerID      uuid.UUID
	addressID    uuid.UUID
	title        string
	description  string
	propertyType entities.PropertyType
	priceCents   int64
	condoCents   int64
	iptuCents    int64
	bedrooms     int
	bathrooms    int
	suites       int
	parking      int
	areaM2       int
	furnished    bool
	petFriendly  bool
	hasBalcony   bool
	hasElevator  bool

	wantErr bool
}

func TestNewProperty(t *testing.T) {
	ownerID := uuid.New()
	addressID := uuid.New()

	tests := []propertyTestCase{
		{
			name:         "imóvel válido",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Apartamento moderno com ótima localização, varanda e garagem.",
			propertyType: entities.PropertyApartment,
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       75,
			furnished:    true,
			petFriendly:  true,
			hasBalcony:   true,
			hasElevator:  true,
			wantErr:      false,
		},
		{
			name:         "imóvel válido sem condomínio e IPTU",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Casa térrea",
			description:  "Casa térrea ampla com quintal e garagem coberta.",
			propertyType: entities.PropertyHouse,
			priceCents:   180000,
			condoCents:   0,
			iptuCents:    0,
			bedrooms:     3,
			bathrooms:    2,
			suites:       1,
			parking:      2,
			areaM2:       120,
			furnished:    false,
			petFriendly:  true,
			hasBalcony:   false,
			hasElevator:  false,
			wantErr:      false,
		},
		{
			name:         "ownerID vazio",
			ownerID:      uuid.Nil,
			addressID:    addressID,
			title:        "Casa térrea",
			description:  "Casa térrea ampla com quintal e garagem coberta.",
			propertyType: entities.PropertyHouse,
			priceCents:   180000,
			condoCents:   0,
			iptuCents:    0,
			bedrooms:     3,
			bathrooms:    2,
			suites:       1,
			parking:      2,
			areaM2:       120,
			wantErr:      true,
		},
		{
			name:         "addressID vazio",
			ownerID:      ownerID,
			addressID:    uuid.Nil,
			title:        "Casa térrea",
			description:  "Casa térrea ampla com quintal e garagem coberta.",
			propertyType: entities.PropertyHouse,
			priceCents:   180000,
			condoCents:   0,
			iptuCents:    0,
			bedrooms:     3,
			bathrooms:    2,
			suites:       1,
			parking:      2,
			areaM2:       120,
			wantErr:      true,
		},
		{
			name:         "título inválido",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "AP",
			description:  "Apartamento moderno com ótima localização.",
			propertyType: entities.PropertyApartment,
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       75,
			wantErr:      true,
		},
		{
			name:         "descrição inválida",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Curta",
			propertyType: entities.PropertyApartment,
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       75,
			wantErr:      true,
		},
		{
			name:         "tipo inválido",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Apartamento moderno com ótima localização.",
			propertyType: "invalid",
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       75,
			wantErr:      true,
		},
		{
			name:         "preço inválido",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Apartamento moderno com ótima localização.",
			propertyType: entities.PropertyApartment,
			priceCents:   0,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       75,
			wantErr:      true,
		},
		{
			name:         "área inválida",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Apartamento moderno com ótima localização.",
			propertyType: entities.PropertyApartment,
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     2,
			bathrooms:    2,
			suites:       1,
			parking:      1,
			areaM2:       0,
			wantErr:      true,
		},
		{
			name:         "suítes maiores que quartos",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Apartamento moderno",
			description:  "Apartamento moderno com ótima localização.",
			propertyType: entities.PropertyApartment,
			priceCents:   250000,
			condoCents:   35000,
			iptuCents:    12000,
			bedrooms:     1,
			bathrooms:    2,
			suites:       2,
			parking:      1,
			areaM2:       75,
			wantErr:      true,
		},
		{
			name:         "terreno não pode ter quartos",
			ownerID:      ownerID,
			addressID:    addressID,
			title:        "Terreno amplo",
			description:  "Terreno amplo em ótima localização.",
			propertyType: entities.PropertyLand,
			priceCents:   100000,
			condoCents:   0,
			iptuCents:    0,
			bedrooms:     1,
			bathrooms:    0,
			suites:       0,
			parking:      0,
			areaM2:       300,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property, err := entities.NewProperty(
				tt.ownerID,
				tt.addressID,
				tt.title,
				tt.description,
				tt.propertyType,
				tt.priceCents,
				tt.condoCents,
				tt.iptuCents,
				tt.bedrooms,
				tt.bathrooms,
				tt.suites,
				tt.parking,
				tt.areaM2,
				tt.furnished,
				tt.petFriendly,
				tt.hasBalcony,
				tt.hasElevator,
			)

			if tt.wantErr {
				assertError(t, err)
				return
			}

			assertNoError(t, err)
			assertProperty(t, property, tt)
		})
	}
}

// assertProperty valida os dados principais do imóvel criado.
func assertProperty(t *testing.T, property *entities.Property, tt propertyTestCase) {
	t.Helper()

	if property == nil {
		t.Fatal("property não deveria ser nil")
	}

	if property.OwnerID != tt.ownerID {
		t.Errorf("ownerID inválido: got %s, want %s", property.OwnerID, tt.ownerID)
	}

	if property.AddressID != tt.addressID {
		t.Errorf("addressID inválido: got %s, want %s", property.AddressID, tt.addressID)
	}

	if property.Title != strings.TrimSpace(tt.title) {
		t.Errorf(
			"title inválido: got %s, want %s",
			property.Title,
			strings.TrimSpace(tt.title),
		)
	}

	if property.Description != strings.TrimSpace(tt.description) {
		t.Errorf(
			"description inválida: got %s, want %s",
			property.Description,
			strings.TrimSpace(tt.description),
		)
	}

	if property.Type != tt.propertyType {
		t.Errorf("type inválido: got %s, want %s", property.Type, tt.propertyType)
	}

	if property.Status != entities.PropertyAvailable {
		t.Errorf(
			"status inválido: got %s, want %s",
			property.Status,
			entities.PropertyAvailable,
		)
	}

	if property.PriceCents != tt.priceCents {
		t.Errorf(
			"priceCents inválido: got %d, want %d",
			property.PriceCents,
			tt.priceCents,
		)
	}

	if property.CondominiumFeeCents != tt.condoCents {
		t.Errorf(
			"condominiumFeeCents inválido: got %d, want %d",
			property.CondominiumFeeCents,
			tt.condoCents,
		)
	}

	if property.IPTUCents != tt.iptuCents {
		t.Errorf(
			"IPTUCents inválido: got %d, want %d",
			property.IPTUCents,
			tt.iptuCents,
		)
	}

	if property.Bedrooms != tt.bedrooms {
		t.Errorf(
			"bedrooms inválido: got %d, want %d",
			property.Bedrooms,
			tt.bedrooms,
		)
	}

	if property.Bathrooms != tt.bathrooms {
		t.Errorf(
			"bathrooms inválido: got %d, want %d",
			property.Bathrooms,
			tt.bathrooms,
		)
	}

	if property.Suites != tt.suites {
		t.Errorf(
			"suites inválido: got %d, want %d",
			property.Suites,
			tt.suites,
		)
	}

	if property.ParkingSpaces != tt.parking {
		t.Errorf(
			"parkingSpaces inválido: got %d, want %d",
			property.ParkingSpaces,
			tt.parking,
		)
	}

	if property.AreaM2 != tt.areaM2 {
		t.Errorf(
			"areaM2 inválida: got %d, want %d",
			property.AreaM2,
			tt.areaM2,
		)
	}

	if property.Furnished != tt.furnished {
		t.Errorf(
			"furnished inválido: got %v, want %v",
			property.Furnished,
			tt.furnished,
		)
	}

	if property.PetFriendly != tt.petFriendly {
		t.Errorf(
			"petFriendly inválido: got %v, want %v",
			property.PetFriendly,
			tt.petFriendly,
		)
	}

	if property.HasBalcony != tt.hasBalcony {
		t.Errorf(
			"hasBalcony inválido: got %v, want %v",
			property.HasBalcony,
			tt.hasBalcony,
		)
	}

	if property.HasElevator != tt.hasElevator {
		t.Errorf(
			"hasElevator inválido: got %v, want %v",
			property.HasElevator,
			tt.hasElevator,
		)
	}

	if property.CreatedAt.IsZero() {
		t.Error("CreatedAt não deveria ser zero")
	}

	if property.UpdatedAt.IsZero() {
		t.Error("UpdatedAt não deveria ser zero")
	}
}
