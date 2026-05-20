// Package property reúne validadores específicos para a entity Property.
package property

// ValidatePriceCents valida o preço do imóvel em centavos.
func ValidatePriceCents(value int64) bool {
	return value > 0
}

// ValidateOptionalFeeCents valida taxas opcionais em centavos.
func ValidateOptionalFeeCents(value int64) bool {
	return value >= 0
}
