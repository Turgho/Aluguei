// tests/repositories/property_repository_test.go
package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestPropertyRepository_Create_Success() {
	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Nova Propriedade"

	err := suite.propertyRepo.Create(newProperty)

	suite.NoError(err)
	suite.NotEmpty(newProperty.ID)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByID_Success() {
	property, err := suite.propertyRepo.FindByID(suite.testProperty.ID.String())

	suite.NoError(err)
	suite.NotNil(property)
	suite.Equal(suite.testProperty.Title, property.Title)
	suite.Len(property.Contracts, 1) // Should have the contract from setup
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByOwnerID_Success() {
	properties, err := suite.propertyRepo.FindByOwnerID(suite.testOwner.ID.String())

	suite.NoError(err)
	suite.Len(properties, 1)
	suite.Equal(suite.testProperty.ID, properties[0].ID)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByOwnerID_Empty() {
	// Create new owner without properties
	newOwner := fixtures.CreateTestOwner()
	newOwner.Email = "novoowner@test.com"
	err := suite.ownerRepo.Create(newOwner)
	suite.NoError(err)

	properties, err := suite.propertyRepo.FindByOwnerID(newOwner.ID.String())

	suite.NoError(err)
	suite.Len(properties, 0)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindAll_Success() {
	// Create second property
	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Segunda Propriedade"
	err := suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	properties, err := suite.propertyRepo.FindAll()

	suite.NoError(err)
	suite.Len(properties, 2)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_Update_Success() {
	suite.testProperty.RentAmount = 1600.00
	suite.testProperty.Status = models.PropertyStatusRented

	err := suite.propertyRepo.Update(suite.testProperty)
	suite.NoError(err)

	updatedProperty, err := suite.propertyRepo.FindByID(suite.testProperty.ID.String())
	suite.NoError(err)
	suite.Equal(1600.00, updatedProperty.RentAmount)
	suite.Equal(models.PropertyStatusRented, updatedProperty.Status)
}
