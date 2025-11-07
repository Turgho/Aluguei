package repositories

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entities.Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Tenant, error)
	GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Tenant, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Tenant, int64, error)
	Update(ctx context.Context, tenant *entities.Tenant) error
	Delete(ctx context.Context, id uuid.UUID) error
}