package integration

import (
	"context"
	"testing"

	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/Turgho/Aluguei/test/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DatabaseIntegrationTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *DatabaseIntegrationTestSuite) SetupTest() {
	db, err := testhelpers.SetupTestDB()
	suite.Require().NoError(err)
	suite.db = db
}

func (suite *DatabaseIntegrationTestSuite) TearDownTest() {
	err := testhelpers.CleanupTestDB(suite.db)
	suite.Require().NoError(err)
}

func (suite *DatabaseIntegrationTestSuite) TestOwnerPropertyRelationship() {
	ctx := context.Background()
	
	// Create repositories
	ownerRepo := persistence.NewOwnerRepository(suite.db)
	propertyRepo := persistence.NewPropertyRepository(suite.db)
	
	// Create owner
	owner := testhelpers.CreateTestOwner()
	err := ownerRepo.Create(ctx, owner)
	suite.Require().NoError(err)
	
	// Create property for owner
	property := testhelpers.CreateTestProperty(owner.ID)
	err = propertyRepo.Create(ctx, property)
	suite.Require().NoError(err)
	
	// Verify relationship
	properties, err := propertyRepo.GetByOwnerID(ctx, owner.ID)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 1)
	assert.Equal(suite.T(), property.ID, properties[0].ID)
	assert.Equal(suite.T(), owner.ID, properties[0].OwnerID)
}

func (suite *DatabaseIntegrationTestSuite) TestTenantOwnerRelationship() {
	ctx := context.Background()
	
	// Create repositories
	ownerRepo := persistence.NewOwnerRepository(suite.db)
	tenantRepo := persistence.NewTenantRepository(suite.db)
	
	// Create owner
	owner := testhelpers.CreateTestOwner()
	err := ownerRepo.Create(ctx, owner)
	suite.Require().NoError(err)
	
	// Create tenant for owner
	tenant := testhelpers.CreateTestTenant(owner.ID)
	err = tenantRepo.Create(ctx, tenant)
	suite.Require().NoError(err)
	
	// Verify relationship
	tenants, err := tenantRepo.GetByOwnerID(ctx, owner.ID)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tenants, 1)
	assert.Equal(suite.T(), tenant.ID, tenants[0].ID)
	assert.Equal(suite.T(), owner.ID, tenants[0].OwnerID)
}

func (suite *DatabaseIntegrationTestSuite) TestCompleteWorkflow() {
	ctx := context.Background()
	
	// Create repositories
	ownerRepo := persistence.NewOwnerRepository(suite.db)
	tenantRepo := persistence.NewTenantRepository(suite.db)
	propertyRepo := persistence.NewPropertyRepository(suite.db)
	contractRepo := persistence.NewContractRepository(suite.db)
	paymentRepo := persistence.NewPaymentRepository(suite.db)
	
	// 1. Create owner
	owner := testhelpers.CreateTestOwner()
	err := ownerRepo.Create(ctx, owner)
	suite.Require().NoError(err)
	
	// 2. Create tenant
	tenant := testhelpers.CreateTestTenant(owner.ID)
	err = tenantRepo.Create(ctx, tenant)
	suite.Require().NoError(err)
	
	// 3. Create property
	property := testhelpers.CreateTestProperty(owner.ID)
	err = propertyRepo.Create(ctx, property)
	suite.Require().NoError(err)
	
	// 4. Create contract
	contract := testhelpers.CreateTestContract(property.ID, tenant.ID)
	err = contractRepo.Create(ctx, contract)
	suite.Require().NoError(err)
	
	// 5. Create payment
	payment := testhelpers.CreateTestPayment(contract.ID)
	err = paymentRepo.Create(ctx, payment)
	suite.Require().NoError(err)
	
	// Verify all relationships
	retrievedOwner, err := ownerRepo.GetByID(ctx, owner.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), owner.Name, retrievedOwner.Name)
	
	retrievedProperty, err := propertyRepo.GetByID(ctx, property.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), property.Title, retrievedProperty.Title)
	
	retrievedContract, err := contractRepo.GetByID(ctx, contract.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), contract.MonthlyRent, retrievedContract.MonthlyRent)
	
	retrievedPayment, err := paymentRepo.GetByID(ctx, payment.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), payment.Amount, retrievedPayment.Amount)
}

func TestDatabaseIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationTestSuite))
}