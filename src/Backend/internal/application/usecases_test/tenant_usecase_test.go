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
)

type MockTenantRepository struct {
	mock.Mock
}

func (m *MockTenantRepository) Create(ctx context.Context, tenant *entities.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}

func (m *MockTenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Tenant, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Tenant), args.Error(1)
}

func (m *MockTenantRepository) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*entities.Tenant, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*entities.Tenant), args.Error(1)
}

func (m *MockTenantRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Tenant, int64, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]*entities.Tenant), args.Get(1).(int64), args.Error(2)
}

func (m *MockTenantRepository) Update(ctx context.Context, tenant *entities.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}

func (m *MockTenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestTenantUseCase_CreateTenant(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	useCase := usecases.NewTenantUseCase(mockRepo)

	ownerID := uuid.New()
	ctx := context.Background()
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Tenant")).Return(nil)

	tenant, err := useCase.CreateTenant(
		ctx,
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		&birthDate,
	)

	assert.NoError(t, err)
	assert.NotNil(t, tenant)
	assert.Equal(t, "Jane Doe", tenant.Name)
	assert.Equal(t, ownerID, tenant.OwnerID)
	mockRepo.AssertExpectations(t)
}

func TestTenantUseCase_GetTenant(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	useCase := usecases.NewTenantUseCase(mockRepo)

	tenantID := uuid.New()
	expectedTenant := &entities.Tenant{
		ID:   tenantID,
		Name: "Jane Doe",
	}
	ctx := context.Background()

	mockRepo.On("GetByID", ctx, tenantID).Return(expectedTenant, nil)

	tenant, err := useCase.GetTenant(ctx, tenantID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTenant, tenant)
	mockRepo.AssertExpectations(t)
}

func TestTenantUseCase_GetTenantsByOwner(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	useCase := usecases.NewTenantUseCase(mockRepo)

	ownerID := uuid.New()
	expectedTenants := []*entities.Tenant{
		{ID: uuid.New(), Name: "Tenant 1", OwnerID: ownerID},
		{ID: uuid.New(), Name: "Tenant 2", OwnerID: ownerID},
	}
	ctx := context.Background()

	mockRepo.On("GetByOwnerID", ctx, ownerID).Return(expectedTenants, nil)

	tenants, err := useCase.GetTenantsByOwner(ctx, ownerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTenants, tenants)
	assert.Len(t, tenants, 2)
	mockRepo.AssertExpectations(t)
}

func TestTenantUseCase_CreateTenant_Error(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	useCase := usecases.NewTenantUseCase(mockRepo)

	ctx := context.Background()

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Tenant")).Return(errors.New("database error"))

	tenant, err := useCase.CreateTenant(
		ctx,
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		uuid.New(),
		nil,
	)

	assert.Error(t, err)
	assert.Nil(t, tenant)
	mockRepo.AssertExpectations(t)
}