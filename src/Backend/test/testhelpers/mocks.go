package testhelpers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/Turgho/Aluguei/internal/domain/entities"
)

// MockOwnerUseCase
type MockOwnerUseCase struct {
	mock.Mock
}

func (m *MockOwnerUseCase) CreateOwner(ctx context.Context, name, email, phone, cpf, password string, birthDate *string) (*entities.Owner, error) {
	args := m.Called(ctx, name, email, phone, cpf, password, birthDate)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerUseCase) GetOwnerByID(ctx context.Context, id uuid.UUID) (*entities.Owner, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerUseCase) GetOwnerByEmail(ctx context.Context, email string) (*entities.Owner, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerUseCase) GetOwners(ctx context.Context, filters entities.OwnerFilters) ([]*entities.Owner, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]*entities.Owner), args.Int(1), args.Error(2)
}

func (m *MockOwnerUseCase) UpdateOwner(ctx context.Context, id uuid.UUID, name, email, phone, cpf *string, birthDate *string) (*entities.Owner, error) {
	args := m.Called(ctx, id, name, email, phone, cpf, birthDate)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func (m *MockOwnerUseCase) DeleteOwner(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockTenantUseCase
type MockTenantUseCase struct {
	mock.Mock
}

func (m *MockTenantUseCase) CreateTenant(ctx context.Context, ownerID uuid.UUID, name, email, phone, cpf string, birthDate *string) (*entities.Tenant, error) {
	args := m.Called(ctx, ownerID, name, email, phone, cpf, birthDate)
	return args.Get(0).(*entities.Tenant), args.Error(1)
}

func (m *MockTenantUseCase) GetTenantByID(ctx context.Context, id uuid.UUID) (*entities.Tenant, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Tenant), args.Error(1)
}

func (m *MockTenantUseCase) GetTenants(ctx context.Context, filters entities.TenantFilters) ([]*entities.Tenant, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]*entities.Tenant), args.Int(1), args.Error(2)
}

func (m *MockTenantUseCase) UpdateTenant(ctx context.Context, id uuid.UUID, name, email, phone, cpf *string, birthDate *string) (*entities.Tenant, error) {
	args := m.Called(ctx, id, name, email, phone, cpf, birthDate)
	return args.Get(0).(*entities.Tenant), args.Error(1)
}

func (m *MockTenantUseCase) DeleteTenant(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTenantUseCase) GetTenantsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entities.Tenant, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*entities.Tenant), args.Error(1)
}

// MockPropertyUseCase
type MockPropertyUseCase struct {
	mock.Mock
}

func (m *MockPropertyUseCase) CreateProperty(ctx context.Context, ownerID uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64) (*entities.Property, error) {
	args := m.Called(ctx, ownerID, title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) GetPropertyByID(ctx context.Context, id uuid.UUID) (*entities.Property, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) GetProperties(ctx context.Context, filters entities.PropertyFilters) ([]*entities.Property, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]*entities.Property), args.Int(1), args.Error(2)
}

func (m *MockPropertyUseCase) UpdateProperty(ctx context.Context, id uuid.UUID, title, description, address, city, state, zipCode *string, bedrooms, bathrooms, area *int, rentAmount *float64, status *string) (*entities.Property, error) {
	args := m.Called(ctx, id, title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount, status)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) DeleteProperty(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPropertyUseCase) GetPropertiesByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entities.Property, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*entities.Property), args.Error(1)
}

// MockContractUseCase
type MockContractUseCase struct {
	mock.Mock
}

func (m *MockContractUseCase) CreateContract(ctx context.Context, propertyID, tenantID uuid.UUID, startDate string, endDate *string, monthlyRent float64, paymentDueDay int, status string) (*entities.Contract, error) {
	args := m.Called(ctx, propertyID, tenantID, startDate, endDate, monthlyRent, paymentDueDay, status)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

func (m *MockContractUseCase) GetContractByID(ctx context.Context, id uuid.UUID) (*entities.Contract, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

func (m *MockContractUseCase) GetContracts(ctx context.Context, filters entities.ContractFilters) ([]*entities.Contract, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]*entities.Contract), args.Int(1), args.Error(2)
}

func (m *MockContractUseCase) UpdateContract(ctx context.Context, id uuid.UUID, startDate, endDate *string, monthlyRent *float64, paymentDueDay *int, status *string) (*entities.Contract, error) {
	args := m.Called(ctx, id, startDate, endDate, monthlyRent, paymentDueDay, status)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

func (m *MockContractUseCase) DeleteContract(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockContractUseCase) GetContractsByProperty(ctx context.Context, propertyID uuid.UUID) ([]*entities.Contract, error) {
	args := m.Called(ctx, propertyID)
	return args.Get(0).([]*entities.Contract), args.Error(1)
}

func (m *MockContractUseCase) GetContractsByTenant(ctx context.Context, tenantID uuid.UUID) ([]*entities.Contract, error) {
	args := m.Called(ctx, tenantID)
	return args.Get(0).([]*entities.Contract), args.Error(1)
}

func (m *MockContractUseCase) GetActiveContractByProperty(ctx context.Context, propertyID uuid.UUID) (*entities.Contract, error) {
	args := m.Called(ctx, propertyID)
	return args.Get(0).(*entities.Contract), args.Error(1)
}

// MockPaymentUseCase
type MockPaymentUseCase struct {
	mock.Mock
}

func (m *MockPaymentUseCase) CreatePayment(ctx context.Context, contractID uuid.UUID, dueDate string, amount float64, paymentMethod, notes string) (*entities.Payment, error) {
	args := m.Called(ctx, contractID, dueDate, amount, paymentMethod, notes)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (m *MockPaymentUseCase) GetPaymentByID(ctx context.Context, id uuid.UUID) (*entities.Payment, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (m *MockPaymentUseCase) GetPayments(ctx context.Context, filters entities.PaymentFilters) ([]*entities.Payment, int, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]*entities.Payment), args.Int(1), args.Error(2)
}

func (m *MockPaymentUseCase) UpdatePayment(ctx context.Context, id uuid.UUID, paidDate *string, amountPaid *float64, status, paymentMethod, notes *string) (*entities.Payment, error) {
	args := m.Called(ctx, id, paidDate, amountPaid, status, paymentMethod, notes)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (m *MockPaymentUseCase) DeletePayment(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPaymentUseCase) GetPaymentsByContract(ctx context.Context, contractID uuid.UUID) ([]*entities.Payment, error) {
	args := m.Called(ctx, contractID)
	return args.Get(0).([]*entities.Payment), args.Error(1)
}

func (m *MockPaymentUseCase) GetOverduePayments(ctx context.Context) ([]*entities.Payment, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entities.Payment), args.Error(1)
}

func (m *MockPaymentUseCase) GetPaymentsByPeriod(ctx context.Context, startDate, endDate string) ([]*entities.Payment, error) {
	args := m.Called(ctx, startDate, endDate)
	return args.Get(0).([]*entities.Payment), args.Error(1)
}