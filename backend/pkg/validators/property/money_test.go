// pkg/validators/property/money_test.go
package property_test

import (
	"testing"

	propertyValidators "github.com/Turgho/Aluguei/pkg/validators/property"
)

func TestValidatePropertyMoney(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  bool
	}{
		{name: "zero", value: 0, want: true},
		{name: "positive", value: 1, want: true},
		{name: "large value", value: 999999, want: true},
		{name: "negative", value: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := propertyValidators.ValidatePropertyMoney(tt.value)
			if got != tt.want {
				t.Errorf("ValidatePropertyMoney(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
