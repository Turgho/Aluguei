// Package address reúne validadores específicos para a entity Address.
package address

import (
	"regexp"
	"strings"
)

var zipCodeRegex = regexp.MustCompile(`^\d{5}-?\d{3}$`)

// ValidateZipCode valida CEP (BR por enquanto).
func ValidateZipCode(country, value string) bool {
	country = strings.ToUpper(strings.TrimSpace(country))
	value = strings.TrimSpace(value)

	switch country {
	case "BR":
		return zipCodeRegex.MatchString(value)
	default:
		return len(value) >= 3 && len(value) <= 20
	}
}

// NormalizeZipCode remove caracteres não numéricos do CEP.
func NormalizeZipCode(country, value string) string {
	value = strings.TrimSpace(value)

	switch strings.ToUpper(country) {
	case "BR":
		re := regexp.MustCompile(`\D`)
		return re.ReplaceAllString(value, "")
	default:
		return value
	}
}
