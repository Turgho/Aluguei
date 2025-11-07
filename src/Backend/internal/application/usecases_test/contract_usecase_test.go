package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockContractRepository struct {
	mock.Mock
}

func (m *MockContractRepository) Create(ctx context.Context, contract *entities.Contract) error {
	args := m.Called(ctx, contract)
	return args.Error(0)
}

func (m *MockContractRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Contract, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

func (m *MockContractRepository) GetByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*entities.Contract, error) {
	args := m.Called(ctx, propertyID)
	return args.Get(0).([]*entities.Contract), args.Error(1)
}

func (m *MockContractRepository) GetByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*entities.Contract, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0).([]*entities.Contract), args.Error(1)
}

func (m *MockContractRepository) GetActiveByPropertyID(ctx context.Context, propertyID uuid.UUID) (*entities.Contract, error) {
	args := m.Called(ctx, propertyID)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

func (m *MockContractRepository) GetAll(ctx context.Context, page, limit int, status string) ([]*entities.Contract, int64, error) {
	args := m.Called(ctx, page, limit, status)
	return args.Get(0).([]*entities.Contract), args.Get(1).(int64), args.Error(2)
}

func (m *MockContractRepository) Update(ctx context.Context, contract *entities.Contract) error {
	args := m.Called(ctx, contract)
	return args.Error(0)
}

func (m *MockContractRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestContractUseCase_CreateContract(t *testing.T) {
	mockRepo := new(MockContractRepository)
	useCase := usecases.NewContractUseCase(mockRepo)

	propertyID := uuid.New()
	tenantID := uuid.New()
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	ctx := context.Background()

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Contract")).Return(nil)

	contract, err := useCase.CreateContract(
		ctx,
		propertyID,
		tenantID,
		startDate,
		endDate,
		1500.0,
		5,
		entities.ContractStatusActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, contract)
	assert.Equal(t, propertyID, contract.PropertyID)
	assert.Equal(t, tenantID, contract.TenantID)
	assert.Equal(t, 1500.0, contract.MonthlyRent)
	assert.Equal(t, 5, contract.PaymentDueDay)
	mockRepo.AssertExpectations(t)
}

func TestContractUseCase_GetActiveContractByProperty(t *testing.T) {
	mockRepo := new(MockContractRepository)
	useCase := usecases.NewContractUseCase(mockRepo)

	propertyID := uuid.New()
	expectedContract := &entities.Contract{
		ID:         uuid.New(),
		PropertyID: propertyID,
		Status:     entities.ContractStatusActive,
	}
	ctx := context.Background()

	mockRepo.On("GetActiveByPropertyID", ctx, propertyID).Return(expectedContract, nil)

	contract, err := useCase.GetActiveContractByProperty(ctx, propertyID)

	assert.NoError(t, err)
	assert.Equal(t, expectedContract, contract)
	mockRepo.AssertExpectations(t)
}