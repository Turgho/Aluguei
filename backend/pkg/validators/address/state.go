// Package address reúne validadores específicos para a entity Address.
package address

import "strings"

var validBRStates = map[string]struct{}{
	"AC": {}, "AL": {}, "AP": {}, "AM": {}, "BA": {},
	"CE": {}, "DF": {}, "ES": {}, "GO": {}, "MA": {},
	"MT": {}, "MS": {}, "MG": {}, "PA": {}, "PB": {},
	"PR": {}, "PE": {}, "PI": {}, "RJ": {}, "RN": {},
	"RS": {}, "RO": {}, "RR": {}, "SC": {}, "SP": {},
	"SE": {}, "TO": {},
}

// ValidateState valida a UF do endereço brasileiro.
func ValidateState(value string) bool {
	value = strings.ToUpper(strings.TrimSpace(value))

	_, exists := validBRStates[value]

	return exists
}
