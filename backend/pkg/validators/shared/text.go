// Package shared reúne validadores genéricos de texto reutilizáveis em várias entities.
package shared

import "strings"

// ValidateRequiredText valida se o texto não está vazio após remover espaços.
func ValidateRequiredText(value string) bool {
	return strings.TrimSpace(value) != ""
}

// ValidateTextLength valida se o texto possui tamanho entre min e max, após trim.
func ValidateTextLength(value string, min, max int) bool {
	value = strings.TrimSpace(value)
	l := len(value)

	return l >= min && l <= max
}
