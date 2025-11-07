package repositories

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type ContractRepository interface {
	Create(ctx context.Context, contract *entities.Contract) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Contract, error)
	GetByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*entities.Contract, error)
	GetByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*entities.Contract, error)
	GetActiveByPropertyID(ctx context.Context, propertyID uuid.UUID) (*entities.Contract, error)
	GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Contract, int64, error)
	Update(ctx context.Context, contract *entities.Contract) error
	Delete(ctx context.Context, id uuid.UUID) error
}