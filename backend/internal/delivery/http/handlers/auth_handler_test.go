// internal/delivery/http/handlers/auth_handler_test.go
package handlers_test

// Esse arquivo está no mesmo package que user_handler_test.go,
// então mockUserUseCase, fakeUser e doRequest já estão disponíveis.
// Aqui só adicionamos os métodos Login/RefreshToken no mock
// e os testes específicos do AuthHandler.

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/delivery/http/handlers"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── Métodos do mock que só o AuthHandler usa ──────────────────────────────────
// (mockUserUseCase está definido em user_handler_test.go)

func (m *mockUserUseCase) Login(email, password string) (string, string, error) {
	return m.loginFn(email, password)
}

func (m *mockUserUseCase) RefreshToken(refreshToken string) (string, error) {
	return m.refreshTokenFn(refreshToken)
}

// ── Router de auth ────────────────────────────────────────────────────────────

func newAuthTestRouter(h *handlers.AuthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/v1/auth/register", h.Register)
	r.POST("/api/v1/auth/login", h.Login)
	r.POST("/api/v1/auth/refresh", h.RefreshToken)
	r.POST("/api/v1/auth/logout", h.Logout)

	// Rota /me precisa de middleware que injeta jwt.Claims no contexto
	r.GET("/api/v1/auth/me", fakeJWTMiddleware("valid-user-id"), h.Me)
	return r
}

// fakeJWTMiddleware simula o middleware de autenticação real
// injetando jwt.Claims no contexto para testar o handler Me.
func fakeJWTMiddleware(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", &jwt.Claims{UserID: userID})
		c.Next()
	}
}

// ── Register ──────────────────────────────────────────────────────────────────

