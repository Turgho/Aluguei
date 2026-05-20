// pkg/validators/property/status_test.go
package property_test

import (
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidateStatus(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "available", value: "available", want: true},
		{name: "rented", value: "rented", want: true},
		{name: "inactive", value: "inactive", want: true},
		{name: "sold", value: "sold", want: true},
		{name: "empty", value: "", want: false},
		{name: "invalid", value: "pending", want: false},
		{name: "uppercase", value: "AVAILABLE", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateStatus(tt.value)
			if got != tt.want {
				t.Errorf("ValidateStatus(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
