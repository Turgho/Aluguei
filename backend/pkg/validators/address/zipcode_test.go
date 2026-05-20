// pkg/validators/address/zipcode_test.go
package address_test

import (
	"testing"

	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

func TestValidateZipCode(t *testing.T) {
	tests := []struct {
		name    string
		country string
		value   string
		want    bool
	}{
		{name: "BR valid with dash", country: "BR", value: "17000-000", want: true},
		{name: "BR valid without dash", country: "BR", value: "17000000", want: true},
		{name: "BR invalid short", country: "BR", value: "17000", want: false},
		{name: "BR invalid letters", country: "BR", value: "abcde-fff", want: false},
		{name: "BR empty", country: "BR", value: "", want: false},

		// fallback genérico
		{name: "US-like valid", country: "US", value: "12345", want: true},
		{name: "US-like extended", country: "US", value: "12345-6789", want: true},
		{name: "US-like invalid", country: "US", value: "12", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.ValidateZipCode(tt.country, tt.value)

			if got != tt.want {
				t.Errorf(
					"ValidateZipCode(%q, %q) = %v, want %v",
					tt.country,
					tt.value,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestNormalizeZipCode(t *testing.T) {
	tests := []struct {
		name    string
		country string
		value   string
		want    string
	}{
		{name: "BR with dash", country: "BR", value: "17000-000", want: "17000000"},
		{name: "BR with spaces", country: "BR", value: " 17000-000 ", want: "17000000"},
		{name: "BR already clean", country: "BR", value: "17000000", want: "17000000"},
		{name: "BR with letters", country: "BR", value: "17a000-0b00", want: "17000000"},

		// fallback genérico (caso suporte outros países)
		{name: "US unchanged", country: "US", value: "12345-6789", want: "12345-6789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressValidators.NormalizeZipCode(tt.country, tt.value)

			if got != tt.want {
				t.Errorf(
					"NormalizeZipCode(%q, %q) = %q, want %q",
					tt.country,
					tt.value,
					got,
					tt.want,
				)
			}
		})
	}
}
