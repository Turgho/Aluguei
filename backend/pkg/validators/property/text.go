// Package property reúne validadores específicos para a entity Property.
package property

import "strings"

// ValidateTitle valida o título do imóvel.
func ValidateTitle(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 5 && len(value) <= 120
}

// ValidateDescription valida a descrição do imóvel.
func ValidateDescription(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 20
}
