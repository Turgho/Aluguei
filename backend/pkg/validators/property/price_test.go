// pkg/validators/property/price_test.go
package property_test

import (
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidatePriceCents(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  bool
	}{
		{name: "zero", value: 0, want: false},
		{name: "positive", value: 1, want: true},
		{name: "large value", value: 250000, want: true},
		{name: "negative", value: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidatePriceCents(tt.value)
			if got != tt.want {
				t.Errorf("ValidatePriceCents(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateOptionalFeeCents(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  bool
	}{
		{name: "zero", value: 0, want: true},
		{name: "positive", value: 1, want: true},
		{name: "large value", value: 35000, want: true},
		{name: "negative", value: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateOptionalFeeCents(tt.value)
			if got != tt.want {
				t.Errorf("ValidateOptionalFeeCents(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
