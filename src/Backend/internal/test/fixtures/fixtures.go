package fixtures

import (
	"time"

	"github.com/Turgho/Aluguei/internal/models"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB cria e configura o banco de testes
func SetupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar tabelas
	err = db.AutoMigrate(
		&models.Owner{},
		&models.Property{},
		&models.Tenant{},
		&models.Contract{},
		&models.Payment{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTestOwner cria um owner de teste
func CreateTestOwner() *models.Owner {
	return &models.Owner{
		Name:     "João Silva",
		Email:    "joao@test.com",
		Phone:    "11999999999",
		Password: "senha123",
	}
}

// CreateTestProperty cria uma propriedade de teste
func CreateTestProperty(ownerID uuid.UUID) *models.Property {
	return &models.Property{
		OwnerID:      ownerID,
		Title:        "Apartamento Teste",
		Description:  "Descrição do apartamento teste",
		Address:      "Rua Teste, 123",
		Number:       "123",
		Complement:   "Apto 45",
		Neighborhood: "Centro",
		City:         "São Paulo",
		State:        "SP",
		ZipCode:      "01234567",
		Type:         models.PropertyTypeApartment,
		Bedrooms:     2,
		Bathrooms:    1,
		Area:         60.0,
		RentAmount:   1500.00,
		Status:       models.PropertyStatusAvailable,
	}
}

// CreateTestTenant cria um tenant de teste
func CreateTestTenant(ownerID uuid.UUID) *models.Tenant {
	return &models.Tenant{
		OwnerID: ownerID,
		Name:    "Maria Santos",
		Email:   "maria@test.com",
		Phone:   "11888888888",
		CPF:     "12345678901",
	}
}

// CreateTestContract cria um contrato de teste
func CreateTestContract(propertyID, tenantID uuid.UUID) *models.Contract {
	return &models.Contract{
		PropertyID:    propertyID,
		TenantID:      tenantID,
		StartDate:     time.Now(),
		EndDate:       time.Now().AddDate(1, 0, 0), // 1 ano
		MonthlyRent:   1500.00,
		DepositAmount: 3000.00,
		PaymentDueDay: 5,
		Status:        models.ContractStatusActive,
	}
}

// CreateTestPayment cria um pagamento de teste
func CreateTestPayment(contractID uuid.UUID) *models.Payment {
	return &models.Payment{
		ContractID:     contractID,
		DueDate:        time.Now().AddDate(0, 1, 0), // Próximo mês
		Amount:         1500.00,
		PaidAmount:     0,
		LateFee:        0,
		Method:         models.PaymentMethodPIX,
		Status:         models.PaymentStatusPending,
		ReferenceMonth: time.Now(),
	}
}

// CreateTestOwnerWithID cria owner com ID específico (para testes de update)
func CreateTestOwnerWithID() *models.Owner {
	owner := CreateTestOwner()
	owner.ID = uuid.New()
	return owner
}

// CreateTestPropertyWithID cria property com ID específico
func CreateTestPropertyWithID(ownerID uuid.UUID) *models.Property {
	property := CreateTestProperty(ownerID)
	property.ID = uuid.New()
	return property
}
