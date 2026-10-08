package validators

import "unicode"

// ValidatePassword valida uma senha com regras mínimas.
//
// Regras:
// - mínimo de 8 caracteres
// - ao menos 1 letra maiúscula
// - ao menos 1 letra minúscula
// - ao menos 1 número
// - ao menos 1 caractere especial
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
