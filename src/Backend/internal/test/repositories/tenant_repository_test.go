package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestTenantRepository_Create_Success() {
	owner := suite.createUniqueTestOwner()
	newTenant := fixtures.CreateTestTenant(owner.ID)
	newTenant.Email = suite.generateUniqueEmail() // Use helper para email único
	newTenant.CPF = suite.generateUniqueCPF()     // Use helper para CPF único

	err := suite.tenantRepo.Create(newTenant)

	suite.Nil(err) // ← CORREÇÃO: Use Nil para *errors.AppError
	suite.NotEmpty(newTenant.ID)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByID_Success() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)

	foundTenant, err := suite.tenantRepo.FindByID(tenant.ID.String())

	suite.Nil(err) // ← CORREÇÃO
	suite.NotNil(foundTenant)
	suite.Equal(tenant.Name, foundTenant.Name)
	suite.Equal(tenant.Email, foundTenant.Email)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByID_NotFound() {
	tenant, err := suite.tenantRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.NotNil(err) // ← CORREÇÃO
	suite.Nil(tenant)
	suite.Equal(errors.ErrorCodeNotFound, err.Code) // ← CORREÇÃO: Use a constante
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Success() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)

	// VERIFIQUE se este método existe no seu repositório
	tenants, err := suite.tenantRepo.FindByOwnerID(owner.ID.String())

	suite.Nil(err)
	suite.Len(tenants, 1)
	suite.Equal(tenant.ID, tenants[0].ID)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Multiple() {
	owner := suite.createUniqueTestOwner()

	tenant1 := suite.createUniqueTestTenant(owner.ID)
	tenant2 := suite.createUniqueTestTenant(owner.ID)

	tenants, err := suite.tenantRepo.FindByOwnerID(owner.ID.String())

	suite.Nil(err)
	suite.Len(tenants, 2)

	tenantIDs := make(map[string]bool)
	for _, t := range tenants {
		tenantIDs[t.ID.String()] = true
	}

	suite.True(tenantIDs[tenant1.ID.String()])
	suite.True(tenantIDs[tenant2.ID.String()])
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Empty() {
	newOwner := suite.createUniqueTestOwner()

	tenants, err := suite.tenantRepo.FindByOwnerID(newOwner.ID.String())

	suite.Nil(err)
	suite.Len(tenants, 0)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindAll_Success() {
	owner1 := suite.createUniqueTestOwner()
	owner2 := suite.createUniqueTestOwner()

	tenant1 := suite.createUniqueTestTenant(owner1.ID)
	tenant2 := suite.createUniqueTestTenant(owner2.ID)

	tenants, err := suite.tenantRepo.FindAll()

	suite.Nil(err)

	tenantIDs := make(map[string]bool)
	for _, t := range tenants {
		tenantIDs[t.ID.String()] = true
	}

	suite.True(tenantIDs[tenant1.ID.String()])
	suite.True(tenantIDs[tenant2.ID.String()])
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Update_Success() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)

	tenant.Name = "Nome Atualizado"
	tenant.Phone = "11999997777"

	err := suite.tenantRepo.Update(tenant)
	suite.Nil(err)

	updatedTenant, err := suite.tenantRepo.FindByID(tenant.ID.String())
	suite.Nil(err)
	suite.Equal("Nome Atualizado", updatedTenant.Name)
	suite.Equal("11999997777", updatedTenant.Phone)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Update_Email() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)

	tenant.Email = suite.generateUniqueEmail() // Use email único

	err := suite.tenantRepo.Update(tenant)
	suite.Nil(err)

	updatedTenant, err := suite.tenantRepo.FindByID(tenant.ID.String())
	suite.Nil(err)
	suite.Equal(tenant.Email, updatedTenant.Email)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Delete_Success() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)

	err := suite.tenantRepo.Delete(tenant.ID.String())
	suite.Nil(err)

	deletedTenant, err := suite.tenantRepo.FindByID(tenant.ID.String())
	suite.NotNil(err)
	suite.Nil(deletedTenant)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Delete_WithActiveContracts() {
	owner := suite.createUniqueTestOwner()
	tenant := suite.createUniqueTestTenant(owner.ID)
	property := suite.createUniqueTestProperty(owner.ID)

	// Criar contrato ativo
	contract := fixtures.CreateTestContract(property.ID, tenant.ID)
	err := suite.contractRepo.Create(contract)
	suite.Nil(err)

	// Não deve conseguir deletar tenant com contratos ativos
	err = suite.tenantRepo.Delete(tenant.ID.String())

	// CORREÇÃO: Verifique se err não é nil antes de acessar .Code
	suite.NotNil(err, "Deveria retornar erro ao tentar deletar tenant com contratos ativos")
	if err != nil {
		suite.Equal(errors.ErrorCodeBusinessRule, err.Code)
	}
}
