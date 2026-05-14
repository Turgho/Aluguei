package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleOwner  Role = "owner"
	RoleTenant Role = "tenant"
)

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	CPF          string
	Email        string
	Phone        string
	PasswordHash string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(firstName, lastName, cpf, email, phone, passwordHash string, role Role) (*User, error) {
	var errs []string

	if firstName == "" {
		errs = append(errs, "nome é obrigatório")
	}

	if lastName == "" {
		errs = append(errs, "sobrenome é obrigatório")
	}

	if cpf == "" {
		errs = append(errs, "CPF é obrigatório")
	}

	if email == "" || !strings.Contains(email, "@") {
		errs = append(errs, "email é obrigatório ou tem formato inválido")
	}

	if passwordHash == "" {
		errs = append(errs, "senha é obrigatório")
	}

	if role != RoleOwner && role != RoleTenant {
		errs = append(errs, "role inválida")
	}

	// retorna todos os erros
	if len(errs) > 0 {
		return nil, errors.New(strings.Join(errs, "; "))
	}

	now := time.Now().UTC()
	return &User{
		FirstName:    firstName,
		LastName:     lastName,
		CPF:          cpf,
		Email:        email,
		Phone:        phone,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
