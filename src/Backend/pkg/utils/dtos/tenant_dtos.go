package dtos

import "time"

type CreateTenantRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=255"`
	Email string `json:"email" validate:"omitempty,email"`
	Phone string `json:"phone" validate:"required,min=10,max=20"`
	CPF   string `json:"cpf" validate:"omitempty,min=11,max=14"`
}

type UpdateTenantRequest struct {
	Name  string `json:"name" validate:"omitempty,min=2,max=255"`
	Email string `json:"email" validate:"omitempty,email"`
	Phone string `json:"phone" validate:"omitempty,min=10,max=20"`
	CPF   string `json:"cpf" validate:"omitempty,min=11,max=14"`
}

type TenantResponse struct {
	ID        string    `json:"id"`
	OwnerID   string    `json:"owner_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CPF       string    `json:"cpf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
