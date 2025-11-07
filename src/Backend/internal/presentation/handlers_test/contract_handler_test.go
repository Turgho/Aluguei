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

func TestContractHandler_CreateContract(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &testhelpers.MockContractUseCase{}
	handler := handlers.NewContractHandler(mockUseCase)

	propertyID := uuid.New()
	tenantID := uuid.New()
	contract := &entities.Contract{
		ID:            uuid.New(),
		PropertyID:    propertyID,
		TenantID:      tenantID,
		MonthlyRent:   1500.00,
		PaymentDueDay: 5,
		Status:        "active",
	}

	mockUseCase.On("CreateContract", mock.Anything, propertyID, tenantID, "2024-01-01", (*string)(nil), 1500.00, 5, "").Return(contract, nil)

	reqBody := map[string]interface{}{
		"property_id":      propertyID,
		"tenant_id":        tenantID,
		"start_date":       "2024-01-01",
		"monthly_rent":     1500.00,
		"payment_due_day":  5,
	}
	jsonBody, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/contracts", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateContract(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUseCase.AssertExpectations(t)
}

func TestContractHandler_GetContractByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &testhelpers.MockContractUseCase{}
	handler := handlers.NewContractHandler(mockUseCase)

	contractID := uuid.New()
	contract := &entities.Contract{
		ID:            contractID,
		MonthlyRent:   1500.00,
		PaymentDueDay: 5,
		Status:        "active",
	}

	mockUseCase.On("GetContractByID", mock.Anything, contractID).Return(contract, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/contracts/"+contractID.String(), nil)
	c.Params = []gin.Param{{Key: "id", Value: contractID.String()}}

	handler.GetContractByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}