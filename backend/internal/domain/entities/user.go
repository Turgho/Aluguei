// Package entities define as entidades de domínio da aplicação.
package entities

import (
	"errors"
	"strings"
	"time"

	userValidators "github.com/Turgho/Aluguei/pkg/validators/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`

	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`

	CPF   string `gorm:"type:varchar(11);uniqueIndex;not null"`
	Email string `gorm:"type:varchar(255);uniqueIndex;not null;index"`
	Phone string `gorm:"type:varchar(15)"`

	PasswordHash string `json:"-" gorm:"type:varchar(255);not null"`

	Role Role `gorm:"type:varchar(50);not null;default:owner;index"`

	// Controle da conta
	IsActive      bool       `gorm:"not null;default:true"`
	EmailVerified bool       `gorm:"not null;default:false"`
	LastLoginAt   *time.Time `gorm:"default:null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// NewUser cria e valida uma nova instância de [User].
//
// Retorna erro se algum campo obrigatório estiver ausente ou inválido.
// Todos os erros de validação são retornados juntos, separados por ponto e vírgula.
func NewUser(firstName, lastName, cpf, email, phone, passwordHash string, role Role) (*User, error) {
	var errs []string

	// Normalização
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	cpf = userValidators.NormalizeCPF(cpf)

	email = strings.TrimSpace(strings.ToLower(email))

	phone = userValidators.NormalizePhone(phone)

	passwordHash = strings.TrimSpace(passwordHash)

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
	} else if !userValidators.ValidateCPF(cpf) {
		errs = append(errs, "CPF inválido")
	}

	// ————— Email —————
	if email == "" {
		errs = append(errs, "email é obrigatório")
	} else if !userValidators.ValidateEmail(email) {
		errs = append(errs, "email inválido")
	}

	// ————— Telefone —————
	if phone != "" && !userValidators.ValidatePhone(phone) {
		errs = append(errs, "telefone inválido")
	}

	// ————— Senha —————
	if passwordHash == "" {
		errs = append(errs, "senha é obrigatória")
	}

	// ————— Role —————
	if !userValidators.ValidateRole(string(role)) {
		errs = append(errs, "role inválida")
	}

	// ————— Erros —————
	if len(errs) > 0 {
		return nil, errors.New(strings.Join(errs, "; "))
	}

	now := time.Now().UTC()

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		CPF:           cpf,
		Email:         email,
		Phone:         phone,
		PasswordHash:  passwordHash,
		Role:          role,
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}
