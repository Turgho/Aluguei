package validators

import "net/mail"

// ValidateEmail verifica se o email possui formato válido.
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
