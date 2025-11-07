package entities_test

import (
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOwner(t *testing.T) {
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		&birthDate,
	)

	assert.NotEqual(t, uuid.Nil, owner.ID)
	assert.Equal(t, "John Doe", owner.Name)
	assert.Equal(t, "john@example.com", owner.Email)
	assert.Equal(t, "hashedpassword", owner.Password)
	assert.Equal(t, "+5511999999999", owner.Phone)
	assert.Equal(t, "12345678901", owner.CPF)
	assert.Equal(t, &birthDate, owner.BirthDate)
}

func TestOwner_UpdateProfile(t *testing.T) {
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)

	newBirthDate := time.Date(1985, 5, 15, 0, 0, 0, 0, time.UTC)
	owner.UpdateProfile(
		"John Smith",
		"johnsmith@example.com",
		"+5511888888888",
		"98765432100",
		&newBirthDate,
	)

	assert.Equal(t, "John Smith", owner.Name)
	assert.Equal(t, "johnsmith@example.com", owner.Email)
	assert.Equal(t, "+5511888888888", owner.Phone)
	assert.Equal(t, "98765432100", owner.CPF)
	assert.Equal(t, &newBirthDate, owner.BirthDate)
}