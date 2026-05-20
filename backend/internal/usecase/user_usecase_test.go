// internal/domain/usecases/user/user_usecase_test.go
package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/usecase"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// mockUserRepository simula o repositório de usuários
// para testes unitários do use case.
type mockUserRepository struct {
	createFn           func(user *entities.User) error
	getByEmailFn       func(email string) (*entities.User, error)
	getByCPFFn         func(cpf string) (*entities.User, error)
	getByIDFn          func(id uuid.UUID) (*entities.User, error)
	updateFn           func(user *entities.User) error
	deleteFn           func(id uuid.UUID) error
	searchFn           func(filters repositories.UserFilters) ([]*entities.User, int64, error)
	setActiveFn        func(id uuid.UUID, active bool) error
	setEmailVerifiedFn func(id uuid.UUID) error
	updateLastLoginFn  func(id uuid.UUID, at time.Time) error
}

// ── Escrita ───────────────────────────────────────────────────────────────────

func (m *mockUserRepository) Create(user *entities.User) error {
	if m.createFn != nil {
		return m.createFn(user)
	}
	return nil
}

func (m *mockUserRepository) Update(user *entities.User) error {
	if m.updateFn != nil {
		return m.updateFn(user)
	}
	return nil
}

func (m *mockUserRepository) Delete(id uuid.UUID) error {
	if m.deleteFn != nil {
		return m.deleteFn(id)
	}
	return nil
}

// ── Leitura ───────────────────────────────────────────────────────────────────

func (m *mockUserRepository) GetByEmail(email string) (*entities.User, error) {
	if m.getByEmailFn != nil {
		return m.getByEmailFn(email)
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockUserRepository) GetByCPF(cpf string) (*entities.User, error) {
	if m.getByCPFFn != nil {
		return m.getByCPFFn(cpf)
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockUserRepository) GetByID(id uuid.UUID) (*entities.User, error) {
	if m.getByIDFn != nil {
		return m.getByIDFn(id)
	}
	return nil, gorm.ErrRecordNotFound
}

// ── Busca ─────────────────────────────────────────────────────────────────────

func (m *mockUserRepository) Search(filters repositories.UserFilters) ([]*entities.User, int64, error) {
	if m.searchFn != nil {
		return m.searchFn(filters)
	}
	return []*entities.User{}, 0, nil
}

// ── Operações de conta ───────────────────────────────────────────────────────

func (m *mockUserRepository) SetActive(id uuid.UUID, active bool) error {
	if m.setActiveFn != nil {
		return m.setActiveFn(id, active)
	}
	return nil
}

func (m *mockUserRepository) SetEmailVerified(id uuid.UUID) error {
	if m.setEmailVerifiedFn != nil {
		return m.setEmailVerifiedFn(id)
	}

	return nil
}

func (m *mockUserRepository) UpdateLastLogin(id uuid.UUID, at time.Time) error {
	if m.updateLastLoginFn != nil {
		return m.updateLastLoginFn(id, at)
	}

	return nil
}

// Garante em tempo de compilação que o mock implementa a interface.
var _ repositories.UserRepository = (*mockUserRepository)(nil)

// ── Create ────────────────────────────────────────────────────────────────────

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		repo    repositories.UserRepository
		wantErr bool
	}{
		{
			name: "cria usuário com sucesso",
			repo: &mockUserRepository{
				getByEmailFn: func(email string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
				getByCPFFn: func(cpf string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
				createFn: func(user *entities.User) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "email já cadastrado",
			repo: &mockUserRepository{
				getByEmailFn: func(email string) (*entities.User, error) {
					return &entities.User{}, nil
				},
				getByCPFFn: func(cpf string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
			wantErr: true,
		},
		{
			name: "cpf já cadastrado",
			repo: &mockUserRepository{
				getByEmailFn: func(email string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
				getByCPFFn: func(cpf string) (*entities.User, error) {
					return &entities.User{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "erro ao salvar usuário",
			repo: &mockUserRepository{
				getByEmailFn: func(email string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
				getByCPFFn: func(cpf string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
				createFn: func(user *entities.User) error {
					return errors.New("db error")
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewUserUseCase(tt.repo)

			_, err := uc.Create(
				"Victor",
				"Hugo",
				"52998224725",
				"victor@email.com",
				"14999999999",
				"Senha@123",
				entities.RoleOwner,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("esperado erro=%v, recebido=%v", tt.wantErr, err)
			}
		})
	}
}
