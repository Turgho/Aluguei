// internal/delivery/http/middleware/zap_logger_test.go
package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

// newObservedLogger retorna um logger zap com observer para inspecionar os logs nos testes.
func newObservedLogger() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zap.DebugLevel)
	return zap.New(core), logs
}

// newLoggerRouter cria um router Gin com ZapLogger e um handler que retorna o status informado.
func newLoggerRouter(log *zap.Logger, status int) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ZapLogger(log))
	r.GET("/test", func(c *gin.Context) {
		c.Status(status)
	})
	return r
}

func TestZapLogger(t *testing.T) {
	t.Run("loga requisição 2xx como Info", func(t *testing.T) {
		log, logs := newObservedLogger()
		r := newLoggerRouter(log, http.StatusOK)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 1, logs.Len())
		assert.Equal(t, zap.InfoLevel, logs.All()[0].Level)
		assert.Equal(t, "request", logs.All()[0].Message)
	})

	t.Run("loga requisição 4xx como Warn", func(t *testing.T) {
		log, logs := newObservedLogger()
		r := newLoggerRouter(log, http.StatusBadRequest)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 1, logs.Len())
		assert.Equal(t, zap.WarnLevel, logs.All()[0].Level)
		assert.Equal(t, "client error", logs.All()[0].Message)
	})

	t.Run("loga requisição 5xx como Error", func(t *testing.T) {
		log, logs := newObservedLogger()
		r := newLoggerRouter(log, http.StatusInternalServerError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 1, logs.Len())
		assert.Equal(t, zap.ErrorLevel, logs.All()[0].Level)
		assert.Equal(t, "server error", logs.All()[0].Message)
	})

	t.Run("campos obrigatórios presentes no log", func(t *testing.T) {
		log, logs := newObservedLogger()
		r := newLoggerRouter(log, http.StatusOK)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		fields := logs.All()[0].ContextMap()
		assert.Contains(t, fields, "method")
		assert.Contains(t, fields, "path")
		assert.Contains(t, fields, "status")
		assert.Contains(t, fields, "latency")
		assert.Contains(t, fields, "ip")
	})
}
