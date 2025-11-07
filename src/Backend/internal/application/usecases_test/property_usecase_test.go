package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPropertyRepository is a mock implementation of PropertyRepository
type MockPropertyRepository struct {
	mock.Mock
}

func (m *MockPropertyRepository) Create(ctx context.Context, property *entities.Property) error {
	args := m.Called(ctx, property)
	return args.Error(0)
}

func (m *MockPropertyRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Property, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyRepository) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Property, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*entities.Property), args.Error(1)
}

func (m *MockPropertyRepository) GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Property, int64, error) {
	args := m.Called(ctx, page, limit, status)
	return args.Get(0).([]*entities.Property), args.Get(1).(int64), args.Error(2)
}

func (m *MockPropertyRepository) Update(ctx context.Context, property *entities.Property) error {
	args := m.Called(ctx, property)
	return args.Error(0)
}

func (m *MockPropertyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestPropertyUseCase_CreateProperty(t *testing.T) {
	mockRepo := new(MockPropertyRepository)
	useCase := usecases.NewPropertyUseCase(mockRepo)

	ownerID := uuid.New()
	ctx := context.Background()

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Property")).Return(nil)

	property, err := useCase.CreateProperty(
		ctx,
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

	assert.NoError(t, err)
	assert.NotNil(t, property)
	assert.Equal(t, ownerID, property.OwnerID)
	assert.Equal(t, "Test Property", property.Title)
	mockRepo.AssertExpectations(t)
}

func TestPropertyUseCase_CreateProperty_Error(t *testing.T) {
	mockRepo := new(MockPropertyRepository)
	useCase := usecases.NewPropertyUseCase(mockRepo)

	ownerID := uuid.New()
	ctx := context.Background()

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Property")).Return(errors.New("database error"))

	property, err := useCase.CreateProperty(
		ctx,
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

	assert.Error(t, err)
	assert.Nil(t, property)
	mockRepo.AssertExpectations(t)
}

func TestPropertyUseCase_GetProperty(t *testing.T) {
	mockRepo := new(MockPropertyRepository)
	useCase := usecases.NewPropertyUseCase(mockRepo)

	propertyID := uuid.New()
	expectedProperty := &entities.Property{
		ID:    propertyID,
		Title: "Test Property",
	}
	ctx := context.Background()

	mockRepo.On("GetByID", ctx, propertyID).Return(expectedProperty, nil)

	property, err := useCase.GetProperty(ctx, propertyID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProperty, property)
	mockRepo.AssertExpectations(t)
}