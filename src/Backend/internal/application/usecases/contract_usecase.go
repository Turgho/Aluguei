package usecases

import (
	"context"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/repositories"
	"github.com/google/uuid"
)

type ContractUseCase struct {
	contractRepo repositories.ContractRepository
}

func NewContractUseCase(contractRepo repositories.ContractRepository) *ContractUseCase {
	return &ContractUseCase{
		contractRepo: contractRepo,
	}
}

func (uc *ContractUseCase) CreateContract(ctx context.Context, propertyID, tenantID uuid.UUID, startDate, endDate time.Time, monthlyRent float64, paymentDueDay int, status entities.ContractStatus) (*entities.Contract, error) {
	contract := entities.NewContract(propertyID, tenantID, startDate, endDate, monthlyRent, paymentDueDay, status)
	
	if err := uc.contractRepo.Create(ctx, contract); err != nil {
		return nil, err
	}
	
	return contract, nil
}

func (uc *ContractUseCase) GetContract(ctx context.Context, id uuid.UUID) (*entities.Contract, error) {
	return uc.contractRepo.GetByID(ctx, id)
}

func (uc *ContractUseCase) GetContractsByProperty(ctx context.Context, propertyID uuid.UUID) ([]*entities.Contract, error) {
	return uc.contractRepo.GetByPropertyID(ctx, propertyID)
}

func (uc *ContractUseCase) GetContractsByTenant(ctx context.Context, tenantID uuid.UUID) ([]*entities.Contract, error) {
	return uc.contractRepo.GetByTenantID(ctx, tenantID)
}

func (uc *ContractUseCase) GetActiveContractByProperty(ctx context.Context, propertyID uuid.UUID) (*entities.Contract, error) {
	return uc.contractRepo.GetActiveByPropertyID(ctx, propertyID)
}

func (uc *ContractUseCase) GetAllContracts(ctx context.Context, page, limit int, status string) ([]*entities.Contract, int64, error) {
	return uc.contractRepo.GetAll(ctx, page, limit, status)
}

func (uc *ContractUseCase) UpdateContract(ctx context.Context, id uuid.UUID, startDate, endDate time.Time, monthlyRent float64, paymentDueDay int, status entities.ContractStatus) error {
	contract, err := uc.contractRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	contract.Update(startDate, endDate, monthlyRent, paymentDueDay, status)
	
	return uc.contractRepo.Update(ctx, contract)
}

func (uc *ContractUseCase) DeleteContract(ctx context.Context, id uuid.UUID) error {
	return uc.contractRepo.Delete(ctx, id)
}

func (uc *ContractUseCase) GetContractByID(ctx context.Context, id uuid.UUID) (*entities.Contract, error) {
	return uc.contractRepo.GetByID(ctx, id)
}