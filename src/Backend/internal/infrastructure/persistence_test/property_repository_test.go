package persistence_test

import (
	"context"
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PropertyRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repositories.PropertyRepository
}

func (suite *PropertyRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	err = db.AutoMigrate(&entities.Property{})
	suite.Require().NoError(err)

	suite.db = db
	suite.repo = persistence.NewPropertyRepository(db)
}

func (suite *PropertyRepositoryTestSuite) TestCreate() {
	ctx := context.Background()
	ownerID := uuid.New()

	property := entities.NewProperty(
		ownerID,
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	err := suite.repo.Create(ctx, property)
	assert.NoError(suite.T(), err)

	// Verify the property was created
	var count int64
	suite.db.Model(&entities.Property{}).Count(&count)
	assert.Equal(suite.T(), int64(1), count)
}

func (suite *PropertyRepositoryTestSuite) TestGetByID() {
	ctx := context.Background()
	ownerID := uuid.New()

	property := entities.NewProperty(
		ownerID,
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	err := suite.repo.Create(ctx, property)
	suite.Require().NoError(err)

	foundProperty, err := suite.repo.GetByID(ctx, property.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), property.ID, foundProperty.ID)
	assert.Equal(suite.T(), property.Title, foundProperty.Title)
}

func (suite *PropertyRepositoryTestSuite) TestGetByOwnerID() {
	ctx := context.Background()
	ownerID := uuid.New()

	property1 := entities.NewProperty(
		ownerID,
		"Property 1",
		"Description 1",
		"Address 1",
		"City 1",
		"State 1",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	property2 := entities.NewProperty(
		ownerID,
		"Property 2",
		"Description 2",
		"Address 2",
		"City 2",
		"State 2",
		"87654321",
		3,
		2,
		75,
		1500.0,
	)

	err := suite.repo.Create(ctx, property1)
	suite.Require().NoError(err)

	err = suite.repo.Create(ctx, property2)
	suite.Require().NoError(err)

	properties, err := suite.repo.GetByOwnerID(ctx, ownerID)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 2)
}

func (suite *PropertyRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()
	ownerID := uuid.New()

	property1 := entities.NewProperty(
		ownerID,
		"Property 1",
		"Description 1",
		"Address 1",
		"City 1",
		"State 1",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	property2 := entities.NewProperty(
		ownerID,
		"Property 2",
		"Description 2",
		"Address 2",
		"City 2",
		"State 2",
		"87654321",
		3,
		2,
		75,
		1500.0,
	)

	err := suite.repo.Create(ctx, property1)
	suite.Require().NoError(err)

	err = suite.repo.Create(ctx, property2)
	suite.Require().NoError(err)

	properties, total, err := suite.repo.GetAll(ctx, 1, 10, "")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), properties, 2)
	assert.Equal(suite.T(), int64(2), total)
}

func (suite *PropertyRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()
	ownerID := uuid.New()

	property := entities.NewProperty(
		ownerID,
		"Original Title",
		"Original Description",
		"Original Address",
		"Original City",
		"Original State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	err := suite.repo.Create(ctx, property)
	suite.Require().NoError(err)

	property.Update(
		"Updated Title",
		"Updated Description",
		"Updated Address",
		"Updated City",
		"Updated State",
		"87654321",
		3,
		2,
		75,
		1500.0,
		entities.PropertyStatusRented,
	)

	err = suite.repo.Update(ctx, property)
	assert.NoError(suite.T(), err)

	updatedProperty, err := suite.repo.GetByID(ctx, property.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Updated Title", updatedProperty.Title)
	assert.Equal(suite.T(), entities.PropertyStatusRented, updatedProperty.Status)
}

func (suite *PropertyRepositoryTestSuite) TestDelete() {
	ctx := context.Background()
	ownerID := uuid.New()

	property := entities.NewProperty(
		ownerID,
		"Test Property",
		"Test Description",
		"Test Address",
		"Test City",
		"Test State",
		"12345678",
		2,
		1,
		50,
		1000.0,
	)

	err := suite.repo.Create(ctx, property)
	suite.Require().NoError(err)

	err = suite.repo.Delete(ctx, property.ID)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByID(ctx, property.ID)
	assert.Error(suite.T(), err)
}

func TestPropertyRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PropertyRepositoryTestSuite))
}