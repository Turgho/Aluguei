// internal/domain/usecases/user_usecase.go
package usecase

import (
	"errors"
	"fmt"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/hash"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/Turgho/Aluguei/pkg/validators"
	"gorm.io/gorm"
)

type userUseCase struct {
	repo repositories.UserRepository
}

// NewUserUseCase retorna uma implementação de [usecases.UserUseCase].
func NewUserUseCase(repo repositories.UserRepository) usecases.UserUseCase {
	return &userUseCase{repo: repo}
}

// Create valida, hasheia a senha e persiste um novo usuário.
func (uc *userUseCase) Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
	// Verifica força da senha antes do hash
	if !validators.ValidatePassword(password) {
		return nil, fmt.Errorf("senha fraca")
	}

	// Verifica se email já existe
	existing, err := uc.repo.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("erro ao verificar email: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("email já cadastrado")
	}

	// Verifica se CPF já existe
	existingCPF, err := uc.repo.GetByCPF(cpf)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("erro ao verificar CPF: %w", err)
	}
	if existingCPF != nil {
		return nil, fmt.Errorf("CPF já cadastrado")
	}

	// Hasheia a senha
	passwordHash, err := hash.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar senha: %w", err)
	}

	// Cria entidade
	user, err := entities.NewUser(
		firstName,
		lastName,
		cpf,
		email,
		phone,
		passwordHash,
		role,
	)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := uc.repo.Create(user); err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return user, nil
}

// GetByID busca um usuário pelo ID.
func (uc *userUseCase) GetByID(id string) (*entities.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

// GetByEmail busca um usuário pelo email.
func (uc *userUseCase) GetByEmail(email string) (*entities.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

// Update atualiza os dados de um usuário existente.
func (uc *userUseCase) Update(user *entities.User) error {
	if err := uc.repo.Update(user); err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

// Delete remove um usuário pelo ID.
func (uc *userUseCase) Delete(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	return nil
}

// Search realiza busca textual por nome, email ou CPF.
func (uc *userUseCase) Search(query string) ([]*entities.User, error) {
	users, err := uc.repo.Search(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return users, nil
}

// Login autentica um usuário e retorna os tokens JWT.
func (uc *userUseCase) Login(email, password string) (string, string, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", fmt.Errorf("credenciais inválidas")
		}

		return "", "", fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	match, err := hash.VerifyPassword(password, user.PasswordHash)
	if err != nil {
		return "", "", fmt.Errorf("erro ao verificar senha: %w", err)
	}

	if !match {
		return "", "", fmt.Errorf("credenciais inválidas")
	}

	accessToken, err := jwt.GenerateAccessToken(
		user.ID.String(),
		user.Email,
		string(user.Role),
	)
	if err != nil {
		return "", "", fmt.Errorf("erro ao gerar access token: %w", err)
	}

	refreshToken, err := jwt.GenerateRefreshToken(
		user.ID.String(),
	)
	if err != nil {
		return "", "", fmt.Errorf("erro ao gerar refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// RefreshToken valida o refresh token e retorna um novo access token.
func (uc *userUseCase) RefreshToken(refreshToken string) (string, error) {
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("refresh token inválido: %w", err)
	}

	// Se precisar de email e role no access token, busca o usuário
	user, err := uc.repo.GetByID(claims.UserID)
	if err != nil {
		return "", fmt.Errorf("usuário não encontrado: %w", err)
	}

	return jwt.GenerateAccessToken(user.ID.String(), user.Email, string(user.Role))
}
