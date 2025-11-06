// tests/repositories/payment_repository_test.go
package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/models"
	"github.com/Turgho/Aluguei/internal/test/fixtures"
)

func (suite *RepositoriesTestSuite) TestPaymentRepository_Create_Success() {
	newPayment := fixtures.CreateTestPayment(suite.testContract.ID)

	err := suite.paymentRepo.Create(newPayment)

	suite.NoError(err)
	suite.NotEmpty(newPayment.ID)
	suite.Equal(models.PaymentStatusPending, newPayment.Status)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByID_Success() {
	// Primeiro criar um payment
	newPayment := fixtures.CreateTestPayment(suite.testContract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.NoError(err)

	payment, err := suite.paymentRepo.FindByID(newPayment.ID.String())

	suite.NoError(err)
	suite.NotNil(payment)
	suite.Equal(newPayment.Amount, payment.Amount)
	suite.NotNil(payment.Contract)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByID_NotFound() {
	payment, err := suite.paymentRepo.FindByID("00000000-0000-0000-0000-000000000000")

	suite.Error(err)
	suite.Nil(payment)
	suite.Equal("NOT_FOUND", err.Code)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByContractID_Success() {
	// Criar alguns payments para o contrato
	payment1 := fixtures.CreateTestPayment(suite.testContract.ID)
	err := suite.paymentRepo.Create(payment1)
	suite.NoError(err)

	payment2 := fixtures.CreateTestPayment(suite.testContract.ID)
	payment2.DueDate = time.Now().AddDate(0, 2, 0)
	err = suite.paymentRepo.Create(payment2)
	suite.NoError(err)

	payments, err := suite.paymentRepo.FindByContractID(suite.testContract.ID.String())

	suite.NoError(err)
	suite.Len(payments, 2)
	// Deve estar ordenado por due_date DESC
	suite.True(payments[0].DueDate.After(payments[1].DueDate))
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByContractID_Empty() {
	// Criar contrato sem payments
	newTenant := fixtures.CreateTestTenant(suite.testOwner.ID)
	newTenant.Email = "contrato.sem.payments@test.com"
	err := suite.tenantRepo.Create(newTenant)
	suite.NoError(err)

	newProperty := fixtures.CreateTestProperty(suite.testOwner.ID)
	newProperty.Title = "Property Sem Payments"
	err = suite.propertyRepo.Create(newProperty)
	suite.NoError(err)

	newContract := fixtures.CreateTestContract(newProperty.ID, newTenant.ID)
	err = suite.contractRepo.Create(newContract)
	suite.NoError(err)

	payments, err := suite.paymentRepo.FindByContractID(newContract.ID.String())

	suite.NoError(err)
	suite.Len(payments, 0)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindOverdue_Success() {
	// Criar payment em atraso
	overduePayment := fixtures.CreateTestPayment(suite.testContract.ID)
	overduePayment.DueDate = time.Now().AddDate(0, -1, 0) // 1 mês atrás
	overduePayment.Status = models.PaymentStatusOverdue
	err := suite.paymentRepo.Create(overduePayment)
	suite.NoError(err)

	payments, err := suite.paymentRepo.FindOverdue()

	suite.NoError(err)
	suite.Len(payments, 1)
	suite.Equal(overduePayment.ID, payments[0].ID)
	suite.Equal(models.PaymentStatusOverdue, payments[0].Status)
	suite.NotNil(payments[0].Contract)
	suite.NotNil(payments[0].Contract.Property)
	suite.NotNil(payments[0].Contract.Tenant)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_FindByPeriod_Success() {
	startDate := time.Now().AddDate(0, -1, 0)
	endDate := time.Now().AddDate(0, 2, 0)

	// Criar payment dentro do período
	paymentInPeriod := fixtures.CreateTestPayment(suite.testContract.ID)
	paymentInPeriod.DueDate = time.Now().AddDate(0, 1, 0)
	err := suite.paymentRepo.Create(paymentInPeriod)
	suite.NoError(err)

	payments, err := suite.paymentRepo.FindByPeriod(startDate, endDate)

	suite.NoError(err)
	suite.Len(payments, 1)
	suite.Equal(paymentInPeriod.ID, payments[0].ID)
	suite.True(payments[0].DueDate.After(startDate))
	suite.True(payments[0].DueDate.Before(endDate))
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_Update_Success() {
	newPayment := fixtures.CreateTestPayment(suite.testContract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.NoError(err)

	newPayment.PaidAmount = 1500.00
	newPayment.Status = models.PaymentStatusPaid
	paidDate := time.Now()
	newPayment.PaidDate = &paidDate

	err = suite.paymentRepo.Update(newPayment)
	suite.NoError(err)

	updatedPayment, err := suite.paymentRepo.FindByID(newPayment.ID.String())
	suite.NoError(err)
	suite.Equal(1500.00, updatedPayment.PaidAmount)
	suite.Equal(models.PaymentStatusPaid, updatedPayment.Status)
	suite.NotNil(updatedPayment.PaidDate)
}

func (suite *RepositoriesTestSuite) TestPaymentRepository_Delete_Success() {
	newPayment := fixtures.CreateTestPayment(suite.testContract.ID)
	err := suite.paymentRepo.Create(newPayment)
	suite.NoError(err)

	err = suite.paymentRepo.Delete(newPayment.ID.String())
	suite.NoError(err)

	// Verificar se foi deletado
	deletedPayment, err := suite.paymentRepo.FindByID(newPayment.ID.String())
	suite.Error(err)
	suite.Nil(deletedPayment)
}
