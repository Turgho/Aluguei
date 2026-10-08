// Package middleware fornece middlewares HTTP para uso com o framework [github.com/gin-gonic/gin].
package middleware

import (
	"net/http"

	"github.com/Turgho/Aluguei/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// Auth retorna um middleware Gin que protege rotas com autenticação via JWT.
//
// O access token JWT deve ser enviado via cookie HttpOnly:
//
//	access_token=<jwt>
//
// Em caso de token ausente, inválido ou expirado, a requisição é abortada
// com status 401 Unauthorized.
//
// Após validação bem-sucedida, os seguintes valores são injetados no contexto:
//
//   - "userID" (string) — identificador do usuário autenticado
//   - "email"  (string) — e-mail do usuário autenticado
//   - "role"   (string) — papel/permissão do usuário autenticado
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token ausente",
			})
			return
		}

		claims, err := jwt.ValidateAccessToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
