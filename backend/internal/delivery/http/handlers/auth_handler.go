// internal/delivery/http/handlers/auth_handler.go
package handlers

import (
	"net/http"

	"github.com/Turgho/Aluguei/internal/domain/entities"
	"github.com/Turgho/Aluguei/internal/domain/usecases"
	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/Turgho/Aluguei/pkg/response"
	"github.com/gin-gonic/gin"
)

// AuthHandler agrupa os handlers HTTP relacionados à autenticação.
type AuthHandler struct {
	uc usecases.UserUseCase
}

// NewAuthHandler retorna uma instância de [AuthHandler].
func NewAuthHandler(uc usecases.UserUseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

// ── Request / Response ─────────────────────────────────────────────────────

// loginRequest representa o corpo da requisição de login.
type loginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// registerRequest representa o corpo da requisição de registro.
type registerRequest struct {
	FirstName string        `json:"first_name" binding:"required"`
	LastName  string        `json:"last_name"  binding:"required"`
	CPF       string        `json:"cpf"        binding:"required"`
	Email     string        `json:"email"      binding:"required,email"`
	Phone     string        `json:"phone"`
	Password  string        `json:"password"   binding:"required,min=8"`
	Role      entities.Role `json:"role"       binding:"required"`
}

// authResponse representa a resposta genérica de autenticação.
type authResponse struct {
	Success bool `json:"success"`
}

// ── Helpers ────────────────────────────────────────────────────────────────

func setAccessCookie(c *gin.Context, token string) {
	c.SetCookieData(&http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		MaxAge:   900, // 15 min
		Secure:   c.Request.TLS != nil,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func setRefreshCookie(c *gin.Context, token string) {
	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/api/v1/auth/refresh",
		Domain:   "",
		MaxAge:   604800, // 7 dias
		Secure:   c.Request.TLS != nil,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearAuthCookies(c *gin.Context) {
	c.SetCookieData(&http.Cookie{
		Name:   "access_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	c.SetCookieData(&http.Cookie{
		Name:   "refresh_token",
		Value:  "",
		Path:   "/api/v1/auth/refresh",
		MaxAge: -1,
	})
}

// ── Handlers ───────────────────────────────────────────────────────────────

// Register godoc
//
//	@Summary		Registra um novo usuário
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		registerRequest	true	"Dados do usuário"
//	@Success		201		{object}	userResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		409		{object}	response.ErrorResponse
//	@Router			/api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	user, err := h.uc.Create(req.FirstName, req.LastName, req.CPF, req.Email, req.Phone, req.Password, req.Role)
	if err != nil {
		response.Error(c, http.StatusConflict, "CONFLICT", err.Error())
		return
	}

	c.JSON(http.StatusCreated, toUserResponse(user))
}

// Login godoc
//
//	@Summary		Autentica um usuário e define cookies JWT HttpOnly
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		loginRequest	true	"Credenciais"
//	@Success		200		{object}	authResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	accessToken, refreshToken, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "credenciais inválidas")
		return
	}

	setAccessCookie(c, accessToken)
	setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusOK, authResponse{Success: true})
}

// RefreshToken godoc
//
//	@Summary		Renova o access token
//	@Description	Valida o refresh token do cookie e emite um novo access token
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	authResponse
//	@Failure		401	{object}	response.ErrorResponse
//	@Router			/api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	tokenStr, err := c.Cookie("refresh_token")
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "MISSING_REFRESH_TOKEN", "refresh token ausente")
		return
	}

	newAccessToken, err := h.uc.RefreshToken(tokenStr)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "INVALID_REFRESH_TOKEN", "refresh token inválido ou expirado")
		return
	}

	setAccessCookie(c, newAccessToken)

	c.JSON(http.StatusOK, authResponse{Success: true})
}

// Logout godoc
//
//	@Summary		Encerra a sessão do usuário
//	@Tags			auth
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	authResponse
//	@Router			/api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	clearAuthCookies(c)
	c.JSON(http.StatusOK, authResponse{Success: true})
}

// Me godoc
//
//	@Summary		Retorna o usuário autenticado
//	@Description	Usa o access token do cookie para identificar o usuário
//	@Tags			auth
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	userResponse
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/api/v1/auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	claims := c.MustGet("user").(*jwt.Claims)

	user, err := h.uc.GetByID(claims.UserID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", "usuário não encontrado")
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}
