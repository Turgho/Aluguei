// internal/domain/entities/user_test.go
package entities_test

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		firstName   string
		lastName    string
		cpf         string
		email       string
		phone       string
		password    string
		role        entities.Role
		expectError bool
	}{
		{
			name:        "usuário válido owner",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "529.982.247-25",
			email:       "victor@email.com",
			phone:       "(14) 99999-9999",
			password:    "hash-da-senha",
			role:        entities.RoleOwner,
			expectError: false,
		},
		{
			name:        "usuário válido tenant",
			firstName:   "Maria",
			lastName:    "Silva",
			cpf:         "52998224725",
			email:       "maria@email.com",
			phone:       "",
			password:    "hash-da-senha",
			role:        entities.RoleTenant,
			expectError: false,
		},
		{
			name:        "nome vazio",
			firstName:   "",
			lastName:    "Silva",
			cpf:         "52998224725",
			email:       "teste@email.com",
			phone:       "",
			password:    "hash-da-senha",
			role:        entities.RoleTenant,
			expectError: true,
		},
		{
			name:        "cpf inválido",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "11111111111",
			email:       "teste@email.com",
			phone:       "",
			password:    "hash-da-senha",
			role:        entities.RoleTenant,
			expectError: true,
		},
		{
			name:        "email inválido",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "52998224725",
			email:       "email-invalido",
			phone:       "",
			password:    "hash-da-senha",
			role:        entities.RoleTenant,
			expectError: true,
		},
		{
			name:        "telefone inválido",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "52998224725",
			email:       "teste@email.com",
			phone:       "123",
			password:    "hash-da-senha",
			role:        entities.RoleTenant,
			expectError: true,
		},
		{
			name:        "senha vazia",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "52998224725",
			email:       "teste@email.com",
			phone:       "",
			password:    "",
			role:        entities.RoleTenant,
			expectError: true,
		},
		{
			name:        "role inválida",
			firstName:   "Victor",
			lastName:    "Hugo",
			cpf:         "52998224725",
			email:       "teste@email.com",
			phone:       "",
			password:    "hash-da-senha",
			role:        "admin",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entities.NewUser(
				tt.firstName,
				tt.lastName,
				tt.cpf,
				tt.email,
				tt.phone,
				tt.password,
				tt.role,
			)

			if tt.expectError {
				if err == nil {
					t.Fatal("esperava erro, recebeu nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("esperava nil error, recebeu %v", err)
			}

			if user == nil {
				t.Fatal("user não deveria ser nil")
			}

			// CPF normalizado
			if user.CPF != "52998224725" && tt.cpf != "" {
				t.Errorf("cpf não normalizado: %s", user.CPF)
			}

			// Telefone normalizado
			if tt.phone == "(14) 99999-9999" {
				if user.Phone != "5514999999999" {
					t.Errorf("telefone não normalizado: %s", user.Phone)
				}
			}

			if user.CreatedAt.IsZero() {
				t.Error("CreatedAt não deveria ser zero")
			}

			if user.UpdatedAt.IsZero() {
				t.Error("UpdatedAt não deveria ser zero")
			}
		})
	}
}
