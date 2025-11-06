// tests/repositories/repositories_suite_test.go
package repositories

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/repositories"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type RepositoriesTestSuite struct {
	suite.Suite
	db           *gorm.DB
	ownerRepo    repositories.OwnerRepository
	propertyRepo repositories.PropertyRepository
	tenantRepo   repositories.TenantRepository
	contractRepo repositories.ContractRepository
	paymentRepo  repositories.PaymentRepository
}

func (suite *RepositoriesTestSuite) SetupTest() {
	// Setup database - criar NOVO banco para CADA teste
	db, err := fixtures.SetupTestDB()
	suite.NoError(err)
	suite.db = db

	// Initialize repositories
	suite.ownerRepo = repositories.NewOwnerRepository(db)
	suite.propertyRepo = repositories.NewPropertyRepository(db)
	suite.tenantRepo = repositories.NewTenantRepository(db)
	suite.contractRepo = repositories.NewContractRepository(db)
	suite.paymentRepo = repositories.NewPaymentRepository(db)
}

func (suite *RepositoriesTestSuite) TearDownTest() {
	// Cleanup - Fechar a conexão para garantir que o banco em memória seja destruído
	if suite.db != nil {
		if sqlDB, err := suite.db.DB(); err == nil {
			sqlDB.Close()
		}
	}
}

// Helper methods para criar dados únicos em cada teste
func (suite *RepositoriesTestSuite) createUniqueTestOwner() *models.Owner {
	owner := fixtures.CreateTestOwner()
	owner.Email = suite.generateUniqueEmail()
	owner.CPF = suite.generateUniqueCPF()

	err := suite.ownerRepo.Create(owner)

	suite.Nil(err, "Deveria criar owner sem erro")

	if err != nil {
		suite.T().Logf("ERRO ao criar owner: %v", err)
		suite.FailNow("Falha ao criar owner de teste")
		return nil
	}

	suite.NotEmpty(owner.ID, "Owner deveria ter ID após criação")
	return owner
}

func (suite *RepositoriesTestSuite) createUniqueTestProperty(ownerID uuid.UUID) *models.Property {
	property := fixtures.CreateTestProperty(ownerID)
	property.Title = "Unique Property " + suite.generateUniqueID()

	err := suite.propertyRepo.Create(property)

	suite.Nil(err, "Deveria criar property sem erro")

	if err != nil {
		suite.T().Logf("ERRO ao criar property: %v", err)
		suite.FailNow("Falha ao criar property de teste")
		return nil
	}

	suite.NotEmpty(property.ID, "Property deveria ter ID após criação")
	return property
}

func (suite *RepositoriesTestSuite) createUniqueTestTenant(ownerID uuid.UUID) *models.Tenant {
	tenant := fixtures.CreateTestTenant(ownerID)
	tenant.Email = suite.generateUniqueEmail()
	tenant.CPF = suite.generateUniqueCPF()

	err := suite.tenantRepo.Create(tenant)

	suite.Nil(err, "Deveria criar tenant sem erro")

	// Verificação adicional para garantir que foi criado
	if err != nil {
		suite.T().Logf("ERRO ao criar tenant: %v", err)
		suite.FailNow("Falha ao criar tenant de teste")
		return nil
	}

	suite.NotEmpty(tenant.ID, "Tenant deveria ter ID após criação")
	return tenant
}

// Helper para gerar emails únicos
func (suite *RepositoriesTestSuite) generateUniqueEmail() string {
	return "test_" + suite.generateUniqueID() + "@test.com"
}

// Helper para gerar IDs únicos
func (suite *RepositoriesTestSuite) generateUniqueID() string {
	return uuid.New().String()[:8] // Usa os primeiros 8 caracteres do UUID
}

// Helper para gerar CPFs únicos
func (suite *RepositoriesTestSuite) generateUniqueCPF() string {
	return suite.generateUniqueID() + "00000" // CPF fictício único
}

// Helper para criar contrato
func (suite *RepositoriesTestSuite) createTestContract(propertyID, tenantID uuid.UUID) *models.Contract {
	contract := fixtures.CreateTestContract(propertyID, tenantID)
	err := suite.contractRepo.Create(contract)
	suite.Nil(err)
	return contract
}

func TestRepositoriesTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoriesTestSuite))
}
