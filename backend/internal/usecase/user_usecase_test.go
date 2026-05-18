// internal/domain/usecases/user/user_usecase_test.go
package usecase_test

import (
	"errors"
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/usecase"
	jwtutil "github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mockUserRepository struct {
	createFn     func(user *entities.User) error
	getByEmailFn func(email string) (*entities.User, error)
	getByCPFFn   func(cpf string) (*entities.User, error)
	getByIDFn    func(id string) (*entities.User, error)
}

func (m *mockUserRepository) Create(user *entities.User) error {
	return m.createFn(user)
}

func (m *mockUserRepository) GetByEmail(email string) (*entities.User, error) {
	return m.getByEmailFn(email)
}

func (m *mockUserRepository) GetByCPF(cpf string) (*entities.User, error) {
	return m.getByCPFFn(cpf)
}

// mocks vazios só pra satisfazer interface
func (m *mockUserRepository) GetByID(id string) (*entities.User, error) {
	if m.getByIDFn != nil {
		return m.getByIDFn(id)
	}
	return nil, nil
}
func (m *mockUserRepository) Update(user *entities.User) error {
	return nil
}

func (m *mockUserRepository) Delete(id string) error {
	return nil
}

func (m *mockUserRepository) Search(query string) ([]*entities.User, error) {
	return nil, nil
}

var _ repositories.UserRepository = (*mockUserRepository)(nil)

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
				createFn: func(user *entities.User) error {
					return nil
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
				createFn: func(user *entities.User) error {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "erro ao salvar",
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

func TestRefreshTokenUseCase(t *testing.T) {
	t.Setenv("JWT_REFRESH_SECRET", "segredo-refresh-para-testes-minimo-32-chars")
	t.Setenv("JWT_ACCESS_SECRET", "segredo-access-para-testes-minimo-32-chars")

	user := &entities.User{
		ID:    uuid.New(),
		Email: "joao@email.com",
		Role:  entities.RoleTenant,
	}

	validToken, err := jwtutil.GenerateRefreshToken(user.ID.String())
	if err != nil {
		t.Fatalf("erro ao gerar refresh token: %v", err)
	}

	tests := []struct {
		name    string
		token   string
		repo    repositories.UserRepository
		wantErr bool
	}{
		{
			name:  "refresh com sucesso",
			token: validToken,
			repo: &mockUserRepository{
				getByIDFn: func(id string) (*entities.User, error) {
					return user, nil
				},
			},
			wantErr: false,
		},
		{
			name:    "token inválido",
			token:   "token.invalido.qualquer",
			repo:    &mockUserRepository{},
			wantErr: true,
		},
		{
			name:  "usuário não encontrado",
			token: validToken,
			repo: &mockUserRepository{
				getByIDFn: func(id string) (*entities.User, error) {
					return nil, gorm.ErrRecordNotFound
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewUserUseCase(tt.repo)
			newAccessToken, err := uc.RefreshToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("esperado erro=%v, recebido=%v", tt.wantErr, err)
			}
			if !tt.wantErr && newAccessToken == "" {
				t.Error("esperado access token não vazio")
			}
		})
	}
}
