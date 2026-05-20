// Package property reúne validadores específicos para a entity Property.
package property

// ValidateStatus valida se o status do imóvel é um valor aceito pela aplicação.
func ValidateStatus(value string) bool {
	switch value {
	case "available",
		"rented",
		"inactive",
		"sold":
		return true
	default:
		return false
	}
}
