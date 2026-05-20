// internal/domain/entities/user_test.go
package entities_test

import (
	"testing"

	"github.com/Turgho/Aluguei/internal/domain/entities"
)

type userTestCase struct {
	name         string
	firstName    string
	lastName     string
	cpf          string
	email        string
	phone        string
	passwordHash string
	role         entities.Role

	wantErr   bool
	wantCPF   string
	wantPhone string
}

func TestNewUser(t *testing.T) {
	tests := []userTestCase{
		{
			name:         "usuário válido owner",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "529.982.247-25",
			email:        "victor@email.com",
			phone:        "(14) 99999-9999",
			passwordHash: "hash-da-senha",
			role:         entities.RoleOwner,
			wantErr:      false,
			wantCPF:      "52998224725",
			wantPhone:    "5514999999999",
		},
		{
			name:         "usuário válido tenant sem telefone",
			firstName:    "Maria",
			lastName:     "Silva",
			cpf:          "52998224725",
			email:        "maria@email.com",
			phone:        "",
			passwordHash: "hash-da-senha",
			role:         entities.RoleTenant,
			wantErr:      false,
			wantCPF:      "52998224725",
			wantPhone:    "",
		},
		{
			name:         "nome vazio",
			firstName:    "",
			lastName:     "Silva",
			cpf:          "52998224725",
			email:        "teste@email.com",
			phone:        "",
			passwordHash: "hash-da-senha",
			role:         entities.RoleTenant,
			wantErr:      true,
		},
		{
			name:         "cpf inválido",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "11111111111",
			email:        "teste@email.com",
			phone:        "",
			passwordHash: "hash-da-senha",
			role:         entities.RoleTenant,
			wantErr:      true,
		},
		{
			name:         "email inválido",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "52998224725",
			email:        "email-invalido",
			phone:        "",
			passwordHash: "hash-da-senha",
			role:         entities.RoleTenant,
			wantErr:      true,
		},
		{
			name:         "telefone inválido",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "52998224725",
			email:        "teste@email.com",
			phone:        "123",
			passwordHash: "hash-da-senha",
			role:         entities.RoleTenant,
			wantErr:      true,
		},
		{
			name:         "senha vazia",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "52998224725",
			email:        "teste@email.com",
			phone:        "",
			passwordHash: "",
			role:         entities.RoleTenant,
			wantErr:      true,
		},
		{
			name:         "role inválida",
			firstName:    "Victor",
			lastName:     "Hugo",
			cpf:          "52998224725",
			email:        "teste@email.com",
			phone:        "",
			passwordHash: "hash-da-senha",
			role:         "admin",
			wantErr:      true,
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
				tt.passwordHash,
				tt.role,
			)

			if tt.wantErr {
				assertError(t, err)
				return
			}

			assertNoError(t, err)
			assertUser(t, user, tt)
		})
	}
}

// assertUser valida os dados principais do usuário criado.
func assertUser(t *testing.T, user *entities.User, tt userTestCase) {
	t.Helper()

	if user == nil {
		t.Fatal("user não deveria ser nil")
	}

	if user.FirstName != tt.firstName {
		t.Errorf("firstName inválido: got %s, want %s", user.FirstName, tt.firstName)
	}

	if user.LastName != tt.lastName {
		t.Errorf("lastName inválido: got %s, want %s", user.LastName, tt.lastName)
	}

	if user.CPF != tt.wantCPF {
		t.Errorf("cpf inválido: got %s, want %s", user.CPF, tt.wantCPF)
	}

	if user.Email != tt.email {
		t.Errorf("email inválido: got %s, want %s", user.Email, tt.email)
	}

	if user.Phone != tt.wantPhone {
		t.Errorf("telefone inválido: got %s, want %s", user.Phone, tt.wantPhone)
	}

	if user.Role != tt.role {
		t.Errorf("role inválida: got %s, want %s", user.Role, tt.role)
	}

	if !user.IsActive {
		t.Error("IsActive deveria ser true")
	}

	if user.EmailVerified {
		t.Error("EmailVerified deveria ser false")
	}

	if user.CreatedAt.IsZero() {
		t.Error("CreatedAt não deveria ser zero")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("UpdatedAt não deveria ser zero")
	}
}
