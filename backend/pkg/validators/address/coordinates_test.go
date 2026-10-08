// pkg/validators/address/coordinates_test.go
package address_test

import (
	"testing"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

func TestValidateLatitude(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  bool
	}{
		{name: "zero", value: 0, want: true},
		{name: "positive valid", value: 22.314, want: true},
		{name: "negative valid", value: -22.314, want: true},
		{name: "upper boundary", value: 90, want: true},
		{name: "lower boundary", value: -90, want: true},
		{name: "above upper boundary", value: 90.1, want: false},
		{name: "below lower boundary", value: -90.1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateLatitude(tt.value)
			if got != tt.want {
				t.Errorf("ValidateLatitude(%f) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateLongitude(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  bool
	}{
		{name: "zero", value: 0, want: true},
		{name: "positive valid", value: 49.058, want: true},
		{name: "negative valid", value: -49.058, want: true},
		{name: "upper boundary", value: 180, want: true},
		{name: "lower boundary", value: -180, want: true},
		{name: "above upper boundary", value: 180.1, want: false},
		{name: "below lower boundary", value: -180.1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateLongitude(tt.value)
			if got != tt.want {
				t.Errorf("ValidateLongitude(%f) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
