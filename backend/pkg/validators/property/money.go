// Package property reúne validadores específicos para a entity Property.
package property

// ValidatePropertyMoney valida valores monetários opcionais do imóvel em centavos.
func ValidatePropertyMoney(value int64) bool {
	return value >= 0
}
