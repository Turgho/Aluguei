// pkg/response/response_test.go
package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newTestRouter(handler gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/test", handler)
	return r
}

func TestError(t *testing.T) {
	t.Run("retorna status e corpo corretos", func(t *testing.T) {
		r := newTestRouter(func(c *gin.Context) {
			response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var res response.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)
		assert.Equal(t, "USER_NOT_FOUND", res.Code)
		assert.Equal(t, "usuário não encontrado", res.Error)
		assert.Nil(t, res.Details)
	})

	t.Run("retorna status 400 em bad request", func(t *testing.T) {
		r := newTestRouter(func(c *gin.Context) {
			response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "requisição inválida")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var res response.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &res)
		assert.Equal(t, "BAD_REQUEST", res.Code)
	})

	t.Run("retorna status 401 em unauthorized", func(t *testing.T) {
		r := newTestRouter(func(c *gin.Context) {
			response.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "credenciais inválidas")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestValidationError(t *testing.T) {
	t.Run("retorna detalhes de validação", func(t *testing.T) {
		r := newTestRouter(func(c *gin.Context) {
			response.ValidationError(c, map[string]string{
				"email": "email inválido",
				"cpf":   "CPF inválido",
			})
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var res response.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &res)
		assert.Equal(t, "VALIDATION_ERROR", res.Code)
		assert.Equal(t, "erro de validação", res.Error)
		assert.Equal(t, "email inválido", res.Details["email"])
		assert.Equal(t, "CPF inválido", res.Details["cpf"])
	})

	t.Run("retorna details vazio quando não há erros", func(t *testing.T) {
		r := newTestRouter(func(c *gin.Context) {
			response.ValidationError(c, map[string]string{})
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var res response.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &res)
		assert.Equal(t, "VALIDATION_ERROR", res.Code)
	})
}