func TestRegister(t *testing.T) {
	validBody := map[string]any{
		"first_name": "João",
		"last_name":  "Silva",
		"cpf":        "000.000.000-00",
		"email":      "joao@email.com",
		"phone":      "11999999999",
		"password":   "senha123",
		"role":       entities.RoleTenant,
	}

	t.Run("registra usuário com sucesso e retorna 201", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			createFn: func(_, _, _, _, _, _ string, _ entities.Role) (*entities.User, error) {
				return user, nil
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(validBody)
		w := doRequest(r, http.MethodPost, "/api/v1/auth/register", body)

		assert.Equal(t, http.StatusCreated, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, user.Email, res["email"])
		assert.NotContains(t, w.Body.String(), "password_hash")
	})

	t.Run("body inválido retorna 400", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		w := doRequest(r, http.MethodPost, "/api/v1/auth/register", []byte(`{invalid}`))

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("campos obrigatórios ausentes retornam 400", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))

		// sem email
		body, _ := json.Marshal(map[string]any{
			"first_name": "João",
			"last_name":  "Silva",
			"cpf":        "000.000.000-00",
			"password":   "senha123",
			"role":       entities.RoleTenant,
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/register", body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("senha menor que 8 caracteres retorna 400", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"first_name": "João",
			"last_name":  "Silva",
			"cpf":        "000.000.000-00",
			"email":      "joao@email.com",
			"password":   "123",
			"role":       entities.RoleTenant,
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/register", body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("e-mail já cadastrado retorna 409", func(t *testing.T) {
		uc := &mockUserUseCase{
			createFn: func(_, _, _, _, _, _ string, _ entities.Role) (*entities.User, error) {
				return nil, errors.New("email already exists")
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(validBody)
		w := doRequest(r, http.MethodPost, "/api/v1/auth/register", body)

		assert.Equal(t, http.StatusConflict, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "CONFLICT", res["code"])
	})
}

// ── Login ─────────────────────────────────────────────────────────────────────

func TestLogin(t *testing.T) {
	t.Run("login com sucesso seta cookies e retorna success:true", func(t *testing.T) {
		uc := &mockUserUseCase{
			loginFn: func(email, password string) (string, string, error) {
				return "access-token-abc", "refresh-token-xyz", nil
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(map[string]any{
			"email":    "joao@email.com",
			"password": "senha123",
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/login", body)

		assert.Equal(t, http.StatusOK, w.Code)

		// Valida body
		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, true, res["success"])

		// Valida cookies
		cookies := w.Result().Cookies()
		var accessCookie, refreshCookie *http.Cookie
		for _, c := range cookies {
			if c.Name == "access_token" {
				accessCookie = c
			}
			if c.Name == "refresh_token" {
				refreshCookie = c
			}
		}
		require.NotNil(t, accessCookie, "cookie access_token não encontrado")
		require.NotNil(t, refreshCookie, "cookie refresh_token não encontrado")
		assert.Equal(t, "access-token-abc", accessCookie.Value)
		assert.Equal(t, "refresh-token-xyz", refreshCookie.Value)
		assert.True(t, accessCookie.HttpOnly)
		assert.True(t, refreshCookie.HttpOnly)
	})

	t.Run("credenciais inválidas retornam 401", func(t *testing.T) {
		uc := &mockUserUseCase{
			loginFn: func(email, password string) (string, string, error) {
				return "", "", errors.New("invalid credentials")
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(map[string]any{
			"email":    "joao@email.com",
			"password": "errada",
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/login", body)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_CREDENTIALS", res["code"])
	})

	t.Run("body inválido retorna 400", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		w := doRequest(r, http.MethodPost, "/api/v1/auth/login", []byte(`{invalid}`))

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("e-mail inválido retorna 400", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(map[string]any{
			"email":    "nao-e-email",
			"password": "senha123",
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/login", body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("não expõe tokens no body da resposta", func(t *testing.T) {
		uc := &mockUserUseCase{
			loginFn: func(_, _ string) (string, string, error) {
				return "access-token-abc", "refresh-token-xyz", nil
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		body, _ := json.Marshal(map[string]any{
			"email":    "joao@email.com",
			"password": "senha123",
		})
		w := doRequest(r, http.MethodPost, "/api/v1/auth/login", body)

		assert.NotContains(t, w.Body.String(), "access-token-abc")
		assert.NotContains(t, w.Body.String(), "refresh-token-xyz")
	})
}

// ── RefreshToken ──────────────────────────────────────────────────────────────

func TestRefreshToken(t *testing.T) {
	t.Run("renova access token com sucesso", func(t *testing.T) {
		uc := &mockUserUseCase{
			refreshTokenFn: func(token string) (string, error) {
				return "new-access-token", nil
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))

		// Simula cookie refresh_token enviado pelo browser
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)
		req.AddCookie(&http.Cookie{Name: "refresh_token", Value: "valid-refresh"})
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, true, res["success"])

		// Novo access_token deve estar no cookie
		cookies := w.Result().Cookies()
		var accessCookie *http.Cookie
		for _, c := range cookies {
			if c.Name == "access_token" {
				accessCookie = c
			}
		}
		require.NotNil(t, accessCookie)
		assert.Equal(t, "new-access-token", accessCookie.Value)
		assert.True(t, accessCookie.HttpOnly)
	})

	t.Run("sem cookie refresh_token retorna 401", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))

		// Requisição sem cookie
		w := doRequest(r, http.MethodPost, "/api/v1/auth/refresh", nil)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "MISSING_REFRESH_TOKEN", res["code"])
	})

	t.Run("refresh token inválido retorna 401", func(t *testing.T) {
		uc := &mockUserUseCase{
			refreshTokenFn: func(token string) (string, error) {
				return "", errors.New("token expirado")
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)
		req.AddCookie(&http.Cookie{Name: "refresh_token", Value: "expired-token"})
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_REFRESH_TOKEN", res["code"])
	})
}

// ── Logout ────────────────────────────────────────────────────────────────────

func TestLogout(t *testing.T) {
	t.Run("logout limpa os cookies e retorna success:true", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		w := doRequest(r, http.MethodPost, "/api/v1/auth/logout", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, true, res["success"])

		// Cookies devem ter MaxAge=-1 (instrução de deletar)
		cookies := w.Result().Cookies()
		for _, c := range cookies {
			if c.Name == "access_token" || c.Name == "refresh_token" {
				assert.Equal(t, -1, c.MaxAge, "cookie %s deveria ter MaxAge=-1", c.Name)
				assert.Empty(t, c.Value)
			}
		}
	})
}

// ── Me ────────────────────────────────────────────────────────────────────────

func TestMe(t *testing.T) {
	t.Run("retorna usuário autenticado com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				assert.Equal(t, "valid-user-id", id)
				return user, nil
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/auth/me", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, user.Email, res["email"])
		assert.NotContains(t, w.Body.String(), "password_hash")
	})

	t.Run("usuário do token não encontrado retorna 404", func(t *testing.T) {
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				return nil, errors.New("not found")
			},
		}
		r := newAuthTestRouter(handlers.NewAuthHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/auth/me", nil)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "USER_NOT_FOUND", res["code"])
	})
}
