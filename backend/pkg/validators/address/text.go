// Package address reúne validadores específicos para a entity Address.
package address

import "strings"

// ValidateStreet valida o nome da rua.
func ValidateStreet(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 3 && len(value) <= 120
}

// ValidateNumber valida o número do endereço.
func ValidateNumber(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 1 && len(value) <= 20
}

// ValidateComplement valida o complemento do endereço.
func ValidateComplement(value string) bool {
	value = strings.TrimSpace(value)

	if value == "" {
		return true
	}

	return len(value) <= 100
}

// ValidateNeighborhood valida o bairro do endereço.
func ValidateNeighborhood(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 2 && len(value) <= 100
}

// ValidateCity valida a cidade do endereço.
func ValidateCity(value string) bool {
	value = strings.TrimSpace(value)
	return len(value) >= 2 && len(value) <= 100
}
