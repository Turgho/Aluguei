package usecases

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
)

type PaymentUseCase struct {
	paymentRepo repositories.PaymentRepository
}

func NewPaymentUseCase(paymentRepo repositories.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
	}
}

func (uc *PaymentUseCase) CreatePayment(ctx context.Context, contractID uuid.UUID, amount float64, dueDate time.Time, status entities.PaymentStatus) (*entities.Payment, error) {
	payment := entities.NewPayment(contractID, amount, dueDate, status)
	
	if err := uc.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}
	
	return payment, nil
}

func (uc *PaymentUseCase) GetPayment(ctx context.Context, id uuid.UUID) (*entities.Payment, error) {
	return uc.paymentRepo.GetByID(ctx, id)
}

func (uc *PaymentUseCase) GetPaymentsByContract(ctx context.Context, contractID uuid.UUID) ([]*entities.Payment, error) {
	return uc.paymentRepo.GetByContractID(ctx, contractID)
}

func (uc *PaymentUseCase) GetOverduePayments(ctx context.Context) ([]*entities.Payment, error) {
	return uc.paymentRepo.GetOverdue(ctx)
}

func (uc *PaymentUseCase) GetPaymentsByPeriod(ctx context.Context, startDate, endDate time.Time) ([]*entities.Payment, error) {
	return uc.paymentRepo.GetByPeriod(ctx, startDate, endDate)
}

func (uc *PaymentUseCase) GetAllPayments(ctx context.Context, page, limit int) ([]*entities.Payment, int64, error) {
	return uc.paymentRepo.GetAll(ctx, page, limit)
}

func (uc *PaymentUseCase) UpdatePayment(ctx context.Context, id uuid.UUID, amount float64, dueDate time.Time, paidAmount *float64, paidDate *time.Time, status entities.PaymentStatus) error {
	payment, err := uc.paymentRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	payment.Update(amount, dueDate, paidAmount, paidDate, status)
	
	return uc.paymentRepo.Update(ctx, payment)
}

func (uc *PaymentUseCase) DeletePayment(ctx context.Context, id uuid.UUID) error {
	return uc.paymentRepo.Delete(ctx, id)
}

func (uc *PaymentUseCase) GetPaymentsByProperty(ctx context.Context, propertyID uuid.UUID, page, limit int) ([]*entities.Payment, int64, error) {
	// Esta função precisaria ser implementada no repositório
	// Por enquanto, retorna lista vazia
	return []*entities.Payment{}, 0, nil
}