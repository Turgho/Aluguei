package usecases

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
)

type TenantUseCase struct {
	tenantRepo repositories.TenantRepository
}

func NewTenantUseCase(tenantRepo repositories.TenantRepository) *TenantUseCase {
	return &TenantUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *TenantUseCase) CreateTenant(ctx context.Context, name, email, phone, cpf string, ownerID uuid.UUID, birthDate *time.Time) (*entities.Tenant, error) {
	tenant := entities.NewTenant(name, email, phone, cpf, ownerID, birthDate)
	
	if err := uc.tenantRepo.Create(ctx, tenant); err != nil {
		return nil, err
	}
	
	return tenant, nil
}

func (uc *TenantUseCase) GetTenant(ctx context.Context, id uuid.UUID) (*entities.Tenant, error) {
	return uc.tenantRepo.GetByID(ctx, id)
}

func (uc *TenantUseCase) GetTenantsByOwner(ctx context.Context, ownerID uuid.UUID) ([]*entities.Tenant, error) {
	return uc.tenantRepo.GetByOwnerID(ctx, ownerID)
}

func (uc *TenantUseCase) GetAllTenants(ctx context.Context, page, limit int) ([]*entities.Tenant, int64, error) {
	return uc.tenantRepo.GetAll(ctx, page, limit)
}

func (uc *TenantUseCase) UpdateTenant(ctx context.Context, id uuid.UUID, name, email, phone, cpf string, birthDate *time.Time) error {
	tenant, err := uc.tenantRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	tenant.UpdateProfile(name, email, phone, cpf, birthDate)
	
	return uc.tenantRepo.Update(ctx, tenant)
}

func (uc *TenantUseCase) DeleteTenant(ctx context.Context, id uuid.UUID) error {
	return uc.tenantRepo.Delete(ctx, id)
}