// tests/repositories/contract_repository_test.go
package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestContractRepository_Create_Success() {
	// Criar novo tenant e property para evitar conflito
	newTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	newTenant.Email = "novotenant@test.com"
	err := suite.tenantRepo.Create(newTenant)
	suite.NoError(err)

	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Nova Prop para Contrato"
	err = suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	newContract := fixtures.CreateTestContract(newProperty.ID, newTenant.ID)

	err = suite.contractRepo.Create(newContract)

	suite.NoError(err)
	suite.NotEmpty(newContract.ID)
	suite.Equal(models.ContractStatusActive, newContract.Status)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByID_Success() {
	contract, err := suite.contractRepo.FindByID(suite.testContract.ID.String())

	suite.NoError(err)
	suite.NotNil(contract)
	suite.Equal(suite.testContract.MonthlyRent, contract.MonthlyRent)
	suite.NotNil(contract.Property)
	suite.NotNil(contract.Tenant)
	suite.NotNil(contract.Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByID_NotFound() {
	contract, err := suite.contractRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.Error(err)
	suite.Nil(contract)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByPropertyID_Success() {
	contracts, err := suite.contractRepo.FindByPropertyID(suite.testProperty.ID.String())

	suite.NoError(err)
	suite.Len(contracts, 1)
	suite.Equal(suite.testContract.ID, contracts[0].ID)
	suite.NotNil(contracts[0].Property)
	suite.NotNil(contracts[0].Tenant)
	suite.NotNil(contracts[0].Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByPropertyID_Empty() {
	// Criar property sem contratos
	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Property Sem Contratos"
	err := suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	contracts, err := suite.contractRepo.FindByPropertyID(newProperty.ID.String())

	suite.NoError(err)
	suite.Len(contracts, 0)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByTenantID_Success() {
	contracts, err := suite.contractRepo.FindByTenantID(suite.testTenant.ID.String())

	suite.NoError(err)
	suite.Len(contracts, 1)
	suite.Equal(suite.testContract.ID, contracts[0].ID)
	suite.NotNil(contracts[0].Property)
	suite.NotNil(contracts[0].Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindActiveByPropertyID_Success() {
	contract, err := suite.contractRepo.FindActiveByPropertyID(suite.testProperty.ID.String())

	suite.NoError(err)
	suite.NotNil(contract)
	suite.Equal(suite.testContract.ID, contract.ID)
	suite.Equal(models.ContractStatusActive, contract.Status)
	suite.NotNil(contract.Property)
	suite.NotNil(contract.Tenant)
	suite.NotNil(contract.Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindActiveByPropertyID_NotFound() {
	// Criar property sem contrato ativo
	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Property Sem Contrato Ativo"
	err := suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	contract, err := suite.contractRepo.FindActiveByPropertyID(newProperty.ID.String())

	suite.Error(err)
	suite.Nil(contract)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindAll_Success() {
	contracts, err := suite.contractRepo.FindAll()

	suite.NoError(err)
	suite.Len(contracts, 1)
	suite.Equal(suite.testContract.ID, contracts[0].ID)
}

func (suite *RepositoriesTestSuite) TestContractRepository_Update_Success() {
	suite.testContract.MonthlyRent = 1600.00
	suite.testContract.PaymentDueDay = 10

	err := suite.contractRepo.Update(suite.testContract)
	suite.NoError(err)

	updatedContract, err := suite.contractRepo.FindByID(suite.testContract.ID.String())
	suite.NoError(err)
	suite.Equal(1600.00, updatedContract.MonthlyRent)
	suite.Equal(10, updatedContract.PaymentDueDay)
}

func (suite *RepositoriesTestSuite) TestContractRepository_Update_Status() {
	suite.testContract.Status = models.ContractStatusExpired

	err := suite.contractRepo.Update(suite.testContract)
	suite.NoError(err)

	updatedContract, err := suite.contractRepo.FindByID(suite.testContract.ID.String())
	suite.NoError(err)
	suite.Equal(models.ContractStatusExpired, updatedContract.Status)
}

func (suite *RepositoriesTestSuite) TestContractRepository_Delete_Success() {
	// Criar contrato específico para deletar
	newTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	newTenant.Email = "deletar@test.com"
	err := suite.tenantRepo.Create(newTenant)
	suite.NoError(err)

	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Property Para Deletar Contrato"
	err = suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	newContract := fixtures.CreateTestContract(newProperty.ID, newTenant.ID)
	err = suite.contractRepo.Create(newContract)
	suite.NoError(err)

	err = suite.contractRepo.Delete(newContract.ID.String())
	suite.NoError(err)

	// Verificar se foi deletado
	deletedContract, err := suite.contractRepo.FindByID(newContract.ID.String())
	suite.Error(err)
	suite.Nil(deletedContract)
}
