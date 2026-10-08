// internal/domain/usecases/user/user_usecase.go
package usecases

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
)

// UserUseCase define o contrato das regras de negócio de [entities.User].
type UserUseCase interface {
	Create(firstName, lastName, cpf, email, phone, password string, role entities.Role) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
	Search(query string) ([]*entities.User, error)
	Login(email, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, error)
}
