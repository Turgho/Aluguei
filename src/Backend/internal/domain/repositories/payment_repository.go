package repositories

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entities.Payment) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Payment, error)
	GetByContractID(ctx context.Context, contractID uuid.UUID) ([]*entities.Payment, error)
	GetOverdue(ctx context.Context) ([]*entities.Payment, error)
	GetByPeriod(ctx context.Context, startDate, endDate time.Time) ([]*entities.Payment, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Payment, int64, error)
	Update(ctx context.Context, payment *entities.Payment) error
	Delete(ctx context.Context, id uuid.UUID) error
}