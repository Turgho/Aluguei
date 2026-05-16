// Package entities define as entidades de domínio da aplicação.
package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/Turgho/Aluguei/pkg/validators"
	"github.com/google/uuid"
)

// Role representa o papel de um usuário no sistema.
type Role string

const (
	// RoleOwner representa um proprietário de imóvel.
	RoleOwner Role = "owner"

	// RoleTenant representa um inquilino.
	RoleTenant Role = "tenant"
)

// User representa um usuário do sistema.
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`
	FirstName    string    `gorm:"type:varchar(100);not null"`
	LastName     string    `gorm:"type:varchar(100);not null"`
	CPF          string    `gorm:"type:varchar(14);uniqueIndex;not null"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone        string    `gorm:"type:varchar(20)"`
	PasswordHash string    `json:"-" gorm:"type:varchar(255);not null"`
	Role         Role      `gorm:"type:varchar(50);not null;default:tenant"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// NewUser cria e valida uma nova instância de [User].
//
// Retorna erro se algum campo obrigatório estiver ausente ou inválido.
// Todos os erros de validação são retornados juntos, separados por ponto e vírgula.
func NewUser(firstName, lastName, cpf, email, phone, passwordHash string, role Role) (*User, error) {
	var errs []string

	// Validações
	// ————— Nome —————
	if firstName == "" {
		errs = append(errs, "nome é obrigatório")
	}
	if lastName == "" {
		errs = append(errs, "sobrenome é obrigatório")
	}

	// ————— CPF —————
	if cpf == "" {
		errs = append(errs, "CPF é obrigatório")
	} else if !validators.ValidateCPF(cpf) {
		errs = append(errs, "CPF inválido")
	}
	cpf = validators.NormalizeCPF(cpf)

	// ————— Email —————
	if email == "" {
		errs = append(errs, "email é obrigatório")
	} else if !validators.ValidateEmail(email) {
		errs = append(errs, "email inválido")
	}

	// ————— Telefone —————
	if !validators.ValidatePhone(phone) {
		errs = append(errs, "telefone inválido")
	}
	phone = validators.NormalizePhone(phone)

	// ————— Senha —————
	if passwordHash == "" {
		errs = append(errs, "senha é obrigatória")
	}
	if role != RoleOwner && role != RoleTenant {
		errs = append(errs, "role inválida")
	}

	// Erros
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
