// internal/delivery/http/handlers/user_handler_test.go
package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/delivery/http/handlers"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/pkg/pagination"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── Mock ──────────────────────────────────────────────────────────────────────

type mockUserUseCase struct {
	createFn       func(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error)
	getByIDFn      func(id uuid.UUID) (*entities.User, error)
	getByEmailFn   func(email string) (*entities.User, error)
	getByCPFFn     func(cpf string) (*entities.User, error)
	updateFn       func(user *entities.User) error
	deleteFn       func(id uuid.UUID) error
	searchFn       func(filters repositories.UserFilters) (pagination.Result[*entities.User], error)
	activateUserFn func(id uuid.UUID) error
	deactivateFn   func(id uuid.UUID) error
	verifyEmailFn  func(id uuid.UUID) error
	recordLoginFn  func(id uuid.UUID) error

	// auth
	loginFn        func(email, password string) (string, string, error)
	refreshTokenFn func(refreshToken string) (string, error)
}

func (m *mockUserUseCase) Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
	return m.createFn(firstName, lastName, cpf, email, phone, password, role)
}
func (m *mockUserUseCase) GetByID(id uuid.UUID) (*entities.User, error) {
	return m.getByIDFn(id)
}
func (m *mockUserUseCase) GetByEmail(email string) (*entities.User, error) {
	return m.getByEmailFn(email)
}
func (m *mockUserUseCase) GetByCPF(cpf string) (*entities.User, error) {
	return m.getByCPFFn(cpf)
}
func (m *mockUserUseCase) Update(user *entities.User) error {
	return m.updateFn(user)
}
func (m *mockUserUseCase) Delete(id uuid.UUID) error {
	return m.deleteFn(id)
}
func (m *mockUserUseCase) Search(filters repositories.UserFilters) (pagination.Result[*entities.User], error) {
	return m.searchFn(filters)
}
func (m *mockUserUseCase) ActivateUser(id uuid.UUID) error   { return m.activateUserFn(id) }
func (m *mockUserUseCase) DeactivateUser(id uuid.UUID) error { return m.deactivateFn(id) }
func (m *mockUserUseCase) VerifyEmail(id uuid.UUID) error    { return m.verifyEmailFn(id) }
func (m *mockUserUseCase) RecordLogin(id uuid.UUID) error    { return m.recordLoginFn(id) }

// ── Helpers ───────────────────────────────────────────────────────────────────

func fakeUser() *entities.User {
	return &entities.User{
		ID:           uuid.New(),
		FirstName:    "João",
		LastName:     "Silva",
		CPF:          "000.000.000-00",
		Email:        "joao@email.com",
		Phone:        "11999999999",
		PasswordHash: "hash123",
		Role:         entities.RoleTenant,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func newTestRouter(h *handlers.UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/v1/users/search", h.Search)
	r.GET("/api/v1/users/:id", h.GetByID)
	r.PUT("/api/v1/users/:id", h.Update)
	r.DELETE("/api/v1/users/:id", h.Delete)
	return r
}

func doRequest(r *gin.Engine, method, url string, body []byte) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var req *http.Request
	if body != nil {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	r.ServeHTTP(w, req)
	return w
}

func fakePagedResult(users []*entities.User) pagination.Result[*entities.User] {
	return pagination.New(users, int64(len(users)), 1, 20)
}

// ── GetByID ───────────────────────────────────────────────────────────────────

func TestGetByID(t *testing.T) {
	t.Run("retorna usuário com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/"+user.ID.String(), nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, user.ID.String(), res["id"])
		assert.Equal(t, user.FirstName, res["first_name"])
		assert.Equal(t, user.LastName, res["last_name"])
		assert.Equal(t, user.Email, res["email"])
	})

	t.Run("ID inválido retorna 400", func(t *testing.T) {
		r := newTestRouter(handlers.NewUserHandler(&mockUserUseCase{}))
		w := doRequest(r, http.MethodGet, "/api/v1/users/nao-e-um-uuid", nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_ID", res["code"])
	})

	t.Run("usuário não encontrado retorna 404", func(t *testing.T) {
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) {
				return nil, errors.New("not found")
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/"+uuid.New().String(), nil)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "USER_NOT_FOUND", res["code"])
	})

	t.Run("não expõe password_hash na resposta", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/"+user.ID.String(), nil)

		assert.NotContains(t, w.Body.String(), "password_hash")
		assert.NotContains(t, w.Body.String(), "hash123")
	})

	t.Run("repassa o UUID correto para o usecase", func(t *testing.T) {
		user := fakeUser()
		var capturedID uuid.UUID
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) {
				capturedID = id
				return user, nil
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		doRequest(r, http.MethodGet, "/api/v1/users/"+user.ID.String(), nil)

		assert.Equal(t, user.ID, capturedID)
	})
}

// ── Update ────────────────────────────────────────────────────────────────────

