package entities_test

import (
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewContract(t *testing.T) {
	propertyID := uuid.New()
	tenantID := uuid.New()
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	contract := entities.NewContract(
		propertyID,
		tenantID,
		startDate,
		endDate,
		1500.0,
		5,
		entities.ContractStatusActive,
	)

	assert.NotEqual(t, uuid.Nil, contract.ID)
	assert.Equal(t, propertyID, contract.PropertyID)
	assert.Equal(t, tenantID, contract.TenantID)
	assert.Equal(t, startDate, contract.StartDate)
	assert.Equal(t, endDate, contract.EndDate)
	assert.Equal(t, 1500.0, contract.MonthlyRent)
	assert.Equal(t, 5, contract.PaymentDueDay)
	assert.Equal(t, entities.ContractStatusActive, contract.Status)
}

func TestContract_Cancel(t *testing.T) {
	contract := entities.NewContract(
		uuid.New(),
		uuid.New(),
		time.Now(),
		time.Now().AddDate(1, 0, 0),
		1500.0,
		5,
		entities.ContractStatusActive,
	)

	contract.Cancel()
	assert.Equal(t, entities.ContractStatusCancelled, contract.Status)
}

func TestContract_IsActive(t *testing.T) {
	// Active contract
	activeContract := entities.NewContract(
		uuid.New(),
		uuid.New(),
		time.Now().AddDate(0, -1, 0), // Started 1 month ago
		time.Now().AddDate(1, 0, 0),  // Ends in 1 year
		1500.0,
		5,
		entities.ContractStatusActive,
	)

	assert.True(t, activeContract.IsActive())

	// Expired contract
	expiredContract := entities.NewContract(
		uuid.New(),
		uuid.New(),
		time.Now().AddDate(-1, 0, 0), // Started 1 year ago
		time.Now().AddDate(0, -1, 0), // Ended 1 month ago
		1500.0,
		5,
		entities.ContractStatusActive,
	)

	assert.False(t, expiredContract.IsActive())

	// Cancelled contract
	cancelledContract := entities.NewContract(
		uuid.New(),
		uuid.New(),
		time.Now().AddDate(0, -1, 0), // Started 1 month ago
		time.Now().AddDate(1, 0, 0),  // Ends in 1 year
		1500.0,
		5,
		entities.ContractStatusCancelled,
	)

	assert.False(t, cancelledContract.IsActive())
}