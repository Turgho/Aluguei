// tests/repositories/contract_repository_test.go
package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestContractRepository_Create_Success() {
	// Criar dados únicos para este teste
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)

	newContract := fixtures.CreateTestContract(property.ID, tenant.ID)

	err := suite.contractRepo.Create(newContract)

	suite.Nil(err) // ← CORREÇÃO
	suite.NotEmpty(newContract.ID)
	suite.Equal(models.ContractStatusActive, newContract.Status)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	foundContract, err := suite.contractRepo.FindByID(contract.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.NotNil(foundContract)
	suite.Equal(contract.MonthlyRent, foundContract.MonthlyRent)
	suite.NotNil(foundContract.Property)
	suite.NotNil(foundContract.Tenant)
	suite.NotNil(foundContract.Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByID_NotFound() {
	contract, err := suite.contractRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.NotNil(err) // ← CORREÇÃO
	suite.Nil(contract)
	suite.Equal(errors.ErrorCodeNotFound, err.Code) // ← CORREÇÃO
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByPropertyID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	contracts, err := suite.contractRepo.FindByPropertyID(property.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.Len(contracts, 1)
	suite.Equal(contract.ID, contracts[0].ID)
	suite.NotNil(contracts[0].Property)
	suite.NotNil(contracts[0].Tenant)
	suite.NotNil(contracts[0].Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByPropertyID_Empty() {
	// Criar property sem contratos
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	contracts, err := suite.contractRepo.FindByPropertyID(property.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.Len(contracts, 0)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindByTenantID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	contracts, err := suite.contractRepo.FindByTenantID(tenant.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.Len(contracts, 1)
	suite.Equal(contract.ID, contracts[0].ID)
	suite.NotNil(contracts[0].Property)
	suite.NotNil(contracts[0].Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindActiveByPropertyID_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	foundContract, err := suite.contractRepo.FindActiveByPropertyID(property.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.NotNil(foundContract)
	suite.Equal(contract.ID, foundContract.ID)
	suite.Equal(models.ContractStatusActive, foundContract.Status)
	suite.NotNil(foundContract.Property)
	suite.NotNil(foundContract.Tenant)
	suite.NotNil(foundContract.Payments)
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindActiveByPropertyID_NotFound() {
	// Criar property sem contrato ativo
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)

	contract, err := suite.contractRepo.FindActiveByPropertyID(property.ID.String())

	suite.NotNil(err) // ← CORREÇÃO
	suite.Nil(contract)
	suite.Equal(errors.ErrorCodeNotFound, err.Code) // ← CORREÇÃO
}

func (suite *RepositoriesTestSuite) TestContractRepository_FindAll_Success() {
	// Criar múltiplos contratos
	owner1 := suite.createUniqueTestOwner()
	owner2 := suite.createUniqueTestOwner()

	property1 := suite.createUniqueTestProperty(owner1.ID)
	property2 := suite.createUniqueTestProperty(owner2.ID)

	tenant1 := suite.createUniqueTestTenant(owner1.ID)
	tenant2 := suite.createUniqueTestTenant(owner2.ID)

	contract1 := suite.createTestContract(property1.ID, tenant1.ID)
	contract2 := suite.createTestContract(property2.ID, tenant2.ID)

	contracts, err := suite.contractRepo.FindAll()

	suite.Nil(err) // ← CORREÇÃO

	// Verificar se os contratos criados estão na lista
	contractIDs := make(map[string]bool)
	for _, c := range contracts {
		contractIDs[c.ID.String()] = true
	}

	suite.True(contractIDs[contract1.ID.String()])
	suite.True(contractIDs[contract2.ID.String()])
}

func (suite *RepositoriesTestSuite) TestContractRepository_Update_Success() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	contract.MonthlyRent = 1600.00
	contract.PaymentDueDay = 10

	err := suite.contractRepo.Update(contract)
	suite.Nil(err) // ← CORREÇÃO

	updatedContract, err := suite.contractRepo.FindByID(contract.ID.String())
	suite.Nil(err) // ← CORREÇÃO
	suite.Equal(1600.00, updatedContract.MonthlyRent)
	suite.Equal(10, updatedContract.PaymentDueDay)
}

func (suite *RepositoriesTestSuite) TestContractRepository_Update_Status() {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	contract.Status = models.ContractStatusExpired

	err := suite.contractRepo.Update(contract)
	suite.Nil(err) // ← CORREÇÃO

	updatedContract, err := suite.contractRepo.FindByID(contract.ID.String())
	suite.Nil(err) // ← CORREÇÃO
	suite.Equal(models.ContractStatusExpired, updatedContract.Status)
}

func (suite *RepositoriesTestSuite) TestContractRepository_Delete_Success() {
	// Criar contrato específico para deletar
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)
	contract := suite.createTestContract(property.ID, tenant.ID)

	err := suite.contractRepo.Delete(contract.ID.String())
	suite.Nil(err) // ← CORREÇÃO

	// Verificar se foi deletado
	deletedContract, err := suite.contractRepo.FindByID(contract.ID.String())
	suite.NotNil(err) // ← CORREÇÃO
	suite.Nil(deletedContract)
	suite.Equal(errors.ErrorCodeNotFound, err.Code) // ← CORREÇÃO
}
