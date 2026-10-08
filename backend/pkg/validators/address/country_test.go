// Package address reúne validadores específicos para a entity Address.
package address_test

import (
	"testing"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

func TestValidateCountry(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "brazil uppercase", value: "BR", want: true},
		{name: "brazil lowercase", value: "br", want: true},
		{name: "usa", value: "US", want: true},
		{name: "portugal", value: "PT", want: true},

		{name: "empty", value: "", want: false},
		{name: "one char", value: "B", want: false},
		{name: "three chars", value: "BRA", want: false},

		{name: "invalid country", value: "ZZ", want: false},
		{name: "numbers not allowed", value: "1P", want: false},
		{name: "symbols not allowed", value: "@#", want: false},
		{name: "mixed invalid", value: "B1", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateCountry(tt.value)

			if got != tt.want {
				t.Errorf("ValidateCountry(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
