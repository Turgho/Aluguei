// pkg/validators/property/text_test.go
package property_test

import (
	"strings"
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid", value: "Apartamento moderno", want: true},
		{name: "min boundary", value: "12345", want: true},
		{name: "too short", value: "abcd", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "   ", want: false},
		{name: "too long", value: strings.Repeat("a", 121), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateTitle(tt.value)
			if got != tt.want {
				t.Errorf("ValidateTitle(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateDescription(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid", value: "Descrição com mais de vinte caracteres.", want: true},
		{name: "min boundary", value: "12345678901234567890", want: true},
		{name: "too short", value: "curta", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "                     ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateDescription(tt.value)
			if got != tt.want {
				t.Errorf("ValidateDescription(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
