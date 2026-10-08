// internal/domain/repositories/property_repository.go
package repositories

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type PropertyFilters struct {
	Type          *entities.PropertyType
	Status        *entities.PropertyStatus
	MinPriceCents *int64
	MaxPriceCents *int64
	MinAreaM2     *int
	Bedrooms      *int
	Furnished     *bool
	PetFriendly   *bool
	City          *string
	State         *string
	Page          int
	PageSize      int
}

// PropertyRepository define o contrato de acesso a dados de [entities.Property].
type PropertyRepository interface {
	// Escrita
	Create(property *entities.Property) error
	Update(property *entities.Property) error
	Delete(id uuid.UUID) error

	// Leitura simples
	GetByID(id uuid.UUID) (*entities.Property, error)
	GetByOwnerID(ownerID uuid.UUID) ([]*entities.Property, error)

	// Busca tipada com paginação
	Search(filters PropertyFilters) ([]*entities.Property, int64, error)
	//                               ↑ resultado       ↑ total para paginação
}
