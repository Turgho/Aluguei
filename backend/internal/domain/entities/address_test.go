// internal/domain/entities/address_test.go
package entities_test

import (
	"strings"
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	addressValidators "github.com/Turgho/Aluguei/pkg/validators/address"
)

type addressTestCase struct {
	name         string
	street       string
	number       string
	complement   string
	neighborhood string
	city         string
	state        string
	country      string
	zipCode      string
	latitude     float64
	longitude    float64

	wantErr bool
}

func TestNewAddress(t *testing.T) {
	tests := []addressTestCase{
		{
			name:         "endereço válido",
			street:       "Rua das Flores",
			number:       "123",
			complement:   "Apto 101",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SP",
			country:      "BR",
			zipCode:      "17000-000",
			latitude:     -22.314,
			longitude:    -49.058,
			wantErr:      false,
		},
		{
			name:         "rua vazia",
			street:       "",
			number:       "123",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SP",
			country:      "BR",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "número vazio",
			street:       "Rua das Flores",
			number:       "",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SP",
			country:      "BR",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "bairro vazio",
			street:       "Rua das Flores",
			number:       "123",
			neighborhood: "",
			city:         "Bauru",
			state:        "SP",
			country:      "BR",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "cidade vazia",
			street:       "Rua das Flores",
			number:       "123",
			neighborhood: "Centro",
			city:         "",
			state:        "SP",
			country:      "BR",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "estado inválido",
			street:       "Rua das Flores",
			number:       "123",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SPO",
			country:      "BR",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "país vazio",
			street:       "Rua das Flores",
			number:       "123",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SP",
			country:      "",
			zipCode:      "17000-000",
			wantErr:      true,
		},
		{
			name:         "cep vazio",
			street:       "Rua das Flores",
			number:       "123",
			neighborhood: "Centro",
			city:         "Bauru",
			state:        "SP",
			zipCode:      "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := entities.NewAddress(
				tt.street,
				tt.number,
				tt.complement,
				tt.neighborhood,
				tt.city,
				tt.state,
				tt.country,
				tt.zipCode,
				tt.latitude,
				tt.longitude,
			)

			if tt.wantErr {
				assertError(t, err)
				return
			}

			assertNoError(t, err)
			assertAddress(t, address, tt)
		})
	}
}

// assertAddress valida os dados principais do endereço criado.
func assertAddress(t *testing.T, address *entities.Address, tt addressTestCase) {
	t.Helper()

	if address == nil {
		t.Fatal("address não deveria ser nil")
	}

	// ID normalmente será gerado pelo banco/GORM

	if address.Street != strings.TrimSpace(tt.street) {
		t.Errorf(
			"street inválida: got %s, want %s",
			address.Street,
			strings.TrimSpace(tt.street),
		)
	}

	if address.Number != strings.TrimSpace(tt.number) {
		t.Errorf(
			"number inválido: got %s, want %s",
			address.Number,
			strings.TrimSpace(tt.number),
		)
	}

	if address.Complement != strings.TrimSpace(tt.complement) {
		t.Errorf(
			"complement inválido: got %s, want %s",
			address.Complement,
			strings.TrimSpace(tt.complement),
		)
	}

	if address.Neighborhood != strings.TrimSpace(tt.neighborhood) {
		t.Errorf(
			"neighborhood inválido: got %s, want %s",
			address.Neighborhood,
			strings.TrimSpace(tt.neighborhood),
		)
	}

	if address.City != strings.TrimSpace(tt.city) {
		t.Errorf(
			"city inválida: got %s, want %s",
			address.City,
			strings.TrimSpace(tt.city),
		)
	}

	wantState := strings.ToUpper(strings.TrimSpace(tt.state))

	if address.State != wantState {
		t.Errorf(
			"state inválido: got %s, want %s",
			address.State,
			wantState,
		)
	}

	wantZipCode := addressValidators.NormalizeZipCode(tt.country, tt.zipCode)

	if address.ZipCode != wantZipCode {
		t.Errorf(
			"zipCode inválido: got %s, want %s",
			address.ZipCode,
			wantZipCode,
		)
	}

	if address.Latitude != tt.latitude {
		t.Errorf(
			"latitude inválida: got %f, want %f",
			address.Latitude,
			tt.latitude,
		)
	}

	if address.Longitude != tt.longitude {
		t.Errorf(
			"longitude inválida: got %f, want %f",
			address.Longitude,
			tt.longitude,
		)
	}

	if address.CreatedAt.IsZero() {
		t.Error("CreatedAt não deveria ser zero")
	}

	if address.UpdatedAt.IsZero() {
		t.Error("UpdatedAt não deveria ser zero")
	}
}
