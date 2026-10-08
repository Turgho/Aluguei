// pkg/validators/phone.go
package validators

import "strings"

// onlyDigits remove todos os caracteres não numéricos.
func onlyDigits(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// ValidatePhone valida números brasileiros.
//
// Aceita:
// - vazio
// - (14) 99999-9999
// - 14999999999
// - +5514999999999
func ValidatePhone(phone string) bool {
	if strings.TrimSpace(phone) == "" {
		return true
	}

	phone = onlyDigits(phone)

	// telefone BR:
	// 10 dígitos = fixo com DDD
	// 11 dígitos = celular com DDD
	// 13 dígitos = 55 + DDD + número
	switch len(phone) {
	case 10, 11:
		return true
	case 13:
		return strings.HasPrefix(phone, "55")
	default:
		return false
	}
}

// NormalizePhone converte o telefone para padrão internacional.
//
// Exemplos:
//
//	(14) 99999-9999 -> 5514999999999
//	14999999999     -> 5514999999999
//	5514999999999   -> 5514999999999
func NormalizePhone(phone string) string {
	phone = onlyDigits(phone)

	if phone == "" {
		return ""
	}

	// adiciona código do Brasil
	if len(phone) == 10 || len(phone) == 11 {
		phone = "55" + phone
	}

	return phone
}
