// pkg/validators/email_test.go
package validators_test

import (
	"testing"

	"github.com/Turgho/Aluguei/pkg/validators"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "email válido",
			email:    "teste@email.com",
			expected: true,
		},
		{
			name:     "email inválido sem arroba",
			email:    "testeemail.com",
			expected: false,
		},
		{
			name:     "email vazio",
			email:    "",
			expected: false,
		},
		{
			name:     "email inválido",
			email:    "@gmail",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidateEmail(tt.email)

			if result != tt.expected {
				t.Errorf("esperado %v, recebido %v", tt.expected, result)
			}
		})
	}
}
