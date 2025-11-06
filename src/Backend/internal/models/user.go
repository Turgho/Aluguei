package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleOwner  UserRole = "owner"
	RoleTenant UserRole = "tenant"
	RoleAdmin  UserRole = "admin"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Phone     string         `json:"phone"`
	Password  string         `json:"-" gorm:"not null"` // - para não serializar
	Role      UserRole       `json:"role" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relacionamentos
	Properties []*Property `json:"properties,omitempty" gorm:"foreignKey:OwnerID"`
	Contracts  []*Contract `json:"contracts,omitempty" gorm:"foreignKey:TenantID"`
}

// BeforeCreate é um hook do GORM para gerar UUID antes de criar
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
