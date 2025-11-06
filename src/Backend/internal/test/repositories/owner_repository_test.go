// tests/repositories/owner_repository_test.go
package repositories

import "github.com/Turgho/Aluguei/internal/test/fixtures"

func (suite *RepositoriesTestSuite) TestOwnerRepository_Create_Success() {
	newOwner := fixtures.CreateTestOwner()
	newOwner.Email = "novo@test.com" // Email único

	err := suite.ownerRepo.Create(newOwner)

	suite.NoError(err)
	suite.NotEmpty(newOwner.ID)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Create_DuplicateEmail() {
	duplicateOwner := fixtures.CreateTestOwner()
	duplicateOwner.Email = suite.testOwner.Email // Email duplicado

	err := suite.ownerRepo.Create(duplicateOwner)

	suite.Error(err)
	suite.Equal("ALREADY_EXISTS", err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByID_Success() {
	owner, err := suite.ownerRepo.FindByID(suite.testOwner.ID.String())

	suite.NoError(err)
	suite.NotNil(owner)
	suite.Equal(suite.testOwner.Name, owner.Name)
	suite.Equal(suite.testOwner.Email, owner.Email)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByID_NotFound() {
	owner, err := suite.ownerRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.Error(err)
	suite.Nil(owner)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByEmail_Success() {
	owner, err := suite.ownerRepo.FindByEmail(suite.testOwner.Email)

	suite.NoError(err)
	suite.NotNil(owner)
	suite.Equal(suite.testOwner.ID, owner.ID)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByEmail_NotFound() {
	owner, err := suite.ownerRepo.FindByEmail("naoexiste@test.com")

	suite.Error(err)
	suite.Nil(owner)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Update_Success() {
	suite.testOwner.Name = "Nome Atualizado"
	suite.testOwner.Phone = "11999998888"

	err := suite.ownerRepo.Update(suite.testOwner)
	suite.NoError(err)

	// Verify update
	updatedOwner, err := suite.ownerRepo.FindByID(suite.testOwner.ID.String())
	suite.NoError(err)
	suite.Equal("Nome Atualizado", updatedOwner.Name)
	suite.Equal("11999998888", updatedOwner.Phone)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Delete_Success() {
	newOwner := fixtures.CreateTestOwner()
	newOwner.Email = "deletar@test.com"
	err := suite.ownerRepo.Create(newOwner)
	suite.NoError(err)

	err = suite.ownerRepo.Delete(newOwner.ID.String())
	suite.NoError(err)

	// Verify deletion
	deletedOwner, err := suite.ownerRepo.FindByID(newOwner.ID.String())
	suite.Error(err)
	suite.Nil(deletedOwner)
}
