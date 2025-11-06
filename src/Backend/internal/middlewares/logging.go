package middlewares

import (
	"runtime"
	"time"

	"github.com/Turgho/Aluguei/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggingMiddleware cria um middleware de logging para requests HTTP
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Processar request
		c.Next()

		latency := time.Since(start)

		if len(c.Errors) > 0 {
			// Log de erros
			for _, e := range c.Errors.Errors() {
				logger.Error("HTTP Request Error",
					zap.String("path", path),
					zap.String("query", query),
					zap.String("method", c.Request.Method),
					zap.Int("status", c.Writer.Status()),
					zap.String("ip", c.ClientIP()),
					zap.String("user_agent", c.Request.UserAgent()),
					zap.Duration("latency", latency),
					zap.String("error", e),
				)
			}
		} else {
			// Log de request bem-sucedida
			logger.Info("HTTP Request",
				zap.String("path", path),
				zap.String("query", query),
				zap.String("method", c.Request.Method),
				zap.Int("status", c.Writer.Status()),
				zap.String("ip", c.ClientIP()),
				zap.String("user_agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			)
		}
	}
}

// RecoveryMiddleware com logging para panics
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("HTTP Panic Recovered",
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("ip", c.ClientIP()),
					zap.Any("panic", err),
					zap.ByteString("stack", getStack()),
				)

				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

func getStack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
