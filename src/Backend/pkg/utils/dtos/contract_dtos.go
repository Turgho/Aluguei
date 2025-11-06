package dtos

import "time"

type CreateContractRequest struct {
	PropertyID    string    `json:"property_id" validate:"required,uuid4"`
	TenantID      string    `json:"tenant_id" validate:"required,uuid4"`
	StartDate     time.Time `json:"start_date" validate:"required"`
	EndDate       time.Time `json:"end_date" validate:"required,gtfield=StartDate"`
	MonthlyRent   float64   `json:"monthly_rent" validate:"required,min=0"`
	DepositAmount float64   `json:"deposit_amount" validate:"required,min=0"`
	PaymentDueDay int       `json:"payment_due_day" validate:"required,min=1,max=28"`
}

type UpdateContractRequest struct {
	EndDate       time.Time `json:"end_date" validate:"omitempty"`
	MonthlyRent   float64   `json:"monthly_rent" validate:"omitempty,min=0"`
	PaymentDueDay int       `json:"payment_due_day" validate:"omitempty,min=1,max=28"`
}

type ContractResponse struct {
	ID            string           `json:"id"`
	PropertyID    string           `json:"property_id"`
	TenantID      string           `json:"tenant_id"`
	StartDate     time.Time        `json:"start_date"`
	EndDate       time.Time        `json:"end_date"`
	MonthlyRent   float64          `json:"monthly_rent"`
	DepositAmount float64          `json:"deposit_amount"`
	PaymentDueDay int              `json:"payment_due_day"`
	Status        string           `json:"status"`
	Property      PropertyResponse `json:"property,omitempty"`
	Tenant        TenantResponse   `json:"tenant,omitempty"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}
