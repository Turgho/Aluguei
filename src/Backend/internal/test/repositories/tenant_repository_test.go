// tests/repositories/tenant_repository_test.go
package repositories

import "github.com/Turgho/Aluguei/internal/test/fixtures"

func (suite *RepositoriesTestSuite) TestTenantRepository_Create_Success() {
	newTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	newTenant.Email = "novo@tenant.com"
	newTenant.CPF = "98765432100"

	err := suite.tenantRepo.Create(newTenant)

	suite.NoError(err)
	suite.NotEmpty(newTenant.ID)
	suite.Equal("novo@tenant.com", newTenant.Email)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByID_Success() {
	tenant, err := suite.tenantRepo.FindByID(suite.testTenant.ID.String())

	suite.NoError(err)
	suite.NotNil(tenant)
	suite.Equal(suite.testTenant.Name, tenant.Name)
	suite.Equal(suite.testTenant.Email, tenant.Email)
	suite.Len(tenant.Contracts, 1) // Deve ter o contrato criado no setup
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByID_NotFound() {
	tenant, err := suite.tenantRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.Error(err)
	suite.Nil(tenant)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Success() {
	tenants, err := suite.tenantRepo.FindByOwnerID(suite.testOwner.ID.String())

	suite.NoError(err)
	suite.Len(tenants, 1)
	suite.Equal(suite.testTenant.ID, tenants[0].ID)
	suite.Len(tenants[0].Contracts, 1)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Multiple() {
	// Criar segundo tenant
	secondTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	secondTenant.Email = "segundo@tenant.com"
	secondTenant.CPF = "11122233344"
	err := suite.tenantRepo.Create(secondTenant)
	suite.NoError(err)

	tenants, err := suite.tenantRepo.FindByOwnerID(suite.testOwner.ID.String())

	suite.NoError(err)
	suite.Len(tenants, 2)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindByOwnerID_Empty() {
	// Criar novo owner sem tenants
	newOwner := fixtures.CreateTestOwner()
	newOwner.Email = "owner.sem.tenants@test.com"
	err := suite.ownerRepo.Create(newOwner)
	suite.NoError(err)

	tenants, err := suite.tenantRepo.FindByOwnerID(newOwner.ID.String())

	suite.NoError(err)
	suite.Len(tenants, 0)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_FindAll_Success() {
	tenants, err := suite.tenantRepo.FindAll()

	suite.NoError(err)
	suite.Len(tenants, 1)
	suite.Equal(suite.testTenant.ID, tenants[0].ID)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Update_Success() {
	suite.testTenant.Name = "Maria Santos Atualizada"
	suite.testTenant.Phone = "11999997777"

	err := suite.tenantRepo.Update(suite.testTenant)
	suite.NoError(err)

	updatedTenant, err := suite.tenantRepo.FindByID(suite.testTenant.ID.String())
	suite.NoError(err)
	suite.Equal("Maria Santos Atualizada", updatedTenant.Name)
	suite.Equal("11999997777", updatedTenant.Phone)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Update_Email() {
	suite.testTenant.Email = "maria.atualizada@test.com"

	err := suite.tenantRepo.Update(suite.testTenant)
	suite.NoError(err)

	updatedTenant, err := suite.tenantRepo.FindByID(suite.testTenant.ID.String())
	suite.NoError(err)
	suite.Equal("maria.atualizada@test.com", updatedTenant.Email)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Delete_Success() {
	newTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	newTenant.Email = "deletar@tenant.com"
	err := suite.tenantRepo.Create(newTenant)
	suite.NoError(err)

	err = suite.tenantRepo.Delete(newTenant.ID.String())
	suite.NoError(err)

	// Verificar se foi deletado
	deletedTenant, err := suite.tenantRepo.FindByID(newTenant.ID.String())
	suite.Error(err)
	suite.Nil(deletedTenant)
}

func (suite *RepositoriesTestSuite) TestTenantRepository_Delete_WithActiveContracts() {
	// Não deve conseguir deletar tenant com contratos ativos
	err := suite.tenantRepo.Delete(suite.testTenant.ID.String())

	suite.Error(err)
	suite.Equal("BUSINESS_RULE", err.Code)
}
