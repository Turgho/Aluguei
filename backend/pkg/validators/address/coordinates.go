// Package address reúne validadores específicos para a entity Address.
package address

// ValidateLatitude valida latitude geográfica.
func ValidateLatitude(value float64) bool {
	return value >= -90 && value <= 90
}

// ValidateLongitude valida longitude geográfica.
func ValidateLongitude(value float64) bool {
	return value >= -180 && value <= 180
}