func TestUpdate(t *testing.T) {
	t.Run("atualiza first_name com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
			updateFn:  func(u *entities.User) error { return nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{"first_name": "Pedro"})
		w := doRequest(r, http.MethodPut, "/api/v1/users/"+user.ID.String(), body)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "Pedro", res["first_name"])
	})

	t.Run("atualiza múltiplos campos com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
			updateFn:  func(u *entities.User) error { return nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"first_name": "Carlos",
			"last_name":  "Mendes",
			"phone":      "11888888888",
		})
		w := doRequest(r, http.MethodPut, "/api/v1/users/"+user.ID.String(), body)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "Carlos", res["first_name"])
		assert.Equal(t, "Mendes", res["last_name"])
		assert.Equal(t, "11888888888", res["phone"])
	})

	t.Run("body vazio não altera campos existentes", func(t *testing.T) {
		user := fakeUser()
		originalFirst := user.FirstName
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
			updateFn:  func(u *entities.User) error { return nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{})
		w := doRequest(r, http.MethodPut, "/api/v1/users/"+user.ID.String(), body)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, originalFirst, res["first_name"])
	})

	t.Run("ID inválido retorna 400", func(t *testing.T) {
		r := newTestRouter(handlers.NewUserHandler(&mockUserUseCase{}))

		body, _ := json.Marshal(map[string]any{"first_name": "Test"})
		w := doRequest(r, http.MethodPut, "/api/v1/users/nao-e-um-uuid", body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("usuário não encontrado retorna 404", func(t *testing.T) {
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) {
				return nil, errors.New("not found")
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{"first_name": "Test"})
		w := doRequest(r, http.MethodPut, "/api/v1/users/"+uuid.New().String(), body)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("body inválido retorna 400", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		w := doRequest(r, http.MethodPut, "/api/v1/users/"+user.ID.String(), []byte(`{invalid json}`))

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("erro interno no update retorna 500", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id uuid.UUID) (*entities.User, error) { return user, nil },
			updateFn:  func(u *entities.User) error { return errors.New("db error") },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{"first_name": "Test"})
		w := doRequest(r, http.MethodPut, "/api/v1/users/"+user.ID.String(), body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

// ── Delete ────────────────────────────────────────────────────────────────────

func TestDelete(t *testing.T) {
	t.Run("deleta com sucesso retorna 204 sem body", func(t *testing.T) {
		uc := &mockUserUseCase{
			deleteFn: func(id uuid.UUID) error { return nil },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodDelete, "/api/v1/users/"+uuid.New().String(), nil)

		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Empty(t, w.Body.String())
	})

	t.Run("ID inválido retorna 400", func(t *testing.T) {
		r := newTestRouter(handlers.NewUserHandler(&mockUserUseCase{}))
		w := doRequest(r, http.MethodDelete, "/api/v1/users/nao-e-um-uuid", nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("usuário não encontrado retorna 404", func(t *testing.T) {
		uc := &mockUserUseCase{
			deleteFn: func(id uuid.UUID) error { return errors.New("not found") },
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodDelete, "/api/v1/users/"+uuid.New().String(), nil)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "USER_NOT_FOUND", res["code"])
	})
}

// ── Search ────────────────────────────────────────────────────────────────────

func TestSearch(t *testing.T) {
	t.Run("retorna resultado paginado com sucesso", func(t *testing.T) {
		users := []*entities.User{fakeUser(), fakeUser()}
		uc := &mockUserUseCase{
			searchFn: func(f repositories.UserFilters) (pagination.Result[*entities.User], error) {
				return fakePagedResult(users), nil
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/search?q=joão", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, float64(2), res["total"])
		assert.Len(t, res["data"], 2)
	})

	t.Run("retorna resultado vazio quando não encontra", func(t *testing.T) {
		uc := &mockUserUseCase{
			searchFn: func(f repositories.UserFilters) (pagination.Result[*entities.User], error) {
				return fakePagedResult([]*entities.User{}), nil
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/search?q=naoexiste", nil)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]any
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, float64(0), res["total"])
	})

	t.Run("aplica defaults de paginação", func(t *testing.T) {
		var capturedFilters repositories.UserFilters
		uc := &mockUserUseCase{
			searchFn: func(f repositories.UserFilters) (pagination.Result[*entities.User], error) {
				capturedFilters = f
				return fakePagedResult([]*entities.User{}), nil
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		doRequest(r, http.MethodGet, "/api/v1/users/search", nil)

		assert.Equal(t, 1, capturedFilters.Page)
		assert.Equal(t, 20, capturedFilters.PageSize)
	})

	t.Run("repassa o termo de busca corretamente", func(t *testing.T) {
		var capturedFilters repositories.UserFilters
		uc := &mockUserUseCase{
			searchFn: func(f repositories.UserFilters) (pagination.Result[*entities.User], error) {
				capturedFilters = f
				return fakePagedResult([]*entities.User{}), nil
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		doRequest(r, http.MethodGet, "/api/v1/users/search?search=maria", nil)

		assert.Equal(t, "maria", capturedFilters.Search)
	})

	t.Run("erro interno retorna 500", func(t *testing.T) {
		uc := &mockUserUseCase{
			searchFn: func(f repositories.UserFilters) (pagination.Result[*entities.User], error) {
				return pagination.Result[*entities.User]{}, errors.New("db error")
			},
		}
		r := newTestRouter(handlers.NewUserHandler(uc))
		w := doRequest(r, http.MethodGet, "/api/v1/users/search?q=joão", nil)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
