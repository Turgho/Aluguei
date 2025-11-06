package repositories

import (
	"github.com/Turgho/Aluguei/internal/errors"
	"github.com/Turgho/Aluguei/internal/models"
	"gorm.io/gorm"
)

type OwnerRepository interface {
	Create(owner *models.Owner) *errors.AppError
	FindByID(id string) (*models.Owner, *errors.AppError)
	FindByEmail(email string) (*models.Owner, *errors.AppError)
	Update(owner *models.Owner) *errors.AppError
	Delete(id string) *errors.AppError
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

func (r *ownerRepository) FindByEmail(email string) (*models.Owner, *errors.AppError) {
	var owner models.Owner
	err := r.db.Where("email = ?", email).First(&owner).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("proprietário", email)
		}
		return nil, errors.NewDatabaseError("erro ao buscar proprietário por email", err)
	}
	return &owner, nil
}
