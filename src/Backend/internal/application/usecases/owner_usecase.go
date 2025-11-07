package usecases

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type OwnerUseCase struct {
	ownerRepo repositories.OwnerRepository
}

func NewOwnerUseCase(ownerRepo repositories.OwnerRepository) *OwnerUseCase {
	return &OwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *OwnerUseCase) CreateOwner(ctx context.Context, name, email, password, phone, cpf string, birthDate *time.Time) (*entities.Owner, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	owner := entities.NewOwner(name, email, string(hashedPassword), phone, cpf, birthDate)
	
	if err := uc.ownerRepo.Create(ctx, owner); err != nil {
		return nil, err
	}
	
	return owner, nil
}

func (uc *OwnerUseCase) GetOwner(ctx context.Context, id uuid.UUID) (*entities.Owner, error) {
	return uc.ownerRepo.GetByID(ctx, id)
}

func (uc *OwnerUseCase) GetOwnerByEmail(ctx context.Context, email string) (*entities.Owner, error) {
	return uc.ownerRepo.GetByEmail(ctx, email)
}

func (uc *OwnerUseCase) GetAllOwners(ctx context.Context, page, limit int) ([]*entities.Owner, int64, error) {
	return uc.ownerRepo.GetAll(ctx, page, limit)
}

func (uc *OwnerUseCase) UpdateOwner(ctx context.Context, id uuid.UUID, name, email, phone, cpf string, birthDate *time.Time) error {
	owner, err := uc.ownerRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	owner.UpdateProfile(name, email, phone, cpf, birthDate)
	
	return uc.ownerRepo.Update(ctx, owner)
}

func (uc *OwnerUseCase) DeleteOwner(ctx context.Context, id uuid.UUID) error {
	return uc.ownerRepo.Delete(ctx, id)
}

func (uc *OwnerUseCase) ValidatePassword(ctx context.Context, email, password string) (*entities.Owner, error) {
	owner, err := uc.ownerRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(password)); err != nil {
		return nil, err
	}

	return owner, nil
}