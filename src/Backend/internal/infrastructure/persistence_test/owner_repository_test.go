package persistence_test

import (
	"context"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OwnerRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repositories.OwnerRepository
}

func (suite *OwnerRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	err = db.AutoMigrate(&entities.Owner{})
	suite.Require().NoError(err)

	suite.db = db
	suite.repo = persistence.NewOwnerRepository(db)
}

func (suite *OwnerRepositoryTestSuite) TestCreate() {
	ctx := context.Background()
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		&birthDate,
	)

	err := suite.repo.Create(ctx, owner)
	assert.NoError(suite.T(), err)

	var count int64
	suite.db.Model(&entities.Owner{}).Count(&count)
	assert.Equal(suite.T(), int64(1), count)
}

func (suite *OwnerRepositoryTestSuite) TestGetByID() {
	ctx := context.Background()
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)

	err := suite.repo.Create(ctx, owner)
	suite.Require().NoError(err)

	foundOwner, err := suite.repo.GetByID(ctx, owner.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), owner.ID, foundOwner.ID)
	assert.Equal(suite.T(), owner.Name, foundOwner.Name)
	assert.Equal(suite.T(), owner.Email, foundOwner.Email)
}

func (suite *OwnerRepositoryTestSuite) TestGetByEmail() {
	ctx := context.Background()
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)

	err := suite.repo.Create(ctx, owner)
	suite.Require().NoError(err)

	foundOwner, err := suite.repo.GetByEmail(ctx, "john@example.com")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), owner.ID, foundOwner.ID)
	assert.Equal(suite.T(), owner.Email, foundOwner.Email)
}

func (suite *OwnerRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()

	owner1 := entities.NewOwner("John Doe", "john@example.com", "pass", "+5511999999999", "12345678901", nil)
	owner2 := entities.NewOwner("Jane Doe", "jane@example.com", "pass", "+5511888888888", "98765432100", nil)

	err := suite.repo.Create(ctx, owner1)
	suite.Require().NoError(err)

	err = suite.repo.Create(ctx, owner2)
	suite.Require().NoError(err)

	owners, total, err := suite.repo.GetAll(ctx, 1, 10)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), owners, 2)
	assert.Equal(suite.T(), int64(2), total)
}

func (suite *OwnerRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)

	err := suite.repo.Create(ctx, owner)
	suite.Require().NoError(err)

	owner.UpdateProfile("John Smith", "johnsmith@example.com", "+5511888888888", "98765432100", nil)

	err = suite.repo.Update(ctx, owner)
	assert.NoError(suite.T(), err)

	updatedOwner, err := suite.repo.GetByID(ctx, owner.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "John Smith", updatedOwner.Name)
	assert.Equal(suite.T(), "johnsmith@example.com", updatedOwner.Email)
}

func (suite *OwnerRepositoryTestSuite) TestDelete() {
	ctx := context.Background()
	owner := entities.NewOwner(
		"John Doe",
		"john@example.com",
		"hashedpassword",
		"+5511999999999",
		"12345678901",
		nil,
	)

	err := suite.repo.Create(ctx, owner)
	suite.Require().NoError(err)

	err = suite.repo.Delete(ctx, owner.ID)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByID(ctx, owner.ID)
	assert.Error(suite.T(), err)
}

func TestOwnerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(OwnerRepositoryTestSuite))
}