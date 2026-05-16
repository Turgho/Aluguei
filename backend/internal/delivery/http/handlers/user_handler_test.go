// internal/delivery/http/handlers/user_handler_test.go
package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/delivery/http/handlers"
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// mockUserUseCase implementa usecases.UserUseCase para testes.
type mockUserUseCase struct {
	createFn     func(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error)
	getByIDFn    func(id string) (*entities.User, error)
	getByEmailFn func(email string) (*entities.User, error)
	updateFn     func(user *entities.User) error
	deleteFn     func(id string) error
	searchFn     func(query string) ([]*entities.User, error)
	loginFn      func(email, password string) (string, error)
}

func (m *mockUserUseCase) Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
	return m.createFn(firstName, lastName, cpf, email, phone, password, role)
}
func (m *mockUserUseCase) GetByID(id string) (*entities.User, error) {
	return m.getByIDFn(id)
}
func (m *mockUserUseCase) GetByEmail(email string) (*entities.User, error) {
	return m.getByEmailFn(email)
}
func (m *mockUserUseCase) Update(user *entities.User) error {
	return m.updateFn(user)
}
func (m *mockUserUseCase) Delete(id string) error {
	return m.deleteFn(id)
}
func (m *mockUserUseCase) Search(query string) ([]*entities.User, error) {
	return m.searchFn(query)
}
func (m *mockUserUseCase) Login(email, password string) (string, error) {
	return m.loginFn(email, password)
}

// fakeUser retorna um usuário fake para uso nos testes.
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

// newTestRouter cria um router Gin em modo test com o handler registrado.
func newTestRouter(h *handlers.UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.POST("/api/v1/register", h.Register)
	r.POST("/api/v1/login", h.Login)
	r.GET("/api/v1/users/search", h.Search)
	r.GET("/api/v1/users/:id", h.GetByID)
	r.PUT("/api/v1/users/:id", h.Update)
	r.DELETE("/api/v1/users/:id", h.Delete)

	return r
}

func TestRegister(t *testing.T) {
	t.Run("registro com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			createFn: func(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
				return user, nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"first_name": "João",
			"last_name":  "Silva",
			"cpf":        "000.000.000-00",
			"email":      "joao@email.com",
			"phone":      "11999999999",
			"password":   "Senha@123",
			"role":       "tenant",
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("body inválido", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBufferString(`{invalid}`))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("email já cadastrado", func(t *testing.T) {
		uc := &mockUserUseCase{
			createFn: func(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
				return nil, assert.AnError
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"first_name": "João",
			"last_name":  "Silva",
			"cpf":        "000.000.000-00",
			"email":      "joao@email.com",
			"password":   "Senha@123",
			"role":       "tenant",
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login com sucesso", func(t *testing.T) {
		uc := &mockUserUseCase{
			loginFn: func(email, password string) (string, error) {
				return "token-jwt", nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"email":    "joao@email.com",
			"password": "Senha@123",
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]string
		json.Unmarshal(w.Body.Bytes(), &res)
		assert.Equal(t, "token-jwt", res["token"])
	})

	t.Run("credenciais inválidas", func(t *testing.T) {
		uc := &mockUserUseCase{
			loginFn: func(email, password string) (string, error) {
				return "", assert.AnError
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"email":    "joao@email.com",
			"password": "errada",
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("usuário encontrado", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				return user, nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/"+user.ID.String(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("usuário não encontrado", func(t *testing.T) {
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				return nil, assert.AnError
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/"+uuid.New().String(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("atualização com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				return user, nil
			},
			updateFn: func(u *entities.User) error {
				return nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		body, _ := json.Marshal(map[string]any{
			"first_name": "Pedro",
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+user.ID.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("usuário não encontrado", func(t *testing.T) {
		uc := &mockUserUseCase{
			getByIDFn: func(id string) (*entities.User, error) {
				return nil, assert.AnError
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+uuid.New().String(), bytes.NewBufferString(`{}`))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deletado com sucesso", func(t *testing.T) {
		uc := &mockUserUseCase{
			deleteFn: func(id string) error {
				return nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/"+uuid.New().String(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("usuário não encontrado", func(t *testing.T) {
		uc := &mockUserUseCase{
			deleteFn: func(id string) error {
				return assert.AnError
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/"+uuid.New().String(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestSearch(t *testing.T) {
	t.Run("busca com sucesso", func(t *testing.T) {
		user := fakeUser()
		uc := &mockUserUseCase{
			searchFn: func(query string) ([]*entities.User, error) {
				return []*entities.User{user}, nil
			},
		}

		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/search?q=joão", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("parâmetro q ausente", func(t *testing.T) {
		uc := &mockUserUseCase{}
		r := newTestRouter(handlers.NewUserHandler(uc))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/search", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
