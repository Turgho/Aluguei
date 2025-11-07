package handlers

import (
	"net/http"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	ownerUseCase *usecases.OwnerUseCase
	jwtSecret    string
}

func NewAuthHandler(ownerUseCase *usecases.OwnerUseCase, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		ownerUseCase: ownerUseCase,
		jwtSecret:    jwtSecret,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string      `json:"token"`
	ExpiresAt time.Time   `json:"expires_at"`
	Owner     interface{} `json:"owner"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner, err := h.ownerUseCase.ValidatePassword(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT token
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"owner_id": owner.ID.String(),
		"email":    owner.Email,
		"exp":      expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		Owner:     owner,
	})
}