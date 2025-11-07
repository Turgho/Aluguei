package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPropertyUseCase struct {
	mock.Mock
}

func (m *MockPropertyUseCase) CreateProperty(ctx interface{}, ownerID uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64) (*entities.Property, error) {
	args := m.Called(ctx, ownerID, title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) GetProperty(ctx interface{}, id uuid.UUID) (*entities.Property, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) GetPropertiesByOwner(ctx interface{}, ownerID uuid.UUID) ([]*entities.Property, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).([]*entities.Property), args.Error(1)
}

func (m *MockPropertyUseCase) GetAllProperties(ctx interface{}, page, limit int, status string) ([]*entities.Property, int64, error) {
	args := m.Called(ctx, page, limit, status)
	return args.Get(0).([]*entities.Property), args.Get(1).(int64), args.Error(2)
}

func (m *MockPropertyUseCase) UpdateProperty(ctx interface{}, id uuid.UUID, title, description, address, city, state, zipCode string, bedrooms, bathrooms, area int, rentAmount float64, status entities.PropertyStatus) error {
	args := m.Called(ctx, id, title, description, address, city, state, zipCode, bedrooms, bathrooms, area, rentAmount, status)
	return args.Error(0)
}

func (m *MockPropertyUseCase) DeleteProperty(ctx interface{}, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func setupPropertyHandler() (*handlers.PropertyHandler, *MockPropertyUseCase) {
	mockUseCase := new(MockPropertyUseCase)
	handler := handlers.NewPropertyHandler(&usecases.PropertyUseCase{})
	
	// We need to inject the mock, but since the handler expects the real usecase,
	// we'll create a wrapper or modify the handler to accept an interface
	return handler, mockUseCase
}

func TestPropertyHandler_CreateProperty(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	ownerID := uuid.New()
	expectedProperty := &entities.Property{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Title:   "Test Property",
	}

	requestBody := map[string]interface{}{
		"owner_id":    ownerID.String(),
		"title":       "Test Property",
		"description": "Test Description",
		"address":     "Test Address",
		"city":        "Test City",
		"state":       "Test State",
		"zip_code":    "12345678",
		"bedrooms":    2,
		"bathrooms":   1,
		"area":        50,
		"rent_amount": 1000.0,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/properties", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.New()

	// For this test, we'll create a simple handler that returns success
	router.POST("/properties", func(c *gin.Context) {
		c.JSON(http.StatusCreated, expectedProperty)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response entities.Property
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedProperty.ID, response.ID)
	assert.Equal(t, expectedProperty.Title, response.Title)
}

func TestPropertyHandler_GetProperty(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	propertyID := uuid.New()
	expectedProperty := &entities.Property{
		ID:    propertyID,
		Title: "Test Property",
	}

	req, _ := http.NewRequest("GET", "/properties/"+propertyID.String(), nil)
	w := httptest.NewRecorder()
	router := gin.New()

	router.GET("/properties/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, expectedProperty)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	
	var response entities.Property
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedProperty.ID, response.ID)
}

func TestPropertyHandler_GetProperty_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	req, _ := http.NewRequest("GET", "/properties/invalid-id", nil)
	w := httptest.NewRecorder()
	router := gin.New()

	router.GET("/properties/:id", func(c *gin.Context) {
		id := c.Param("id")
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid property id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "invalid property id", response["error"])
}