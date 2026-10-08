// pkg/validators/phone_test.go
package validators_test

import (
	"testing"

	"github.com/Turgho/Aluguei/pkg/validators"
)

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		{
			name:     "telefone válido com máscara",
			phone:    "(14) 99999-9999",
			expected: true,
		},
		{
			name:     "telefone válido internacional",
			phone:    "+5514999999999",
			expected: true,
		},
		{
			name:     "telefone vazio",
			phone:    "",
			expected: true,
		},
		{
			name:     "telefone inválido",
			phone:    "123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidatePhone(tt.phone)

			if result != tt.expected {
				t.Errorf("esperado %v, recebido %v", tt.expected, result)
			}
		})
	}
}

func TestNormalizePhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "telefone com máscara",
			input:    "(14) 99999-9999",
			expected: "5514999999999",
		},
		{
			name:     "telefone já internacional",
			input:    "5514999999999",
			expected: "5514999999999",
		},
		{
			name:     "telefone vazio",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.NormalizePhone(tt.input)

			if result != tt.expected {
				t.Errorf("esperado %s, recebido %s", tt.expected, result)
			}
		})
	}
}
