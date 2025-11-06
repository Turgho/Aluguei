package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
	"github.com/google/uuid"
)

func (suite *RepositoriesTestSuite) TestOwnerRepository_Create_Success() {
	newOwner := fixtures.CreateTestOwner()
	newOwner.Email = suite.generateUniqueEmail()

	err := suite.ownerRepo.Create(newOwner)

	suite.Nil(err) // Deve ser nil, não *errors.AppError
	suite.NotEmpty(newOwner.ID)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Create_DuplicateEmail() {
	// Primeiro cria um owner
	firstOwner := fixtures.CreateTestOwner()
	firstOwner.Email = "duplicado_especifico@test.com"
	err := suite.ownerRepo.Create(firstOwner)
	suite.Nil(err)

	// Tenta criar outro com mesmo email
	duplicateOwner := fixtures.CreateTestOwner()
	duplicateOwner.Email = "duplicado_especifico@test.com"
	duplicateOwner.ID = uuid.New() // ID diferente

	err = suite.ownerRepo.Create(duplicateOwner)

	suite.NotNil(err) // Deve retornar *errors.AppError
	suite.Equal(errors.ErrorCodeAlreadyExists, err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByID_Success() {
	testOwner := suite.createUniqueTestOwner()

	owner, err := suite.ownerRepo.FindByID(testOwner.ID.String())

	suite.Nil(err)
	suite.NotNil(owner)
	suite.Equal(testOwner.Name, owner.Name)
	suite.Equal(testOwner.Email, owner.Email)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByID_NotFound() {
	owner, err := suite.ownerRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.NotNil(err)
	suite.Nil(owner)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByEmail_Success() {
	testOwner := suite.createUniqueTestOwner()

	owner, err := suite.ownerRepo.FindByEmail(testOwner.Email)

	suite.Nil(err)
	suite.NotNil(owner)
	suite.Equal(testOwner.ID, owner.ID)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_FindByEmail_NotFound() {
	owner, err := suite.ownerRepo.FindByEmail("email_que_nao_existe@test.com")

	suite.NotNil(err)
	suite.Nil(owner)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Update_Success() {
	testOwner := suite.createUniqueTestOwner()

	testOwner.Name = "Nome Atualizado"
	testOwner.Phone = "11999998888"

	err := suite.ownerRepo.Update(testOwner)
	suite.Nil(err)

	// Verify update
	updatedOwner, err := suite.ownerRepo.FindByID(testOwner.ID.String())
	suite.Nil(err)
	suite.Equal("Nome Atualizado", updatedOwner.Name)
	suite.Equal("11999998888", updatedOwner.Phone)
}

func (suite *RepositoriesTestSuite) TestOwnerRepository_Delete_Success() {
	newOwner := suite.createUniqueTestOwner()

	err := suite.ownerRepo.Delete(newOwner.ID.String())
	suite.Nil(err)

	// Verify deletion
	deletedOwner, err := suite.ownerRepo.FindByID(newOwner.ID.String())
	suite.NotNil(err)
	suite.Nil(deletedOwner)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}
