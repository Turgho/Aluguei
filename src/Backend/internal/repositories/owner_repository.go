package repositories

import (
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type OwnerRepository interface {
	Create(owner *models.Owner) error
	FindByID(id string) (*models.Owner, error)
	FindByEmail(email string) (*models.Owner, error)
	Update(owner *models.Owner) error
	Delete(id string) error
}

type ownerRepository struct {
	*BaseRepository[models.Owner]
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository {
	return &ownerRepository{
		BaseRepository: NewBaseRepository[models.Owner](db),
		db:             db,
	}
}

func (r *ownerRepository) FindByEmail(email string) (*models.Owner, error) {
	var owner models.Owner
	err := r.db.Where("email = ?", email).First(&owner).Error
	if err != nil {
		return nil, err
	}
	return &owner, nil
}
