// pkg/validators/cpf_test.go
package validators_test

import (
	"testing"

	"github.com/Turgho/Aluguei/pkg/validators"
)

func TestValidateCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{
			name:     "cpf válido sem máscara",
			cpf:      "52998224725",
			expected: true,
		},
		{
			name:     "cpf válido com máscara",
			cpf:      "529.982.247-25",
			expected: true,
		},
		{
			name:     "cpf inválido",
			cpf:      "11111111111",
			expected: false,
		},
		{
			name:     "cpf vazio",
			cpf:      "",
			expected: false,
		},
		{
			name:     "cpf com letras",
			cpf:      "abc",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.ValidateCPF(tt.cpf)

			if result != tt.expected {
				t.Errorf("esperado %v, recebido %v", tt.expected, result)
			}
		})
	}
}

func TestNormalizeCPF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "cpf com máscara",
			input:    "529.982.247-25",
			expected: "52998224725",
		},
		{
			name:     "cpf com espaços",
			input:    "529 982 247 25",
			expected: "52998224725",
		},
		{
			name:     "cpf já normalizado",
			input:    "52998224725",
			expected: "52998224725",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validators.NormalizeCPF(tt.input)

			if result != tt.expected {
				t.Errorf("esperado %s, recebido %s", tt.expected, result)
			}
		})
	}
}
