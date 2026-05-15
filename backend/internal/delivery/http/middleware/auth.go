// Package middleware fornece middlewares HTTP para uso com o framework [github.com/gin-gonic/gin].
package middleware

import (
	"net/http"
	"strings"

	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// Auth retorna um middleware Gin que protege rotas com autenticação via JWT.
//
// O token deve ser enviado no header Authorization no formato:
//
//	Authorization: Bearer <token>
//
// Em caso de token ausente, malformado ou expirado, a requisição é abortada
// com status 401 Unauthorized.
//
// Após validação bem-sucedida, os seguintes valores são injetados no contexto
// e ficam disponíveis nos handlers subsequentes via [gin.Context.Get]:
//
//   - "userID" (uint)   — identificador do usuário autenticado
//   - "email"  (string) — e-mail do usuário autenticado
//   - "role"   (string) — papel/permissão do usuário autenticado
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
