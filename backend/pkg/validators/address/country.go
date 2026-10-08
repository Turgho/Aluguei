// Package address reúne validadores específicos para a entity Address.
package address

import "strings"

// ValidateCountry valida o código de país no padrão ISO 3166-1 alpha-2.
//
// Regras:
// - exatamente 2 caracteres
// - apenas letras A-Z
// - precisa existir na lista suportada pela aplicação
func ValidateCountry(value string) bool {
	value = strings.ToUpper(strings.TrimSpace(value))

	return validCountries[value]
}

// lista inicial de países suportados (pode crescer depois)
var validCountries = map[string]bool{
	"BR": true,
	"US": true,
	"PT": true,
}
