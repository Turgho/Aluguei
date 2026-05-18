// Package jwt fornece utilitários para geração e validação de tokens JWT
// utilizando o algoritmo HMAC-SHA256 (HS256).
//
// O token gerado contém as claims do usuário (ID, e-mail e papel) além das
// claims padrão do JWT (expiração, emissão e emissor).
package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims representa o payload do token JWT.
//
// Estende [jwt.RegisteredClaims] com informações específicas da aplicação:
//   - UserID: identificador único do usuário
//   - Email: endereço de e-mail do usuário
//   - Role: papel/permissão do usuário (ex: "admin", "user")
//   - Type: tipo do token ("access" ou "refresh")
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email,omitempty"`
	Role   string `json:"role,omitempty"`
	Type   string `json:"type"`

	jwt.RegisteredClaims
}

// getSecretKey retorna a chave JWT carregada da variável de ambiente.
func getSecretKey(env string) ([]byte, error) {
	secret := os.Getenv(env)

	if secret == "" {
		return nil, fmt.Errorf("%s não configurado", env)
	}

	if len(secret) < 32 {
		return nil, fmt.Errorf("%s deve ter no mínimo 32 caracteres", env)
	}

	return []byte(secret), nil
}

// GenerateAccessToken gera um access token JWT com duração curta.
//
// O token contém:
//   - user_id
//   - email
//   - role
//
// O token expira em 15 minutos.
func GenerateAccessToken(userID string, email, role string) (string, error) {
	secretKey, err := getSecretKey("JWT_ACCESS_SECRET")
	if err != nil {
		return "", err
	}

	now := time.Now()

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "aluguei-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar access token: %w", err)
	}

	return signed, nil
}

// GenerateRefreshToken gera um refresh token JWT.
//
// O refresh token contém apenas:
//   - user_id
//   - type
//
// O token expira em 7 dias.
func GenerateRefreshToken(userID string) (string, error) {
	secretKey, err := getSecretKey("JWT_REFRESH_SECRET")
	if err != nil {
		return "", err
	}

	now := time.Now()

	claims := Claims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "aluguei-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar refresh token: %w", err)
	}

	return signed, nil
}

// ValidateAccessToken valida um access token JWT.
func ValidateAccessToken(tokenStr string) (*Claims, error) {
	return validateToken(tokenStr, "JWT_ACCESS_SECRET", "access")
}

// ValidateRefreshToken valida um refresh token JWT.
func ValidateRefreshToken(tokenStr string) (*Claims, error) {
	return validateToken(tokenStr, "JWT_REFRESH_SECRET", "refresh")
}

// validateToken valida assinatura, expiração e tipo do token.
func validateToken(tokenStr, envSecret, expectedType string) (*Claims, error) {
	secretKey, err := getSecretKey(envSecret)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("algoritmo inesperado: %v", t.Header["alg"])
			}

			return secretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("erro ao validar token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	if claims.Type != expectedType {
		return nil, fmt.Errorf("tipo de token inválido")
	}

	return claims, nil
}
