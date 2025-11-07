package testhelpers

import (
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

// CreateTestOwner creates a test owner
func CreateTestOwner() *entities.Owner {
	return entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)
}

// CreateTestTenant creates a test tenant
func CreateTestTenant(ownerID uuid.UUID) *entities.Tenant {
	return entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		nil,
	)
}

// CreateTestProperty creates a test property
func CreateTestProperty(ownerID uuid.UUID) *entities.Property {
	return entities.NewProperty(
		ownerID,
		"Test Property",
		"A nice test property",
		"123 Main St",
		"São Paulo",
		"SP",
		"01234567",
		2,
		1,
		50,
		1500.0,
	)
}

// CreateTestContract creates a test contract
func CreateTestContract(propertyID, tenantID uuid.UUID) *entities.Contract {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	
	return entities.NewContract(
		propertyID,
		tenantID,
		startDate,
		endDate,
		1500.0,
		5,
		entities.ContractStatusActive,
	)
}

// CreateTestPayment creates a test payment
func CreateTestPayment(contractID uuid.UUID) *entities.Payment {
	dueDate := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	
	return entities.NewPayment(
		contractID,
		1500.0,
		dueDate,
		entities.PaymentStatusPending,
	)
}