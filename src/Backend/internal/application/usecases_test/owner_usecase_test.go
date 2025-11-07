package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockOwnerRepository is a mock implementation of OwnerRepository
type MockOwnerRepository struct {
	mock.Mock
}

func (m *MockOwnerRepository) Create(ctx context.Context, owner *entities.Owner) error {
	args := m.Called(ctx, owner)
	return args.Error(0)
}

func (m *MockOwnerRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Owner, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerRepository) GetByEmail(ctx context.Context, email string) (*entities.Owner, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Owner, int64, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]*entities.Owner), args.Get(1).(int64), args.Error(2)
}

func (m *MockOwnerRepository) Update(ctx context.Context, owner *entities.Owner) error {
	args := m.Called(ctx, owner)
	return args.Error(0)
}

func (m *MockOwnerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestOwnerUseCase_CreateOwner(t *testing.T) {
	mockRepo := new(MockOwnerRepository)
	useCase := usecases.NewOwnerUseCase(mockRepo)

	ctx := context.Background()
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Owner")).Return(nil)

	owner, err := useCase.CreateOwner(
		ctx,
		"John Doe",
		"john@example.com",
		"password123",
		"+5511999999999",
		"12345678901",
		&birthDate,
	)

	assert.NoError(t, err)
	assert.NotNil(t, owner)
	assert.Equal(t, "John Doe", owner.Name)
	assert.Equal(t, "john@example.com", owner.Email)
	assert.NotEqual(t, "password123", owner.Password) // Should be hashed
	mockRepo.AssertExpectations(t)
}

func TestOwnerUseCase_ValidatePassword(t *testing.T) {
	mockRepo := new(MockOwnerRepository)
	useCase := usecases.NewOwnerUseCase(mockRepo)

	ctx := context.Background()
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	expectedOwner := &entities.Owner{
		ID:       uuid.New(),
		Email:    "john@example.com",
		Password: string(hashedPassword),
	}

	mockRepo.On("GetByEmail", ctx, "john@example.com").Return(expectedOwner, nil)

	owner, err := useCase.ValidatePassword(ctx, "john@example.com", password)

	assert.NoError(t, err)
	assert.Equal(t, expectedOwner, owner)
	mockRepo.AssertExpectations(t)
}

func TestOwnerUseCase_ValidatePassword_InvalidPassword(t *testing.T) {
	mockRepo := new(MockOwnerRepository)
	useCase := usecases.NewOwnerUseCase(mockRepo)

	ctx := context.Background()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	expectedOwner := &entities.Owner{
		ID:       uuid.New(),
		Email:    "john@example.com",
		Password: string(hashedPassword),
	}

	mockRepo.On("GetByEmail", ctx, "john@example.com").Return(expectedOwner, nil)

	owner, err := useCase.ValidatePassword(ctx, "john@example.com", "wrongpassword")

	assert.Error(t, err)
	assert.Nil(t, owner)
	mockRepo.AssertExpectations(t)
}

func TestOwnerUseCase_ValidatePassword_UserNotFound(t *testing.T) {
	mockRepo := new(MockOwnerRepository)
	useCase := usecases.NewOwnerUseCase(mockRepo)

	ctx := context.Background()

	mockRepo.On("GetByEmail", ctx, "nonexistent@example.com").Return((*entities.Owner)(nil), errors.New("user not found"))

	owner, err := useCase.ValidatePassword(ctx, "nonexistent@example.com", "password123")

	assert.Error(t, err)
	assert.Nil(t, owner)
	mockRepo.AssertExpectations(t)
}