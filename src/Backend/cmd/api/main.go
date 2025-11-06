package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Turgho/Aluguei/internal/config"
	"github.com/Turgho/Aluguei/internal/database"
	"github.com/Turgho/Aluguei/internal/server"
	"github.com/Turgho/Aluguei/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Carregar configurações
	cfg := config.Load()

	// Inicializar logger
	if err := logger.Init(logger.Config{
		Environment: cfg.Environment,
		Level:       cfg.LogLevel,
		FilePath:    cfg.LogFilePath,
		MaxSize:     100, // MB
		MaxBackups:  3,
		MaxAge:      30, // dias
	}); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Get().Sync()

	logger.Info("Starting application",
		zap.String("name", "Aluguei API"),
		zap.String("version", "1.0.0"),
		zap.String("environment", cfg.Environment),
	)

	// Conectar ao banco - CORREÇÃO: passar cfg em vez de cfg.DatabaseURL
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database",
			zap.String("url", maskDBURL(cfg.DatabaseURL)), // Mascarar senha
			zap.Error(err),
		)
	}
	defer db.Close() // Fechar conexão ao final

	logger.Info("Database connected successfully")

	// Health check do banco
	if err := db.HealthCheck(); err != nil {
		logger.Fatal("Database health check failed", zap.Error(err))
	}
	logger.Info("Database health check passed")

	// Inicializar servidor
	srv := server.New(cfg, db.DB)

	// Context para graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Iniciar servidor em goroutine
	go func() {
		logger.Info("Starting HTTP server", zap.String("port", cfg.Port))
		if err := srv.Start(); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Aguardar sinal de shutdown
	<-ctx.Done()
	logger.Info("Shutdown signal received")

	// Shutdown graceful com timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info("Shutting down server gracefully...")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("Error during server shutdown", zap.Error(err))
	} else {
		logger.Info("Server stopped gracefully")
	}
}

// maskDBURL mascara a senha na URL do banco para logs
func maskDBURL(url string) string {
	// Implementação simples - em produção use uma lib adequada
	// Isso é apenas para não expor credenciais nos logs
	return "postgresql://***:***@***/***"
}
