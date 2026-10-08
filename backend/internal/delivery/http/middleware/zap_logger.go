// Package middleware fornece middlewares HTTP para uso com o framework [github.com/gin-gonic/gin].
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger retorna um middleware Gin que registra informações de cada requisição HTTP
// utilizando o logger estruturado [go.uber.org/zap].
//
// Para cada requisição são registrados:
//   - Método HTTP (GET, POST, etc.)
//   - Caminho da URL
//   - Código de status da resposta
//   - Latência total de processamento
//   - IP do cliente
//   - Erros ocorridos durante o processamento (se houver)
//
// Deve ser registrado antes dos handlers no roteador:
//
//	router.Use(middleware.ZapLogger(logger.Log))
func ZapLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Processa a requisição antes de logar
		c.Next()

		// Campos base de toda requisição
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("ip", c.ClientIP()),
		}

		// Anexa erros do Gin ao log, se existirem
		if errs := c.Errors.ByType(gin.ErrorTypePrivate); len(errs) > 0 {
			fields = append(fields, zap.String("errors", errs.String()))
		}

		// Nível de log varia de acordo com o status HTTP
		switch {
		case c.Writer.Status() >= 500:
			log.Error("server error", fields...)
		case c.Writer.Status() >= 400:
			log.Warn("client error", fields...)
		default:
			log.Info("request", fields...)
		}
	}
}
