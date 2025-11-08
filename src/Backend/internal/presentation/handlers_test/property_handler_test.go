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

func TestPropertyHandler_CreateProperty_InvalidOwnerID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewPropertyHandler(nil)
	
	reqBody := handlers.CreatePropertyRequest{
		OwnerID:    "invalid-uuid",
		Title:      "Casa 3 quartos",
		Address:    "Rua das Flores, 123",
		City:       "São Paulo",
		State:      "SP",
		RentAmount: 2500.00,
	}
	
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/properties", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	handler.CreateProperty(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid owner_id", response["error"])
}

func TestPropertyHandler_GetProperty_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewPropertyHandler(nil)
	
	req := httptest.NewRequest("GET", "/properties/invalid-uuid", nil)
	w := httptest.NewRecorder()
	
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}
	
	handler.GetProperty(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "invalid property id", response["error"])
}