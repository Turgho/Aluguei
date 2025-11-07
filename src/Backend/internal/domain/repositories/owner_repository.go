package repositories

import (
	"context"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type OwnerRepository interface {
	Create(ctx context.Context, owner *entities.Owner) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Owner, error)
	GetByEmail(ctx context.Context, email string) (*entities.Owner, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Owner, int64, error)
	Update(ctx context.Context, owner *entities.Owner) error
	Delete(ctx context.Context, id uuid.UUID) error
}