// pkg/validators/address/state_test.go
package address_test

import (
	"testing"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

func TestValidateState(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "valid uppercase", value: "SP", want: true},
		{name: "valid lowercase", value: "sp", want: true},
		{name: "valid with spaces", value: " sp ", want: true},
		{name: "valid rio de janeiro", value: "RJ", want: true},

		{name: "empty", value: "", want: false},
		{name: "one char", value: "S", want: false},
		{name: "three chars", value: "SPO", want: false},

		{name: "invalid state", value: "ZZ", want: false},
		{name: "numbers not allowed", value: "1P", want: false},
		{name: "symbols not allowed", value: "@#", want: false},
		{name: "mixed invalid", value: "S1", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateState(tt.value)

			if got != tt.want {
				t.Errorf(
					"ValidateState(%q) = %v, want %v",
					tt.value,
					got,
					tt.want,
				)
			}
		})
	}
}
