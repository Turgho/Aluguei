// Package shared reúne validadores genéricos de números reutilizáveis em várias entities.
package shared

// ValidatePositiveInt valida se o número é maior que zero.
func ValidatePositiveInt(value int) bool {
	return value > 0
}

// ValidateNonNegativeInt valida se o número é maior ou igual a zero.
func ValidateNonNegativeInt(value int) bool {
	return value >= 0
}

// ValidatePositiveInt64 valida se o número é maior que zero.
func ValidatePositiveInt64(value int64) bool {
	return value > 0
}

// ValidateNonNegativeInt64 valida se o número é maior ou igual a zero.
func ValidateNonNegativeInt64(value int64) bool {
	return value >= 0
}
