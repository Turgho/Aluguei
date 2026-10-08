// Package shared reúne validadores genéricos de valores monetários em centavos.
package shared

// ValidateMoney valida valores monetários aceitando zero ou positivo.
func ValidateMoney(value int64) bool {
	return value >= 0
}

// ValidateRequiredMoney valida valores monetários que precisam ser maiores que zero.
func ValidateRequiredMoney(value int64) bool {
	return value > 0
}
