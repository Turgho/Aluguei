// internal/infra/repositories/user_repository.go
package repositories

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository retorna uma implementação de [domain.UserRepository].
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByCPF(cpf string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("cpf = ?", cpf).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}

// Search realiza busca textual por nome, email ou CPF.
func (r *userRepository) Search(query string) ([]*entities.User, error) {
	var users []*entities.User

	pattern := "%" + query + "%"
	if err := r.db.
		Where(`
			first_name ILIKE ? OR
			last_name  ILIKE ? OR
			email      ILIKE ? OR
			cpf        ILIKE ?
		`, pattern, pattern, pattern, pattern).
		Order("first_name ASC").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
