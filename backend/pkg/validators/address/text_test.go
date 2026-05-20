// pkg/validators/address/text_test.go
package address_test

import (
	"strings"
	"testing"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

func TestValidateStreet(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid", value: "Rua das Flores", want: true},
		{name: "min boundary", value: "Rua", want: true},
		{name: "too short", value: "Ru", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "   ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateStreet(tt.value)
			if got != tt.want {
				t.Errorf("ValidateStreet(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid number", value: "123", want: true},
		{name: "valid with letter", value: "123A", want: true},
		{name: "min boundary", value: "1", want: true},
		{name: "too long", value: "123456789012345678901", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "   ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateNumber(tt.value)
			if got != tt.want {
				t.Errorf("ValidateNumber(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateComplement(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "empty allowed", value: "", want: true},
		{name: "valid", value: "Apto 101", want: true},
		{name: "spaces only", value: "   ", want: true},
		{name: "max valid", value: "1234567890", want: true},
		{name: "too long", value: strings.Repeat("a", 121), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateComplement(tt.value)
			if got != tt.want {
				t.Errorf("ValidateComplement(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateNeighborhood(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid", value: "Centro", want: true},
		{name: "min boundary", value: "AB", want: true},
		{name: "too short", value: "A", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "   ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateNeighborhood(tt.value)
			if got != tt.want {
				t.Errorf("ValidateNeighborhood(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateCity(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid", value: "Bauru", want: true},
		{name: "min boundary", value: "AB", want: true},
		{name: "too short", value: "A", want: false},
		{name: "empty", value: "", want: false},
		{name: "spaces only", value: "   ", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateCity(tt.value)
			if got != tt.want {
				t.Errorf("ValidateCity(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
