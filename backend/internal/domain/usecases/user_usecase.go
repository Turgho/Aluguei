// internal/domain/usecases/user/user_usecase.go
package usecases

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/Turgho/Aluguei/pkg/pagination"
	"github.com/google/uuid"
)

// UserUseCase define o contrato das regras de negócio de [entities.User].
type UserUseCase interface {
	// Escrita
	Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID) error

	// Leitura — lookups por chave única
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByCPF(cpf string) (*entities.User, error)

	// Busca administrativa
	Search(filters repositories.UserFilters) (pagination.Result[*entities.User], error)

	// Operações de conta
	ActivateUser(id uuid.UUID) error
	DeactivateUser(id uuid.UUID) error
	VerifyEmail(id uuid.UUID) error
	RecordLogin(id uuid.UUID) error
}
