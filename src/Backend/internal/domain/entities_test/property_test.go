package entities_test

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProperty(t *testing.T) {
	ownerID := uuid.New()
	property := entities.NewProperty(
		ownerID,
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	assert.NotEqual(t, uuid.Nil, property.ID)
	assert.Equal(t, ownerID, property.OwnerID)
	assert.Equal(t, "Test Property", property.Title)
	assert.Equal(t, "Test Description", property.Description)
	assert.Equal(t, "Test Address", property.Address)
	assert.Equal(t, "Test City", property.City)
	assert.Equal(t, "Test State", property.State)
	assert.Equal(t, "12345678", property.ZipCode)
	assert.Equal(t, 2, property.Bedrooms)
	assert.Equal(t, 1, property.Bathrooms)
	assert.Equal(t, 50, property.Area)
	assert.Equal(t, 1000.0, property.RentAmount)
	assert.Equal(t, entities.PropertyStatusAvailable, property.Status)
}

func TestProperty_UpdateStatus(t *testing.T) {
	property := entities.NewProperty(
		uuid.New(),
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	property.UpdateStatus(entities.PropertyStatusRented)
	assert.Equal(t, entities.PropertyStatusRented, property.Status)
}

func TestProperty_IsAvailable(t *testing.T) {
	property := entities.NewProperty(
		uuid.New(),
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	assert.True(t, property.IsAvailable())

	property.UpdateStatus(entities.PropertyStatusRented)
	assert.False(t, property.IsAvailable())
}