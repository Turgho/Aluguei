package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTenantHandler_CreateTenant_InvalidCPF(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewTenantHandler(nil)
	
	ownerID := uuid.New()
	reqBody := handlers.CreateTenantRequest{
		Name:    "Maria Silva",
		Email:   "maria@email.com",
		Phone:   "11888888888",
		CPF:     "11111111111", // CPF inválido
		OwnerID: ownerID.String(),
	}
	
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/tenants", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	handler.CreateTenant(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid CPF", response["error"])
}

func TestTenantHandler_CreateTenant_InvalidOwnerID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewTenantHandler(nil)
	
	reqBody := handlers.CreateTenantRequest{
		Name:    "Maria Silva",
		Email:   "maria@email.com",
		Phone:   "11888888888",
		CPF:     "11144477735",
		OwnerID: "invalid-uuid",
	}
	
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/tenants", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	handler.CreateTenant(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid owner_id", response["error"])
}