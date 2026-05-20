// Package property reúne validadores específicos para a entity Property.
package property

// ValidateAreaM2 valida a área do imóvel.
func ValidateAreaM2(value int) bool {
	return value > 0
}

// ValidateRooms validates counts like bedrooms, bathrooms, suites and parking spaces.
func ValidateRooms(value int) bool {
	return value >= 0
}
