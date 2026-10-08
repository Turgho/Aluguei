// internal/domain/repositories/user_repository.go
package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/google/uuid"
)

type UserFilters struct {
	Role          *entities.Role
	IsActive      *bool
	EmailVerified *bool
	Search        string // busca textual em nome/email
	Page          int
	PageSize      int
}

// UserRepository define o contrato de acesso a dados de [entities.User].
type UserRepository interface {
	// Escrita
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id uuid.UUID) error

	// Leitura — lookups por chave única
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByCPF(cpf string) (*entities.User, error)

	// Busca administrativa
	Search(filters UserFilters) ([]*entities.User, int64, error)

	// Operações de conta (evitam UPDATE do objeto inteiro)
	SetActive(id uuid.UUID, active bool) error
	SetEmailVerified(id uuid.UUID) error
	UpdateLastLogin(id uuid.UUID, at time.Time) error
}
