// internal/domain/repositories/user_repository.go
package repositories

import "github.com/Turgho/Aluguei/internal/domain/entities"

// UserRepository define o contrato de acesso a dados de [entities.User].
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByCPF(cpf string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
	Search(query string) ([]*entities.User, error) // raw para buscas
}
