package repositories

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

// NewPropertyRepository retorna uma implementação de [domain.AddressRepository].
func NewAddressRepository(db *gorm.DB) domain.AddressRepository {
	return &addressRepository{db: db}
}

// ── Escrita ──────────────────────────────────────────────────────────────────

// Create implements [repositories.AddressRepository].
func (r *addressRepository) Create(address *entities.Address) error {
	return r.db.Create(address).Error
}

// Update implements [repositories.AddressRepository].
func (r *addressRepository) Update(address *entities.Address) error {
	return r.db.Save(address).Error
}

// Delete implements [repositories.AddressRepository].
func (r *addressRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Address{}, "id = ?", id).Error
}

// ── Leitura por chave única ───────────────────────────────────────────────────

// GetByID implements [repositories.AddressRepository].
func (r *addressRepository) GetByID(id uuid.UUID) (*entities.Address, error) {
	var address entities.Address
	if err := r.db.Where("id = ?", id).Take(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// GetByZipCode implements [repositories.AddressRepository].
func (r *addressRepository) GetByZipCode(zipCode string) (*entities.Address, error) {
	var address entities.Address
	if err := r.db.Where("zip_code = ?", zipCode).Take(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// ── Busca com filtros ─────────────────────────────────────────────────────────

// Search implements [repositories.AddressRepository].
func (r *addressRepository) Search(filters domain.AddressFilters) ([]*entities.Address, int64, error) {
	const PAGE_SIZE = 20

	var address []*entities.Address
	var total int64

	q := r.db.Model(&entities.Address{})

	if filters.City != nil {
		q = q.Where("city = ?", *filters.City)
	}
	if filters.State != nil {
		q = q.Where("state = ?", *filters.State)
	}
	if filters.Country != nil {
		q = q.Where("city = ?", *filters.City)
	}
	if filters.Neighborhood != nil {
		q = q.Where("neighborhood = ?", *filters.Neighborhood)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	pageSize := filters.PageSize
	if pageSize <= 0 {
		pageSize = PAGE_SIZE // valor padrão
	}
	page := filters.Page
	if page <= 0 {
		page = 1
	}

	if err := q.
		Order("created_at ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&address).Error; err != nil {
		return nil, 0, err
	}

	return address, total, nil
}
