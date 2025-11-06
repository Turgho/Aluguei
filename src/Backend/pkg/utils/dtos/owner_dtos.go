package dtos

import "time"

type CreateOwnerRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"omitempty,min=10,max=20"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateOwnerRequest struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=255"`
	Phone    string `json:"phone" validate:"omitempty,min=10,max=20"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

type OwnerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Owner OwnerResponse `json:"owner"`
	Token string        `json:"token"`
}
