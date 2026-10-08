// Package property reúne validadores específicos para a entity Property.
package property

// ValidateType valida se o tipo do imóvel é um valor aceito pela aplicação.
func ValidateType(value string) bool {
	switch value {
	case "house",
		"apartment",
		"studio",
		"loft",
		"business",
		"commercial",
		"office",
		"store",
		"land":
		return true
	default:
		return false
	}
}
