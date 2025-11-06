package repositories

import "gorm.io/gorm"

type Repository struct {
	Owner    OwnerRepository
	Property PropertyRepository
	Tenant   TenantRepository
	Contract ContractRepository
	Payment  PaymentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Owner:    NewOwnerRepository(db),
		Property: NewPropertyRepository(db),
		Tenant:   NewTenantRepository(db),
		Contract: NewContractRepository(db),
		Payment:  NewPaymentRepository(db),
	}
}
