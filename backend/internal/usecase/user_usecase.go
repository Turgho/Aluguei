// internal/domain/usecases/user_usecase.go
package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/hash"
	"github.com/Turgho/Aluguei/pkg/pagination"
	userValidators "github.com/Turgho/Aluguei/pkg/validators/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userUseCase struct {
	repo domain.UserRepository
}

// NewUserUseCase retorna uma implementação de [usecases.UserUseCase].
func NewUserUseCase(repo domain.UserRepository) usecases.UserUseCase {
	return &userUseCase{repo: repo}
}

// ── Escrita ──────────────────────────────────────────────────────────────────

// Create valida, hasheia a senha e persiste um novo usuário.
func (uc *userUseCase) Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error) {
	// Verifica força da senha antes do hash
	if !userValidators.ValidatePassword(password) {
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

// Update atualiza os dados de um usuário existente.
func (uc *userUseCase) Update(user *entities.User) error {
	if err := uc.repo.Update(user); err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

// Delete remove um usuário pelo ID.
func (uc *userUseCase) Delete(id uuid.UUID) error {
	if err := uc.repo.Delete(id); err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	return nil
}

// ── Leitura por chave única ───────────────────────────────────────────────────

// GetByID busca um usuário pelo ID.
func (uc *userUseCase) GetByID(id uuid.UUID) (*entities.User, error) {
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

// GetByCPF busca um usuário pelo CPF.
func (uc *userUseCase) GetByCPF(cpf string) (*entities.User, error) {
	user, err := uc.repo.GetByCPF(cpf)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

// ── Busca com filtros ─────────────────────────────────────────────────────────

// Search realiza busca textual por nome, email ou CPF.
func (uc *userUseCase) Search(filter domain.UserFilters) (pagination.Result[*entities.User], error) {
	users, total, err := uc.repo.Search(filter)
	if err != nil {
		return pagination.Result[*entities.User]{}, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	return pagination.New(users, total, filter.Page, filter.PageSize), nil
}

// ── Operações de conta ────────────────────────────────────────────────────────

// ActivateUser atualiza o status para ATIVO da conta de um usuário.
func (uc *userUseCase) ActivateUser(id uuid.UUID) error {
	if err := uc.repo.SetActive(id, true); err != nil {
		return fmt.Errorf("erro ao ativar conta do usuário: %w", err)
	}
	return nil
}

// DeactivateUser atualiza o status para INATIVO da conta de um usuário.
func (uc *userUseCase) DeactivateUser(id uuid.UUID) error {
	if err := uc.repo.SetActive(id, false); err != nil {
		return fmt.Errorf("erro ao desativar conta do usuário: %w", err)
	}
	return nil
}

// VerifyEmail atualiza o status do email do usuário para VERIFICADO.
func (uc *userUseCase) VerifyEmail(id uuid.UUID) error {
	if err := uc.repo.SetEmailVerified(id); err != nil {
		return fmt.Errorf("erro ao verificar email do usuário: %w", err)
	}
	return nil
}

// RecordLogin atualiza a última data de login de um usuário.
func (uc *userUseCase) RecordLogin(id uuid.UUID) error {
	now := time.Now().UTC()

	if err := uc.repo.UpdateLastLogin(id, now); err != nil {
		return fmt.Errorf("erro ao atualizar horário de login do usuário: %w", err)
	}
	return nil
}
