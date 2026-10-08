// internal/infra/repositories/user_repository.go
package repositories

import (
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const DefaultPageSize = 20

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository retorna uma implementação de [domain.UserRepository].
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// ── Escrita ──────────────────────────────────────────────────────────────────

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}

// ── Leitura por chave única ───────────────────────────────────────────────────

func (r *userRepository) GetByID(id uuid.UUID) (*entities.User, error) {
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

// ── Busca com filtros ─────────────────────────────────────────────────────────

func (r *userRepository) Search(filters domain.UserFilters) ([]*entities.User, int64, error) {
	var users []*entities.User
	var total int64

	q := r.db.Model(&entities.User{})

	if filters.Role != nil {
		q = q.Where("role = ?", *filters.Role)
	}
	if filters.IsActive != nil {
		q = q.Where("is_active = ?", *filters.IsActive)
	}
	if filters.EmailVerified != nil {
		q = q.Where("email_verified = ?", *filters.EmailVerified)
	}
	if filters.Search != "" {
		pattern := "%" + filters.Search + "%"
		q = q.Where(
			"first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ? OR cpf ILIKE ?",
			pattern, pattern, pattern, pattern,
		)
	}

	// Sessão separada para o count
	if err := q.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Defaults de paginação
	pageSize := filters.PageSize
	if pageSize <= 0 {
		pageSize = DefaultPageSize // valor padrão
	}
	page := filters.Page
	if page <= 0 {
		page = 1
	}

	if err := q.
		Order("first_name ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ── Operações de conta ────────────────────────────────────────────────────────

func (r *userRepository) SetActive(id uuid.UUID, active bool) error {
	return r.db.Model(&entities.User{}).
		Where("id = ?", id).
		Update("is_active", active).Error
}

func (r *userRepository) SetEmailVerified(id uuid.UUID) error {
	return r.db.Model(&entities.User{}).
		Where("id = ?", id).
		Update("email_verified", true).Error
}

func (r *userRepository) UpdateLastLogin(id uuid.UUID, at time.Time) error {
	return r.db.Model(&entities.User{}).
		Where("id = ?", id).
		Update("last_login_at", at).Error
}
