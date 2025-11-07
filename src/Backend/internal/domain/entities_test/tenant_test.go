package entities_test

import (
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTenant(t *testing.T) {
	ownerID := uuid.New()
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		&birthDate,
	)

	assert.NotEqual(t, uuid.Nil, tenant.ID)
	assert.Equal(t, "Jane Doe", tenant.Name)
	assert.Equal(t, "jane@example.com", tenant.Email)
	assert.Equal(t, "+5511888888888", tenant.Phone)
	assert.Equal(t, "98765432100", tenant.CPF)
	assert.Equal(t, ownerID, tenant.OwnerID)
	assert.Equal(t, &birthDate, tenant.BirthDate)
}

func TestTenant_UpdateProfile(t *testing.T) {
	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		uuid.New(),
		nil,
	)

	newBirthDate := time.Date(1985, 5, 15, 0, 0, 0, 0, time.UTC)
	tenant.UpdateProfile(
		"Jane Smith",
		"janesmith@example.com",
		"+5511777777777",
		"12345678900",
		&newBirthDate,
	)

	assert.Equal(t, "Jane Smith", tenant.Name)
	assert.Equal(t, "janesmith@example.com", tenant.Email)
	assert.Equal(t, "+5511777777777", tenant.Phone)
	assert.Equal(t, "12345678900", tenant.CPF)
	assert.Equal(t, &newBirthDate, tenant.BirthDate)
}