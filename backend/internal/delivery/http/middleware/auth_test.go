// Package middleware_test testa os middlewares HTTP.
package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/delivery/http/middleware"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── Helpers ───────────────────────────────────────────────────────────────────

// newMiddlewareRouter cria um router com o middleware Auth aplicado
// e uma rota protegida que devolve os valores injetados no contexto.
func newMiddlewareRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/protected", middleware.Auth(), func(c *gin.Context) {
		claims := c.MustGet("user").(*jwt.Claims)
		c.JSON(http.StatusOK, gin.H{
			"user_id": claims.UserID,
			"email":   claims.Email,
			"role":    claims.Role,
		})
	})

	return r
}

// doMiddlewareRequest dispara uma requisição e retorna o recorder.
func doMiddlewareRequest(r *gin.Engine, cookie *http.Cookie) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	if cookie != nil {
		req.AddCookie(cookie)
	}
	r.ServeHTTP(w, req)
	return w
}

// validToken gera um access token JWT real para usar nos testes.
// Ajuste os parâmetros conforme a assinatura de jwt.GenerateAccessToken do seu pkg.
func validToken(t *testing.T) string {
	t.Helper()
	token, err := jwt.GenerateAccessToken("user-id-123", "joao@email.com", "tenant")
	require.NoError(t, err, "falha ao gerar token de teste")
	return token
}

// ── Testes ────────────────────────────────────────────────────────────────────

func TestAuthMiddleware(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "test-access-secret-key-with-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "test-refresh-secret-key-with-32ch")

	t.Run("token válido — injeta claims e passa para o handler", func(t *testing.T) {
		r := newMiddlewareRouter()
		token := validToken(t)

		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "access_token",
			Value: token,
		})

		assert.Equal(t, http.StatusOK, w.Code)

		// Valida que os valores foram injetados corretamente no contexto
		assert.Contains(t, w.Body.String(), "user-id-123")
		assert.Contains(t, w.Body.String(), "joao@email.com")
		assert.Contains(t, w.Body.String(), "tenant")
	})

	t.Run("sem cookie access_token retorna 401 com mensagem correta", func(t *testing.T) {
		r := newMiddlewareRouter()

		// Nenhum cookie enviado
		w := doMiddlewareRequest(r, nil)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "token ausente")
	})

	t.Run("cookie com valor vazio retorna 401", func(t *testing.T) {
		r := newMiddlewareRouter()

		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "access_token",
			Value: "",
		})

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("token com assinatura inválida retorna 401", func(t *testing.T) {
		r := newMiddlewareRouter()

		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "access_token",
			Value: "header.payload.assinatura-invalida",
		})

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "token inválido")
	})

	t.Run("token malformado retorna 401", func(t *testing.T) {
		r := newMiddlewareRouter()

		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "access_token",
			Value: "isso-nao-e-um-jwt",
		})

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "token inválido")
	})

	t.Run("token expirado retorna 401", func(t *testing.T) {
		r := newMiddlewareRouter()

		// Token já expirado — gerado com MaxAge negativo ou fixo no passado.
		// Substitua por jwt.GenerateAccessTokenWithExpiry se seu pkg suportar,
		// ou use um token expirado hardcoded de testes anteriores.
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJ1c2VySUQiOiJ4eHgiLCJleHAiOjF9." +
			"assinatura-qualquer"

		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "access_token",
			Value: expiredToken,
		})

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "token inválido")
	})

	t.Run("requisição abortada — handler não é executado", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		handlerExecuted := false

		r := gin.New()
		r.GET("/protected", middleware.Auth(), func(c *gin.Context) {
			handlerExecuted = true
			c.Status(http.StatusOK)
		})

		// Sem cookie → middleware deve abortar antes do handler
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.False(t, handlerExecuted, "o handler não deveria ter sido executado")
	})

	t.Run("cookie com nome errado é ignorado — retorna 401", func(t *testing.T) {
		r := newMiddlewareRouter()
		token := validToken(t)

		// Envia o token no cookie errado
		w := doMiddlewareRequest(r, &http.Cookie{
			Name:  "token", // nome errado, deveria ser "access_token"
			Value: token,
		})

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "token ausente")
	})
}
