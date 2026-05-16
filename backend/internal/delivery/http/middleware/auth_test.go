// internal/delivery/http/middleware/auth_test.go
package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/delivery/http/middleware"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newAuthRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/protected", middleware.Auth(), func(c *gin.Context) {
		userID, _ := c.Get("userID")
		email, _ := c.Get("email")
		role, _ := c.Get("role")
		c.JSON(http.StatusOK, gin.H{
			"userID": userID,
			"email":  email,
			"role":   role,
		})
	})

	return r
}

func TestAuth(t *testing.T) {
	t.Run("token válido", func(t *testing.T) {
		token, err := jwt.GenerateToken("123", "joao@email.com", "tenant")
		assert.NoError(t, err)

		r := newAuthRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("token ausente", func(t *testing.T) {
		r := newAuthRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("token inválido", func(t *testing.T) {
		r := newAuthRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer token-invalido")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("token sem Bearer", func(t *testing.T) {
		token, err := jwt.GenerateToken("123", "joao@email.com", "tenant")
		assert.NoError(t, err)

		r := newAuthRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", token) // sem "Bearer "
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("claims injetadas no contexto", func(t *testing.T) {
		token, err := jwt.GenerateToken("123", "joao@email.com", "tenant")
		assert.NoError(t, err)

		r := newAuthRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "joao@email.com")
		assert.Contains(t, w.Body.String(), "tenant")
	})
}
