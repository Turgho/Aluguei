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
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// getSecretKey retorna a chave JWT carregada da variável de ambiente.
func getSecretKey() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET não configurado")
	}

	return []byte(secret), nil
}

// GenerateToken gera um token JWT assinado com HS256 contendo as informações
// do usuário autenticado.
//
// O token gerado expira em 24 horas a partir do momento da criação.
//
// Retorna o token assinado como string ou um erro caso a assinatura falhe.
func GenerateToken(userID string, email, role string) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "aluguei-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar token: %w", err)
	}

	return signed, nil
}

// ValidateToken valida um token JWT e retorna as [Claims] contidas nele.
//
// A validação garante que:
//   - O token foi assinado com o algoritmo HS256 esperado
//   - A assinatura é válida e corresponde à chave secreta
//   - O token não está expirado
//
// Retorna erro nos seguintes casos:
//   - Token malformado ou vazio
//   - Algoritmo de assinatura diferente de HS256
//   - Assinatura inválida
//   - Token expirado
func ValidateToken(tokenStr string) (*Claims, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			// Rejeita qualquer algoritmo diferente de HS256
			if t.Method != jwt.SigningMethodHS256 {
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

	return claims, nil
}
