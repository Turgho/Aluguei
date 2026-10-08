// pkg/validators/password_test.go
package validators_test

import (
	"testing"

	"github.com/Turgho/Aluguei/pkg/validators"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "senha forte",
			password: "Senha123@",
			expected: true,
		},
		{
			name:     "senha curta",
			password: "123",
			expected: false,
		},
		{
			name:     "sem número",
			password: "SenhaForte@",
			expected: false,
		},
		{
			name:     "sem letra maiúscula",
			password: "senha123@",
			expected: false,
		},
		{
			name:     "sem caractere especial",
			password: "Senha123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidatePassword(tt.password)

			if result != tt.expected {
				t.Errorf("esperado %v, recebido %v", tt.expected, result)
			}
		})
	}
}
