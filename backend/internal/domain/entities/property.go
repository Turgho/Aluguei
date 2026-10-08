// Package entities define as entidades de domínio da aplicação.
package entities

import (
	"errors"
	"strings"
	"time"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
	sharedValidators "github.com/Turgho/Aluguei/pkg/validators/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyType representa o tipo do imóvel.
type PropertyType string

const (
	PropertyHouse      PropertyType = "house"
	PropertyApartment  PropertyType = "apartment"
	PropertyStudio     PropertyType = "studio"
	PropertyLoft       PropertyType = "loft"
	PropertyBusiness   PropertyType = "business"
	PropertyCommercial PropertyType = "commercial"
	PropertyOffice     PropertyType = "office"
	PropertyStore      PropertyType = "store"
	PropertyLand       PropertyType = "land"
)

// PropertyStatus representa o estado do anúncio.
type PropertyStatus string

const (
	PropertyAvailable PropertyStatus = "available"
	PropertyRented    PropertyStatus = "rented"
	PropertyInactive  PropertyStatus = "inactive"
	PropertySold      PropertyStatus = "sold"
)

// Property representa um imóvel no sistema.
type Property struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`

	OwnerID uuid.UUID `gorm:"type:uuid;not null;index"`

	AddressID uuid.UUID `gorm:"type:uuid;not null"`
	Address   Address   `gorm:"foreignKey:AddressID"`

	Title       string         `gorm:"type:varchar(120);not null"`
	Description string         `gorm:"type:text"`
	Type        PropertyType   `gorm:"type:varchar(50);not null;index"`
	Status      PropertyStatus `gorm:"type:varchar(50);not null;default:available;index"`

	// Valores em centavos
	PriceCents          int64 `gorm:"not null"`
	CondominiumFeeCents int64 `gorm:"not null;default:0"`
	IPTUCents           int64 `gorm:"not null;default:0"`

	Bedrooms      int `gorm:"not null;default:0"`
	Bathrooms     int `gorm:"not null;default:0"`
	Suites        int `gorm:"not null;default:0"`
	ParkingSpaces int `gorm:"not null;default:0"`
	AreaM2        int `gorm:"not null;default:0"`

	Furnished   bool `gorm:"not null;default:false"`
	PetFriendly bool `gorm:"not null;default:false"`
	HasBalcony  bool `gorm:"not null;default:false"`
	HasElevator bool `gorm:"not null;default:false"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// NewProperty cria e valida uma nova instância de Property.
func NewProperty(
	ownerID uuid.UUID,
	addressID uuid.UUID,
	title, description string,
	propertyType PropertyType,
	priceCents, condominiumFeeCents, iptuCents int64,
	bedrooms, bathrooms, suites, parkingSpaces, areaM2 int,
	furnished, petFriendly, hasBalcony, hasElevator bool,
) (*Property, error) {
	var errs []string

	// Normalização
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	// ————— IDs —————
	if !sharedValidators.ValidateUUID(ownerID) {
		errs = append(errs, "ownerID é obrigatório")
	}

	if !sharedValidators.ValidateUUID(addressID) {
		errs = append(errs, "addressID é obrigatório")
	}

	// ————— Texto —————
	if !propertyValidators.ValidateTitle(title) {
		errs = append(errs, "título inválido")
	}

	if !propertyValidators.ValidateDescription(description) {
		errs = append(errs, "descrição inválida")
	}

	// ————— Tipo e status —————
	if !propertyValidators.ValidateType(string(propertyType)) {
		errs = append(errs, "tipo do imóvel inválido")
	}

	// ————— Valores monetários —————
	if !propertyValidators.ValidatePriceCents(priceCents) {
		errs = append(errs, "preço deve ser maior que zero")
	}

	if !propertyValidators.ValidateOptionalFeeCents(condominiumFeeCents) {
		errs = append(errs, "valor do condomínio inválido")
	}

	if !propertyValidators.ValidateOptionalFeeCents(iptuCents) {
		errs = append(errs, "valor do IPTU inválido")
	}

	// ————— Números —————
	if !propertyValidators.ValidateRooms(bedrooms) {
		errs = append(errs, "número de quartos inválido")
	}

	if !propertyValidators.ValidateRooms(bathrooms) {
		errs = append(errs, "número de banheiros inválido")
	}

	if !propertyValidators.ValidateRooms(suites) {
		errs = append(errs, "número de suítes inválido")
	}

	if !propertyValidators.ValidateRooms(parkingSpaces) {
		errs = append(errs, "número de vagas inválido")
	}

	if !propertyValidators.ValidateAreaM2(areaM2) {
		errs = append(errs, "área do imóvel inválida")
	}

	// ————— Regras de domínio —————
	if suites > bedrooms {
		errs = append(errs, "suítes não podem ser maiores que quartos")
	}

	if propertyType == PropertyLand {
		if bedrooms > 0 ||
			bathrooms > 0 ||
			suites > 0 {
			errs = append(errs, "terrenos não podem possuir cômodos")
		}
	}

	// ————— Erros —————
	if len(errs) > 0 {
		return nil, errors.New(strings.Join(errs, "; "))
	}

	now := time.Now().UTC()

	return &Property{
		OwnerID:             ownerID,
		AddressID:           addressID,
		Title:               title,
		Description:         description,
		Type:                propertyType,
		Status:              PropertyAvailable,
		PriceCents:          priceCents,
		CondominiumFeeCents: condominiumFeeCents,
		IPTUCents:           iptuCents,
		Bedrooms:            bedrooms,
		Bathrooms:           bathrooms,
		Suites:              suites,
		ParkingSpaces:       parkingSpaces,
		AreaM2:              areaM2,
		Furnished:           furnished,
		PetFriendly:         petFriendly,
		HasBalcony:          hasBalcony,
		HasElevator:         hasElevator,
		CreatedAt:           now,
		UpdatedAt:           now,
	}, nil
}
