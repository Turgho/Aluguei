package validators

import "strings"

// NormalizeCPF remove caracteres não numéricos.
func NormalizeCPF(cpf string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)
}

// ValidateCPF valida um CPF brasileiro.
func ValidateCPF(cpf string) bool {
	// remove pontos e traço
	cpf = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)

	if len(cpf) != 11 {
		return false
	}

	// rejeita CPFs com todos os dígitos iguais
	if cpf == strings.Repeat(string(cpf[0]), 11) {
		return false
	}

	digits := make([]int, 11)
	for i := range 11 {
		digits[i] = int(cpf[i] - '0')
	}

	// primeiro dígito
	sum := 0
	for i := range 9 {
		sum += digits[i] * (10 - i)
	}
	rem := sum % 11
	d1 := 0
	if rem >= 2 {
		d1 = 11 - rem
	}
	if digits[9] != d1 {
		return false
	}

	// segundo dígito
	sum = 0
	for i := range 10 {
		sum += digits[i] * (11 - i)
	}
	rem = sum % 11
	d2 := 0
	if rem >= 2 {
		d2 = 11 - rem
	}
	return digits[10] == d2
}
