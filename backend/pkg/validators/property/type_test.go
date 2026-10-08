// pkg/validators/property/type_test.go
package property_test

import (
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidateType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "house", value: "house", want: true},
		{name: "apartment", value: "apartment", want: true},
		{name: "studio", value: "studio", want: true},
		{name: "loft", value: "loft", want: true},
		{name: "business", value: "business", want: true},
		{name: "commercial", value: "commercial", want: true},
		{name: "office", value: "office", want: true},
		{name: "store", value: "store", want: true},
		{name: "land", value: "land", want: true},
		{name: "empty", value: "", want: false},
		{name: "invalid", value: "farm", want: false},
		{name: "uppercase", value: "HOUSE", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidateType(tt.value)
			if got != tt.want {
				t.Errorf("ValidateType(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
