package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestPropertyRepository_Create_Success() {
	owner := suite.createUniqueTestOwner()
	newProperty := fixtures.CreateTestProperty(owner.ID)
	newProperty.Title = "Nova Propriedade Única"

	err := suite.propertyRepo.Create(newProperty)

	suite.Nil(err)
	suite.NotEmpty(newProperty.ID)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	foundProperty, err := suite.propertyRepo.FindByID(property.ID.String())

	suite.Nil(err)
	suite.NotNil(foundProperty)
	suite.Equal(property.Title, foundProperty.Title)
	suite.Len(foundProperty.Contracts, 0) // Não tem contratos ainda
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByOwnerID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	properties, err := suite.propertyRepo.FindByOwnerID(owner.ID.String())

	suite.Nil(err)
	suite.Len(properties, 1)
	suite.Equal(property.ID, properties[0].ID)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByOwnerID_Multiple() {
	owner := suite.createUniqueTestOwner()

	// Criar múltiplas propriedades
	property1 := suite.createUniqueTestProperty(owner.ID)
	property2 := suite.createUniqueTestProperty(owner.ID)

	properties, err := suite.propertyRepo.FindByOwnerID(owner.ID.String())

	suite.Nil(err)
	suite.Len(properties, 2)

	// Verificar se ambas as propriedades estão na lista
	propertyIDs := make(map[string]bool)
	for _, p := range properties {
		propertyIDs[p.ID.String()] = true
	}

	suite.True(propertyIDs[property1.ID.String()])
	suite.True(propertyIDs[property2.ID.String()])
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindByOwnerID_Empty() {
	// Criar novo owner sem propriedades
	newOwner := suite.createUniqueTestOwner()

	properties, err := suite.propertyRepo.FindByOwnerID(newOwner.ID.String())

	suite.Nil(err)
	suite.Len(properties, 0)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_FindAll_Success() {
	// Criar propriedades com diferentes owners
	owner1 := suite.createUniqueTestOwner()
	owner2 := suite.createUniqueTestOwner()

	property1 := suite.createUniqueTestProperty(owner1.ID)
	property2 := suite.createUniqueTestProperty(owner2.ID)

	properties, err := suite.propertyRepo.FindAll()

	suite.Nil(err)

	// Verificar se as propriedades criadas estão na lista
	propertyIDs := make(map[string]bool)
	for _, p := range properties {
		propertyIDs[p.ID.String()] = true
	}

	suite.True(propertyIDs[property1.ID.String()])
	suite.True(propertyIDs[property2.ID.String()])
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_Update_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	property.RentAmount = 1600.00
	property.Status = models.PropertyStatusRented

	err := suite.propertyRepo.Update(property)

	suite.Nil(err)

	updatedProperty, err := suite.propertyRepo.FindByID(property.ID.String())

	suite.Nil(err)
	suite.Equal(1600.00, updatedProperty.RentAmount)
	suite.Equal(models.PropertyStatusRented, updatedProperty.Status)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_Delete_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	err := suite.propertyRepo.Delete(property.ID.String())
	suite.Nil(err)

	// Verify deletion
	deletedProperty, err := suite.propertyRepo.FindByID(property.ID.String())

	suite.NotNil(err)
	suite.Nil(deletedProperty)

	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}

func (suite *RepositoriesTestSuite) TestPropertyRepository_Delete_WithActiveContracts() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)

	// Criar contrato ativo
	contract := fixtures.CreateTestContract(property.ID, tenant.ID)
	err := suite.contractRepo.Create(contract)
	suite.Nil(err)

	// Não deve conseguir deletar propriedade com contratos ativos
	err = suite.propertyRepo.Delete(property.ID.String())

	suite.NotNil(err)
	if err != nil {
		suite.Equal(errors.ErrorCodeBusinessRule, err.Code)
	}
}
