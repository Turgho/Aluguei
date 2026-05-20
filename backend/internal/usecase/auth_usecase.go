package usecase

import (
	"fmt"

	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/hash"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/google/uuid"
)

type authUseCase struct {
	userUC usecases.UserUseCase
}

func NewAuthUseCase(userUC usecases.UserUseCase) usecases.AuthUseCase {
	return &authUseCase{userUC: userUC}
}

// Login envia token de access_token e refresh_token para autenticação.
func (uc *authUseCase) Login(email, password string) (string, string, error) {
	user, err := uc.userUC.GetByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("credenciais inválidas")
	}

	match, err := hash.VerifyPassword(password, user.PasswordHash)
	if err != nil {
		return "", "", fmt.Errorf("erro ao verificar senha: %w", err)
	}
	if !match {
		return "", "", fmt.Errorf("credenciais inválidas")
	}

	accessToken, err := jwt.GenerateAccessToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		return "", "", fmt.Errorf("erro ao gerar access token: %w", err)
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return "", "", fmt.Errorf("erro ao gerar refresh token: %w", err)
	}

	_ = uc.userUC.RecordLogin(user.ID)

	return accessToken, refreshToken, nil
}

// RefreshToken envia um novo token de acesso para usuário.
func (uc *authUseCase) RefreshToken(refreshToken string) (string, error) { // ← authUseCase
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("refresh token inválido: %w", err)
	}

	parsedID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return "", fmt.Errorf("ID inválido no token: %w", err)
	}

	user, err := uc.userUC.GetByID(parsedID) // ← userUC, não repo
	if err != nil {
		return "", fmt.Errorf("usuário não encontrado: %w", err)
	}

	return jwt.GenerateAccessToken(user.ID.String(), user.Email, string(user.Role))
}
