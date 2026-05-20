// Package shared reúne validadores genéricos de UUID reutilizáveis em várias entities.
package shared

import "github.com/google/uuid"

// ValidateUUID valida se o UUID foi preenchido.
func ValidateUUID(value uuid.UUID) bool {
	return value != uuid.Nil
}
