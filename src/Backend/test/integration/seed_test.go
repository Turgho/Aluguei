package integration

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/infrastructure/seeds"
	"github.com/Turgho/Aluguei/test/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SeedTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *SeedTestSuite) SetupTest() {
	db, err := testhelpers.SetupTestDB()
	suite.Require().NoError(err)
	suite.db = db
}

func (suite *SeedTestSuite) TestSeedAll() {
	seeder := seeds.NewSeeder(suite.db)
	
	err := seeder.SeedAll()
	assert.NoError(suite.T(), err)
	
	// Verify owners
	var ownerCount int64
	suite.db.Model(&entities.Owner{}).Count(&ownerCount)
	assert.Equal(suite.T(), int64(3), ownerCount)
	
	// Verify tenants
	var tenantCount int64
	suite.db.Model(&entities.Tenant{}).Count(&tenantCount)
	assert.Equal(suite.T(), int64(3), tenantCount)
	
	// Verify properties
	var propertyCount int64
	suite.db.Model(&entities.Property{}).Count(&propertyCount)
	assert.Equal(suite.T(), int64(4), propertyCount)
	
	// Verify contracts
	var contractCount int64
	suite.db.Model(&entities.Contract{}).Count(&contractCount)
	assert.Equal(suite.T(), int64(2), contractCount)
	
	// Verify payments
	var paymentCount int64
	suite.db.Model(&entities.Payment{}).Count(&paymentCount)
	assert.Equal(suite.T(), int64(4), paymentCount)
}

func (suite *SeedTestSuite) TestSeedIdempotent() {
	seeder := seeds.NewSeeder(suite.db)
	
	// Run seeds twice
	err := seeder.SeedAll()
	assert.NoError(suite.T(), err)
	
	err = seeder.SeedAll()
	assert.NoError(suite.T(), err)
	
	// Should still have same counts
	var ownerCount int64
	suite.db.Model(&entities.Owner{}).Count(&ownerCount)
	assert.Equal(suite.T(), int64(3), ownerCount)
}

func (suite *SeedTestSuite) TestSeedRelationships() {
	seeder := seeds.NewSeeder(suite.db)
	err := seeder.SeedAll()
	assert.NoError(suite.T(), err)
	
	// Test owner-tenant relationship
	var tenant entities.Tenant
	err = suite.db.First(&tenant, "email = ?", "ana.costa@email.com").Error
	assert.NoError(suite.T(), err)
	
	var owner entities.Owner
	err = suite.db.First(&owner, "id = ?", tenant.OwnerID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "João Silva", owner.Name)
	
	// Test property-contract relationship
	var contract entities.Contract
	err = suite.db.First(&contract, "status = ?", entities.ContractStatusActive).Error
	assert.NoError(suite.T(), err)
	
	var property entities.Property
	err = suite.db.First(&property, "id = ?", contract.PropertyID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Casa Vila Madalena - 3 Quartos", property.Title)
}

func TestSeedTestSuite(t *testing.T) {
	suite.Run(t, new(SeedTestSuite))
}