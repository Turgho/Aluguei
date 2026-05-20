package repositories

import (
	"github.com/Turgho/Aluguei/internal/domain/entities"
	domain "github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

// NewPropertyRepository retorna uma implementação de [domain.PropertyRepository].
func NewPropertyRepository(db *gorm.DB) domain.PropertyRepository {
	return &propertyRepository{db: db}
}

// ── Escrita ──────────────────────────────────────────────────────────────────

// Create implements [repositories.PropertyRepository].
func (r *propertyRepository) Create(property *entities.Property) error {
	return r.db.Create(property).Error
}

// Update implements [repositories.PropertyRepository].
func (r *propertyRepository) Update(property *entities.Property) error {
	return r.db.Save(property).Error
}

// Delete implements [repositories.PropertyRepository].
func (r *propertyRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Property{}, "id = ?", id).Error
}

// ── Leitura por chave única ───────────────────────────────────────────────────

// GetByID implements [repositories.PropertyRepository].
func (r *propertyRepository) GetByID(id uuid.UUID) (*entities.Property, error) {
	var property entities.Property
	if err := r.db.Where("id = ?", id).Take(&property).Error; err != nil {
		return nil, err
	}
	return &property, nil
}

// GetByOwnerID implements [repositories.PropertyRepository].
func (r *propertyRepository) GetByOwnerID(ownerID uuid.UUID) ([]*entities.Property, error) {
	var properties []*entities.Property
	if err := r.db.Where("owner_id = ?", ownerID).Find(&properties).Error; err != nil {
		return nil, err
	}
	return properties, nil
}

// ── Busca com filtros ─────────────────────────────────────────────────────────

// Search implements [repositories.PropertyRepository].
func (r *propertyRepository) Search(filters domain.PropertyFilters) ([]*entities.Property, int64, error) {
	const PAGE_SIZE = 20

	var properties []*entities.Property
	var total int64

	q := r.db.Model(&entities.Property{})

	if filters.Type != nil {
		q = q.Where("type = ?", *filters.Type)
	}
	if filters.Status != nil {
		q = q.Where("status = ?", *filters.Status)
	}
	if filters.MinPriceCents != nil {
		q = q.Where("price_cents >= ?", *filters.MinPriceCents)
	}
	if filters.MaxPriceCents != nil {
		q = q.Where("price_cents <= ?", *filters.MaxPriceCents)
	}
	if filters.MinAreaM2 != nil {
		q = q.Where("area_m2 >= ?", *filters.MinAreaM2)
	}
	if filters.Bedrooms != nil {
		q = q.Where("bedrooms = ?", *filters.Bedrooms)
	}
	if filters.Furnished != nil {
		q = q.Where("furnished = ?", *filters.Furnished)
	}
	if filters.PetFriendly != nil {
		q = q.Where("pet_friendly = ?", *filters.PetFriendly)
	}
	if filters.PetFriendly != nil {
		q = q.Where("pet_friendly = ?", *filters.PetFriendly)
	}
	// Filtro junto para não duplicar join
	if filters.City != nil || filters.State != nil {
		q = q.Joins("JOIN addresses ON addresses.id = properties.address_id")

		if filters.City != nil {
			q = q.Where("addresses.city ILIKE ?", *filters.City)
		}
		if filters.State != nil {
			q = q.Where("addresses.state ILIKE ?", *filters.State)
		}
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
		Find(&properties).Error; err != nil {
		return nil, 0, err
	}

	return properties, total, nil
}

// ── Controle de propriedade ─────────────────────────────────────────────────────────

// SetPublished implements [repositories.PropertyRepository].
func (r *propertyRepository) SetPublished(id uuid.UUID, published bool) error {
	return r.db.Model(&entities.Property{}).
		Where("id = ?", id).
		Update("is_published", true).Error
}
