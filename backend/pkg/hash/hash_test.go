// pkg/hash/hash_test.go
package hash_test

import (
	"testing"

	"github.com/Turgho/Aluguei/pkg/hash"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "senha válida",
			password: "Senha123!",
		},
		{
			name:     "senha longa",
			password: "MinhaSenhaMuitoForte123@",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := hash.HashPassword(tt.password)

			if err != nil {
				t.Fatalf("esperava nil error, recebeu %v", err)
			}

			if result == "" {
				t.Fatal("hash não deveria ser vazio")
			}

			if result == tt.password {
				t.Fatal("hash não deveria ser igual à senha")
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		wrongPass   string
		expectMatch bool
	}{
		{
			name:        "senha correta",
			password:    "Senha123!",
			wrongPass:   "Errada123!",
			expectMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := hash.HashPassword(tt.password)
			if err != nil {
				t.Fatalf("erro ao gerar hash: %v", err)
			}

			match, err := hash.VerifyPassword(tt.password, encoded)
			if err != nil {
				t.Fatalf("erro ao verificar hash: %v", err)
			}

			if match != tt.expectMatch {
				t.Errorf("esperado %v, recebido %v", tt.expectMatch, match)
			}

			wrongMatch, err := hash.VerifyPassword(tt.wrongPass, encoded)
			if err != nil {
				t.Fatalf("erro ao verificar senha inválida: %v", err)
			}

			if wrongMatch {
				t.Fatal("senha inválida não deveria bater")
			}
		})
	}
}

func TestHashPasswordGeneratesDifferentHashes(t *testing.T) {
	password := "Senha123!"

	hash1, err := hash.HashPassword(password)
	if err != nil {
		t.Fatalf("erro ao gerar hash1: %v", err)
	}

	hash2, err := hash.HashPassword(password)
	if err != nil {
		t.Fatalf("erro ao gerar hash2: %v", err)
	}

	if hash1 == hash2 {
		t.Fatal("hashes deveriam ser diferentes")
	}
}

func TestVerifyPasswordInvalidHash(t *testing.T) {
	match, err := hash.VerifyPassword("Senha123!", "hash-invalido")

	if err == nil {
		t.Fatal("esperava erro para hash inválido")
	}

	if match {
		t.Fatal("match deveria ser false")
	}
}
