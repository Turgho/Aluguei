package repositories

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type AddressFilters struct {
	City         *string
	State        *string
	Country      *string
	Neighborhood *string
	Page         int
	PageSize     int
}

// AddressRepository define o contrato de acesso a dados de [entities.Address].
type AddressRepository interface {
	// Escrita
	Create(address *entities.Address) error
	Update(address *entities.Address) error
	Delete(id uuid.UUID) error

	// Leitura
	GetByID(id uuid.UUID) (*entities.Address, error)
	GetByZipCode(zipCode string) (*entities.Address, error)

	// Busca
	Search(filters AddressFilters) ([]*entities.Address, int64, error)
}
