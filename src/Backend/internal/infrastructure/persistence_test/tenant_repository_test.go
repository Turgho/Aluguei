package persistence_test

import (
	"context"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TenantRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repositories.TenantRepository
}

func (suite *TenantRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	err = db.AutoMigrate(&entities.Tenant{})
	suite.Require().NoError(err)

	suite.db = db
	suite.repo = persistence.NewTenantRepository(db)
}

func (suite *TenantRepositoryTestSuite) TestCreate() {
	ctx := context.Background()
	ownerID := uuid.New()
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		&birthDate,
	)

	err := suite.repo.Create(ctx, tenant)
	assert.NoError(suite.T(), err)

	var count int64
	suite.db.Model(&entities.Tenant{}).Count(&count)
	assert.Equal(suite.T(), int64(1), count)
}

func (suite *TenantRepositoryTestSuite) TestGetByID() {
	ctx := context.Background()
	ownerID := uuid.New()

	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		nil,
	)

	err := suite.repo.Create(ctx, tenant)
	suite.Require().NoError(err)

	foundTenant, err := suite.repo.GetByID(ctx, tenant.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), tenant.ID, foundTenant.ID)
	assert.Equal(suite.T(), tenant.Name, foundTenant.Name)
	assert.Equal(suite.T(), tenant.OwnerID, foundTenant.OwnerID)
}

func (suite *TenantRepositoryTestSuite) TestGetByOwnerID() {
	ctx := context.Background()
	ownerID := uuid.New()

	tenant1 := entities.NewTenant("Jane Doe", "jane@example.com", "+5511888888888", "98765432100", ownerID, nil)
	tenant2 := entities.NewTenant("John Smith", "john@example.com", "+5511777777777", "12345678900", ownerID, nil)

	err := suite.repo.Create(ctx, tenant1)
	suite.Require().NoError(err)

	err = suite.repo.Create(ctx, tenant2)
	suite.Require().NoError(err)

	tenants, err := suite.repo.GetByOwnerID(ctx, ownerID)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tenants, 2)
}

func (suite *TenantRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()
	ownerID := uuid.New()

	tenant1 := entities.NewTenant("Jane Doe", "jane@example.com", "+5511888888888", "98765432100", ownerID, nil)
	tenant2 := entities.NewTenant("John Smith", "john@example.com", "+5511777777777", "12345678900", ownerID, nil)

	err := suite.repo.Create(ctx, tenant1)
	suite.Require().NoError(err)

	err = suite.repo.Create(ctx, tenant2)
	suite.Require().NoError(err)

	tenants, total, err := suite.repo.GetAll(ctx, 1, 10)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tenants, 2)
	assert.Equal(suite.T(), int64(2), total)
}

func (suite *TenantRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()
	ownerID := uuid.New()

	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		nil,
	)

	err := suite.repo.Create(ctx, tenant)
	suite.Require().NoError(err)

	tenant.UpdateProfile("Jane Smith", "janesmith@example.com", "+5511777777777", "12345678900", nil)

	err = suite.repo.Update(ctx, tenant)
	assert.NoError(suite.T(), err)

	updatedTenant, err := suite.repo.GetByID(ctx, tenant.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Jane Smith", updatedTenant.Name)
	assert.Equal(suite.T(), "janesmith@example.com", updatedTenant.Email)
}

func (suite *TenantRepositoryTestSuite) TestDelete() {
	ctx := context.Background()
	ownerID := uuid.New()

	tenant := entities.NewTenant(
		"Jane Doe",
		"jane@example.com",
		"+5511888888888",
		"98765432100",
		ownerID,
		nil,
	)

	err := suite.repo.Create(ctx, tenant)
	suite.Require().NoError(err)

	err = suite.repo.Delete(ctx, tenant.ID)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByID(ctx, tenant.ID)
	assert.Error(suite.T(), err)
}

func TestTenantRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TenantRepositoryTestSuite))
}