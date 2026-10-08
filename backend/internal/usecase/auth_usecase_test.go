// internal/usecase/auth_usecase_test.go
package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/usecase"
	"github.com/Turgho/Aluguei/pkg/hash"
	jwtutil "github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── Login ─────────────────────────────────────────────────────────────────────

// TestLogin valida o fluxo completo de autenticação:
// busca usuário, valida senha e gera tokens.
func TestLogin(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "segredo-access-para-testes-minimo-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "segredo-refresh-para-testes-minimo-32-chars")

	// Gera hash real usando Argon2id
	passwordHash, err := hash.HashPassword("Senha@123")
	require.NoError(t, err)

	user := &entities.User{
		ID:           uuid.New(),
		Email:        "joao@email.com",
		PasswordHash: passwordHash,
		Role:         entities.RoleTenant,
	}

	repo := &mockUserRepository{
		getByEmailFn: func(email string) (*entities.User, error) {
			return user, nil
		},
		updateLastLoginFn: func(id uuid.UUID, at time.Time) error {
			return nil
		},
	}

	userUC := usecase.NewUserUseCase(repo)
	authUC := usecase.NewAuthUseCase(userUC)

	accessToken, refreshToken, err := authUC.Login(
		"joao@email.com",
		"Senha@123",
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
}

// TestLoginInvalidPassword garante que senha inválida
// bloqueia autenticação.
func TestLoginInvalidPassword(t *testing.T) {
	// Hash válido usando Argon2id
	passwordHash, err := hash.HashPassword("Senha@123")
	require.NoError(t, err)

	user := &entities.User{
		ID:           uuid.New(),
		Email:        "joao@email.com",
		PasswordHash: passwordHash,
		Role:         entities.RoleTenant,
	}

	repo := &mockUserRepository{
		getByEmailFn: func(email string) (*entities.User, error) {
			return user, nil
		},
	}

	userUC := usecase.NewUserUseCase(repo)
	authUC := usecase.NewAuthUseCase(userUC)

	_, _, err = authUC.Login(
		"joao@email.com",
		"senha-errada",
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "credenciais inválidas")
}

// ── Refresh Token ──────────────────────────────────────────────────────────────

// TestRefreshToken gera um novo access token
// usando um refresh token válido.
func TestRefreshToken(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "segredo-access-para-testes-minimo-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "segredo-refresh-para-testes-minimo-32-chars")

	user := &entities.User{
		ID:    uuid.New(),
		Email: "joao@email.com",
		Role:  entities.RoleTenant,
	}

	refreshToken, err := jwtutil.GenerateRefreshToken(user.ID.String())
	require.NoError(t, err)

	repo := &mockUserRepository{
		getByIDFn: func(id uuid.UUID) (*entities.User, error) {
			assert.Equal(t, user.ID, id)
			return user, nil
		},
	}

	userUC := usecase.NewUserUseCase(repo)
	authUC := usecase.NewAuthUseCase(userUC)

	newAccessToken, newRefreshToken, err := authUC.RefreshToken(refreshToken)
	require.NoError(t, err)
	assert.NotEmpty(t, newAccessToken)
	assert.NotEmpty(t, newRefreshToken)
}

// TestRefreshTokenInvalid garante que refresh token inválido
// retorna erro.
func TestRefreshTokenInvalid(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "segredo-access-para-testes-minimo-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "segredo-refresh-para-testes-minimo-32-chars")

	repo := &mockUserRepository{}
	userUC := usecase.NewUserUseCase(repo)
	authUC := usecase.NewAuthUseCase(userUC)

	newAccessToken, _, err := authUC.RefreshToken("token-invalido")
	require.Error(t, err)
	assert.Empty(t, newAccessToken)
	assert.Contains(t, err.Error(), "refresh token inválido")
}

// TestRefreshTokenUserNotFound garante erro
// quando o usuário do token não existe.
func TestRefreshTokenUserNotFound(t *testing.T) {
	t.Setenv("JWT_ACCESS_SECRET", "segredo-access-para-testes-minimo-32-chars")
	t.Setenv("JWT_REFRESH_SECRET", "segredo-refresh-para-testes-minimo-32-chars")

	userID := uuid.New()

	refreshToken, err := jwtutil.GenerateRefreshToken(userID.String())
	require.NoError(t, err)

	repo := &mockUserRepository{
		getByIDFn: func(id uuid.UUID) (*entities.User, error) {
			return nil, errors.New("not found")
		},
	}

	userUC := usecase.NewUserUseCase(repo)
	authUC := usecase.NewAuthUseCase(userUC)

	newAccessToken, _, err := authUC.RefreshToken(refreshToken)
	require.Error(t, err)
	assert.Empty(t, newAccessToken)
	assert.Contains(t, err.Error(), "usuário não encontrado")
}
