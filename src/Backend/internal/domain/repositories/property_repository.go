package repositories

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type PropertyRepository interface {
	Create(ctx context.Context, property *entities.Property) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Property, error)
	GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Property, error)
	GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Property, int64, error)
	Update(ctx context.Context, property *entities.Property) error
	Delete(ctx context.Context, id uuid.UUID) error
}