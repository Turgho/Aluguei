// pkg/jwt/jwt_test.go
package jwt_test

import (
	"os"
	"testing"

	"github.com/Turgho/Aluguei/pkg/jwt"
)

func setupSecrets(t *testing.T) {
	t.Helper()

	t.Setenv("JWT_ACCESS_SECRET", "test-access-secret-key-with-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "test-refresh-secret-key-with-32ch")
}

func TestGenerateAccessToken(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateAccessToken(
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

func TestValidateAccessToken(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateAccessToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	claims, err := jwt.ValidateAccessToken(token)
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

	if claims.Type != "access" {
		t.Errorf("esperado access, recebido %s", claims.Type)
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateRefreshToken("user-123")
	if err != nil {
		t.Fatalf("erro ao gerar refresh token: %v", err)
	}

	if token == "" {
		t.Fatal("refresh token não deveria ser vazio")
	}
}

func TestValidateRefreshToken(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateRefreshToken("user-123")
	if err != nil {
		t.Fatalf("erro ao gerar refresh token: %v", err)
	}

	claims, err := jwt.ValidateRefreshToken(token)
	if err != nil {
		t.Fatalf("erro ao validar refresh token: %v", err)
	}

	if claims.UserID != "user-123" {
		t.Errorf("esperado user-123, recebido %s", claims.UserID)
	}

	if claims.Type != "refresh" {
		t.Errorf("esperado refresh, recebido %s", claims.Type)
	}
}

func TestValidateAccessTokenInvalid(t *testing.T) {
	setupSecrets(t)

	_, err := jwt.ValidateAccessToken("token-invalido")

	if err == nil {
		t.Fatal("esperava erro para token inválido")
	}
}

func TestValidateRefreshTokenInvalid(t *testing.T) {
	setupSecrets(t)

	_, err := jwt.ValidateRefreshToken("token-invalido")

	if err == nil {
		t.Fatal("esperava erro para refresh token inválido")
	}
}

func TestValidateAccessTokenWrongSecret(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateAccessToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	os.Setenv("JWT_ACCESS_SECRET", "wrong-secret-with-32-characters")

	_, err = jwt.ValidateAccessToken(token)

	if err == nil {
		t.Fatal("esperava erro com secret diferente")
	}
}

func TestGenerateAccessTokenWithoutSecret(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "")

	_, err := jwt.GenerateAccessToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)

	if err == nil {
		t.Fatal("esperava erro sem JWT_ACCESS_SECRET")
	}
}

func TestGenerateRefreshTokenWithoutSecret(t *testing.T) {
	t.Setenv("JWT_REFRESH_SECRET", "")

	_, err := jwt.GenerateRefreshToken("user-123")

	if err == nil {
		t.Fatal("esperava erro sem JWT_REFRESH_SECRET")
	}
}

func TestRefreshTokenCannotBeUsedAsAccess(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateRefreshToken("user-123")
	if err != nil {
		t.Fatalf("erro ao gerar refresh token: %v", err)
	}

	_, err = jwt.ValidateAccessToken(token)

	if err == nil {
		t.Fatal("refresh token não deveria validar como access")
	}
}

func TestAccessTokenCannotBeUsedAsRefresh(t *testing.T) {
	setupSecrets(t)

	token, err := jwt.GenerateAccessToken(
		"user-123",
		"teste@email.com",
		"tenant",
	)
	if err != nil {
		t.Fatalf("erro ao gerar access token: %v", err)
	}

	_, err = jwt.ValidateRefreshToken(token)

	if err == nil {
		t.Fatal("access token não deveria validar como refresh")
	}
}
