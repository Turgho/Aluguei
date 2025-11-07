package persistence

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) repositories.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(ctx context.Context, payment *entities.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Payment, error) {
	var payment entities.Payment
	err := r.db.WithContext(ctx).First(&payment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetByContractID(ctx context.Context, contractID uuid.UUID) ([]*entities.Payment, error) {
	var payments []*entities.Payment
	err := r.db.WithContext(ctx).Where("contract_id = ?", contractID).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetOverdue(ctx context.Context) ([]*entities.Payment, error) {
	var payments []*entities.Payment
	err := r.db.WithContext(ctx).Where("status = ?", entities.PaymentStatusOverdue).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetByPeriod(ctx context.Context, startDate, endDate time.Time) ([]*entities.Payment, error) {
	var payments []*entities.Payment
	err := r.db.WithContext(ctx).Where("due_date BETWEEN ? AND ?", startDate, endDate).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Payment, int64, error) {
	var payments []*entities.Payment
	var total int64

	r.db.WithContext(ctx).Model(&entities.Payment{}).Count(&total)

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&payments).Error
	return payments, total, err
}

func (r *paymentRepository) Update(ctx context.Context, payment *entities.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

func (r *paymentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Payment{}, "id = ?", id).Error
}