package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/Turgho/Aluguei/test/testhelpers"
)

func TestPaymentHandler_CreatePayment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &testhelpers.MockPaymentUseCase{}
	handler := handlers.NewPaymentHandler(mockUseCase)

	contractID := uuid.New()
	payment := &entities.Payment{
		ID:         uuid.New(),
		ContractID: contractID,
		Amount:     1500.00,
		Status:     "pending",
	}

	mockUseCase.On("CreatePayment", mock.Anything, contractID, "2024-01-05", 1500.00, "", "").Return(payment, nil)

	reqBody := map[string]interface{}{
		"contract_id": contractID,
		"due_date":    "2024-01-05",
		"amount":      1500.00,
	}
	jsonBody, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/payments", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreatePayment(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUseCase.AssertExpectations(t)
}

func TestPaymentHandler_GetPaymentByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &testhelpers.MockPaymentUseCase{}
	handler := handlers.NewPaymentHandler(mockUseCase)

	paymentID := uuid.New()
	payment := &entities.Payment{
		ID:     paymentID,
		Amount: 1500.00,
		Status: "pending",
	}

	mockUseCase.On("GetPaymentByID", mock.Anything, paymentID).Return(payment, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/payments/"+paymentID.String(), nil)
	c.Params = []gin.Param{{Key: "id", Value: paymentID.String()}}

	handler.GetPaymentByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}