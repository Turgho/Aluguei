// pkg/jwt/jwt_test.go
package jwt_test

import (
	"os"
	"testing"

	"github.com/Turgho/Aluguei/pkg/jwt"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "super-secret-key-123")

	token, err := jwt.GenerateToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)

	if err != nil {
		t.Fatalf("esperava nil error, recebeu %v", err)
	}

	if token == "" {
		t.Fatal("token não deveria ser vazio")
	}
}

func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "super-secret-key-123")

	token, err := jwt.GenerateToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	claims, err := jwt.ValidateToken(token)
	if err != nil {
		t.Fatalf("erro ao validar token: %v", err)
	}

	if claims.UserID != "user-123" {
		t.Errorf("esperado user-123, recebido %s", claims.UserID)
	}

	if claims.Email != "teste@email.com" {
		t.Errorf("esperado teste@email.com, recebido %s", claims.Email)
	}

	if claims.Role != "tenant" {
		t.Errorf("esperado tenant, recebido %s", claims.Role)
	}
}

func TestValidateTokenInvalid(t *testing.T) {
	os.Setenv("JWT_SECRET", "super-secret-key-123")

	_, err := jwt.ValidateToken("token-invalido")

	if err == nil {
		t.Fatal("esperava erro para token inválido")
	}
}

func TestValidateTokenWrongSecret(t *testing.T) {
	os.Setenv("JWT_SECRET", "secret-1")

	token, err := jwt.GenerateToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	os.Setenv("JWT_SECRET", "secret-2")

	_, err = jwt.ValidateToken(token)

	if err == nil {
		t.Fatal("esperava erro com secret diferente")
	}
}

func TestGenerateTokenWithoutSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := jwt.GenerateToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)

	if err == nil {
		t.Fatal("esperava erro sem JWT_SECRET")
	}
}
