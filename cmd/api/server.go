package main

import (
	"log"
	"os"

	"github.com/Turgho/Aluguei/internal/infra/database"
	"github.com/Turgho/Aluguei/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

// NewServer inicializa as dependências e retorna uma instância do servidor
func NewServer() *Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Init(os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Database
	db, err := database.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log.Fatal("falha ao conectar no banco")
	}

	// New() ao invés de Default() para melhor controle de logs via zap
	// Usando apenas Recovery() por segurança
	router := gin.New()
	router.Use(gin.Recovery())

	server := &Server{router: router, db: db}
	server.setupRoutes()

	return server
}

// Run inicia o servidor HTTP na porta configurada.
func (s *Server) Run() {
	defer logger.Log.Sync()

	// Servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	logger.Log.Info("servidor rodando", zap.String("port", port))
	s.router.Run(":" + port)
}

func (s *Server) setupRoutes() {

}
