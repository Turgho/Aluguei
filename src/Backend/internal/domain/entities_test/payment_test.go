package entities_test

import (
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPayment(t *testing.T) {
	contractID := uuid.New()
	dueDate := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)

	payment := entities.NewPayment(
		contractID,
		1500.0,
		dueDate,
		entities.PaymentStatusPending,
	)

	assert.NotEqual(t, uuid.Nil, payment.ID)
	assert.Equal(t, contractID, payment.ContractID)
	assert.Equal(t, 1500.0, payment.Amount)
	assert.Equal(t, dueDate, payment.DueDate)
	assert.Equal(t, entities.PaymentStatusPending, payment.Status)
	assert.Nil(t, payment.PaidAmount)
	assert.Nil(t, payment.PaidDate)
}

func TestPayment_MarkAsPaid(t *testing.T) {
	payment := entities.NewPayment(
		uuid.New(),
		1500.0,
		time.Now(),
		entities.PaymentStatusPending,
	)

	paidDate := time.Now()
	payment.MarkAsPaid(1500.0, paidDate)

	assert.Equal(t, entities.PaymentStatusPaid, payment.Status)
	assert.NotNil(t, payment.PaidAmount)
	assert.Equal(t, 1500.0, *payment.PaidAmount)
	assert.NotNil(t, payment.PaidDate)
	assert.Equal(t, paidDate, *payment.PaidDate)
}

func TestPayment_IsOverdue(t *testing.T) {
	// Not overdue - pending but not past due date
	futurePayment := entities.NewPayment(
		uuid.New(),
		1500.0,
		time.Now().AddDate(0, 0, 1), // Tomorrow
		entities.PaymentStatusPending,
	)
	assert.False(t, futurePayment.IsOverdue())

	// Overdue - pending and past due date
	overduePayment := entities.NewPayment(
		uuid.New(),
		1500.0,
		time.Now().AddDate(0, 0, -1), // Yesterday
		entities.PaymentStatusPending,
	)
	assert.True(t, overduePayment.IsOverdue())

	// Not overdue - paid
	paidPayment := entities.NewPayment(
		uuid.New(),
		1500.0,
		time.Now().AddDate(0, 0, -1), // Yesterday
		entities.PaymentStatusPaid,
	)
	assert.False(t, paidPayment.IsOverdue())
}