package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/infrastructure/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestJWT(ownerID, email, secret string) (string, error) {
	claims := jwt.MapClaims{
		"owner_id": ownerID,
		"email":    email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtSecret := "test-secret"
	ownerID := uuid.New().String()
	email := "test@example.com"

	token, err := createTestJWT(ownerID, email, jwtSecret)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	authMiddleware(c)

	assert.False(t, c.IsAborted())
	assert.Equal(t, ownerID, c.GetString("owner_id"))
	assert.Equal(t, email, c.GetString("email"))
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtSecret := "test-secret"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	authMiddleware(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}