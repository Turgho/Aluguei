package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) createTestContractWithData() *models.Contract {
	owner := suite.createUniqueTestOwner()
	property := suite.createUniqueTestProperty(owner.ID)
	tenant := suite.createUniqueTestTenant(owner.ID)

	contract := fixtures.CreateTestContract(property.ID, tenant.ID)
	err := suite.contractRepo.Create(contract)
	suite.Nil(err) // ← CORREÇÃO
	return contract
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_Create_Success() {
	contract := suite.createTestContractWithData()
	newPayment := fixtures.CreateTestPayment(contract.ID)

	err := suite.paymentRepo.Create(newPayment)

	suite.Nil(err)
	suite.NotEmpty(newPayment.ID)
	suite.Equal(models.PaymentStatusPending, newPayment.Status)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByID_Success() {
	contract := suite.createTestContractWithData()
	newPayment := fixtures.CreateTestPayment(contract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.Nil(err)

	payment, err := suite.paymentRepo.FindByID(newPayment.ID.String())

	suite.Nil(err)
	suite.NotNil(payment)
	suite.Equal(newPayment.Amount, payment.Amount)
	suite.NotNil(payment.Contract)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByID_NotFound() {
	payment, err := suite.paymentRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.NotNil(err)
	suite.Nil(payment)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByContractID_Success() {
	contract := suite.createTestContractWithData()

	// Criar alguns payments para o contrato
	payment1 := fixtures.CreateTestPayment(contract.ID)
	err := suite.paymentRepo.Create(payment1)
	suite.Nil(err)

	payment2 := fixtures.CreateTestPayment(contract.ID)
	payment2.DueDate = time.Now().AddDate(0, 2, 0)
	err = suite.paymentRepo.Create(payment2)
	suite.Nil(err)

	payments, err := suite.paymentRepo.FindByContractID(contract.ID.String())

	suite.Nil(err)
	suite.Len(payments, 2)
	// Deve estar ordenado por due_date DESC
	suite.True(payments[0].DueDate.After(payments[1].DueDate))
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByContractID_Empty() {
	// Criar contrato sem payments
	contract := suite.createTestContractWithData()

	payments, err := suite.paymentRepo.FindByContractID(contract.ID.String())

	suite.Nil(err)
	suite.Len(payments, 0)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindOverdue_Success() {
	contract := suite.createTestContractWithData()

	// Criar payment em atraso
	overduePayment := fixtures.CreateTestPayment(contract.ID)
	overduePayment.DueDate = time.Now().AddDate(0, -1, 0) // 1 mês atrás
	overduePayment.Status = models.PaymentStatusOverdue
	err := suite.paymentRepo.Create(overduePayment)
	suite.Nil(err)

	payments, err := suite.paymentRepo.FindOverdue()

	suite.Nil(err)

	// Verificar se payment está na lista
	found := false
	for _, p := range payments {
		if p.ID == overduePayment.ID {
			found = true
			suite.Equal(models.PaymentStatusOverdue, p.Status)
			suite.NotNil(p.Contract)
			suite.NotNil(p.Contract.Property)
			suite.NotNil(p.Contract.Tenant)
			break
		}
	}

	suite.True(found, "Payment overdue should be found in the list")
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByPeriod_Success() {
	contract := suite.createTestContractWithData()
	startDate := time.Now().AddDate(0, -1, 0)
	endDate := time.Now().AddDate(0, 2, 0)

	// Criar payment dentro do período
	paymentInPeriod := fixtures.CreateTestPayment(contract.ID)
	paymentInPeriod.DueDate = time.Now().AddDate(0, 1, 0)
	err := suite.paymentRepo.Create(paymentInPeriod)
	suite.Nil(err)

	payments, err := suite.paymentRepo.FindByPeriod(startDate, endDate)

	suite.Nil(err)

	// Verificar se payment está na lista
	found := false
	for _, p := range payments {
		if p.ID == paymentInPeriod.ID {
			found = true
			suite.True(p.DueDate.After(startDate))
			suite.True(p.DueDate.Before(endDate))
			break
		}
	}

	suite.True(found, "Payment in period should be found in the list")
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_Update_Success() {
	contract := suite.createTestContractWithData()
	newPayment := fixtures.CreateTestPayment(contract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.Nil(err)

	newPayment.PaidAmount = 1500.00
	newPayment.Status = models.PaymentStatusPaid
	paidDate := time.Now()
	newPayment.PaidDate = &paidDate

	err = suite.paymentRepo.Update(newPayment)
	suite.Nil(err)

	updatedPayment, err := suite.paymentRepo.FindByID(newPayment.ID.String())
	suite.Nil(err)
	suite.Equal(1500.00, updatedPayment.PaidAmount)
	suite.Equal(models.PaymentStatusPaid, updatedPayment.Status)
	suite.NotNil(updatedPayment.PaidDate)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_Delete_Success() {
	contract := suite.createTestContractWithData()
	newPayment := fixtures.CreateTestPayment(contract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.Nil(err)

	err = suite.paymentRepo.Delete(newPayment.ID.String())
	suite.Nil(err)

	// Verificar se foi deletado
	deletedPayment, err := suite.paymentRepo.FindByID(newPayment.ID.String())
	suite.NotNil(err)
	suite.Nil(deletedPayment)
	suite.Equal(errors.ErrorCodeNotFound, err.Code)
}
