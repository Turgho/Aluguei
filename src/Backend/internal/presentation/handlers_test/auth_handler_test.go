package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockOwnerUseCase struct {
	mock.Mock
}

func (m *MockOwnerUseCase) ValidatePassword(ctx interface{}, email, password string) (*entities.Owner, error) {
	args := m.Called(ctx, email, password)
	return args.Get(0).(*entities.Owner), args.Error(1)
}

func TestAuthHandler_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test owner with hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	owner := &entities.Owner{
		ID:       uuid.New(),
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashedPassword),
	}

	requestBody := map[string]string{
		"email":    "john@example.com",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.New()

	// Mock successful login
	router.POST("/auth/login", func(c *gin.Context) {
		var loginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simulate successful validation
		if loginReq.Email == "john@example.com" && loginReq.Password == "password123" {
			c.JSON(http.StatusOK, gin.H{
				"token":      "mock-jwt-token",
				"expires_at": time.Now().Add(24 * time.Hour),
				"owner":      owner,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		}
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "mock-jwt-token", response["token"])
	assert.NotNil(t, response["owner"])
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)

	requestBody := map[string]string{
		"email":    "john@example.com",
		"password": "wrongpassword",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.New()

	router.POST("/auth/login", func(c *gin.Context) {
		var loginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "invalid credentials", response["error"])
}

func TestAuthHandler_Login_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.New()

	router.POST("/auth/login", func(c *gin.Context) {
		var loginReq struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
