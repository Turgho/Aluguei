// pkg/validators/property/area_test.go
package property_test

import (
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidateAreaM2(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "zero", value: 0, want: false},
		{name: "positive", value: 1, want: true},
		{name: "large value", value: 120, want: true},
		{name: "negative", value: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateAreaM2(tt.value)
			if got != tt.want {
				t.Errorf("ValidateAreaM2(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestValidateRooms(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "zero", value: 0, want: true},
		{name: "positive", value: 1, want: true},
		{name: "large value", value: 5, want: true},
		{name: "negative", value: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateRooms(tt.value)
			if got != tt.want {
				t.Errorf("ValidateRooms(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
