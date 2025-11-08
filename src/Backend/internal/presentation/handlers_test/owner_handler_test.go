package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOwnerHandler_CreateOwner_InvalidCPF(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	// Test with nil use case to focus on validation
	handler := handlers.NewOwnerHandler(nil)
	
	reqBody := handlers.CreateOwnerRequest{
		Name:     "João Silva",
		Email:    "joao@email.com",
		Password: "senha123",
		Phone:    "11999999999",
		CPF:      "11111111111", // CPF inválido
	}
	
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/owners", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	handler.CreateOwner(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid CPF", response["error"])
}

func TestOwnerHandler_CreateOwner_MissingRequiredFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewOwnerHandler(nil)
	
	reqBody := handlers.CreateOwnerRequest{
		Name: "João Silva",
		// Missing required fields
	}
	
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/owners", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	handler.CreateOwner(c)
	
	// Should fail validation due to missing required fields
	assert.Equal(t, http.StatusBadRequest, w.Code)
}