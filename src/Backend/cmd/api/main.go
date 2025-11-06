package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"my-api/internal/config"
	"my-api/internal/database"
	"my-api/internal/server"
	"my-api/pkg/logger"

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

	// Conectar ao banco
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database",
			zap.String("url", cfg.DatabaseURL),
			zap.Error(err),
		)
	}
	logger.Info("Database connected successfully")

	// Inicializar servidor
	srv := server.New(cfg, db)

	// Graceful shutdown
	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Aguardar sinal de shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	if err := srv.Shutdown(); err != nil {
		logger.Error("Error during server shutdown", zap.Error(err))
	}
	logger.Info("Server stopped")
}
