package repositories

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/repositories"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
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
	testOwner    *models.Owner
	testProperty *models.Property
	testTenant   *models.Tenant
	testContract *models.Contract
}

func (suite *RepositoriesTestSuite) SetupTest() {
	// Setup database
	db, err := fixtures.SetupTestDB()
	suite.NoError(err)
	suite.db = db

	// Initialize repositories
	suite.ownerRepo = repositories.NewOwnerRepository(db)
	suite.propertyRepo = repositories.NewPropertyRepository(db)
	suite.tenantRepo = repositories.NewTenantRepository(db)
	suite.contractRepo = repositories.NewContractRepository(db)
	suite.paymentRepo = repositories.NewPaymentRepository(db)

	// Create test data
	suite.setupTestData()
}

func (suite *RepositoriesTestSuite) setupTestData() {
	// Create owner
	owner := fixtures.CreateTestOwner()
	err := suite.ownerRepo.Create(owner)
	suite.NoError(err)
	suite.testOwner = owner

	// Create property
	property := fixtures.CreateTestProperty(owner.ID)
	err = suite.propertyRepo.Create(property)
	suite.NoError(err)
	suite.testProperty = property

	// Create tenant
	tenant := fixtures.CreateTestTenant(owner.ID)
	err = suite.tenantRepo.Create(tenant)
	suite.NoError(err)
	suite.testTenant = tenant

	// Create contract
	contract := fixtures.CreateTestContract(property.ID, tenant.ID)
	err = suite.contractRepo.Create(contract)
	suite.NoError(err)
	suite.testContract = contract
}

func (suite *RepositoriesTestSuite) TearDownTest() {
	// Cleanup - o banco em memória é automaticamente limpo
}

func TestRepositoriesTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoriesTestSuite))
}
