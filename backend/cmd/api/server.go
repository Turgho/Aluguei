// Package main é o ponto de entrada da aplicação Aluguei.
// Responsável por inicializar dependências e subir o servidor HTTP.
package main

import (
	"log"
	"os"
	"time"

	_ "github.com/Turgho/Aluguei/docs"
	"github.com/Turgho/Aluguei/internal/delivery/http/handlers"
	"github.com/Turgho/Aluguei/internal/delivery/http/middleware"
	"github.com/Turgho/Aluguei/internal/infra/database"
	"github.com/Turgho/Aluguei/internal/infra/repositories"
	userUseCase "github.com/Turgho/Aluguei/internal/usecase"
	"github.com/Turgho/Aluguei/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Server representa a instância principal do servidor HTTP.
// Encapsula o roteador Gin e a conexão com o banco de dados.
type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

// NewServer inicializa todas as dependências da aplicação e retorna
// uma instância configurada de [Server].
//
// A inicialização segue a ordem:
//  1. Carrega variáveis de ambiente via .env
//  2. Inicializa o logger (zap)
//  3. Configura o modo do Gin (debug/release)
//  4. Conecta ao banco de dados
//  5. Configura o roteador e os middlewares
//  6. Registra as rotas
//
// Encerra a aplicação com log.Fatal em caso de erro crítico
// (arquivo .env ausente ou falha na conexão com o banco).
func NewServer() *Server {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Init(os.Getenv("APP_ENV"))

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log.Fatal("falha ao conectar no banco")
	}

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())
	router.Use(middleware.ZapLogger(logger.Log))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// APENAS PARA DEV
			"http://localhost:4200",
			"http://127.0.0.1:4200",
			"http://192.168.1.7:4200", // PC
			"http://192.168.1.6:4200", // CELULAR
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server := &Server{router: router, db: db}
	server.setupRoutes()

	return server
}

// Run inicia o servidor HTTP na porta definida pela variável de ambiente APP_PORT.
// Caso APP_PORT não esteja definida, utiliza a porta padrão 3000.
func (s *Server) Run() {
	defer logger.Log.Sync()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log.Info("servidor rodando", zap.String("port", port))
	s.router.Run("0.0.0.0:" + port)
}

// setupRoutes registra todas as rotas da aplicação, organizadas em dois grupos:
//
//   - Rotas públicas (/api/v1): acessíveis sem autenticação (login, registro)
//   - Rotas privadas (/api/v1): protegidas pelo middleware [middleware.Auth]
func (s *Server) setupRoutes() {
	// Users
	userRepo := repositories.NewUserRepository(s.db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userH := handlers.NewUserHandler(userUC)
	authH := handlers.NewAuthHandler(userUC)

	// Swagger Docs
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := s.router.Group("/api/v1/auth")
	{
		auth.POST("/login", authH.Login)
		auth.POST("/logout", authH.Logout)
		auth.POST("/refresh", authH.RefreshToken)
		auth.POST("/register", authH.Register)
		auth.GET("/me", middleware.Auth(), authH.Me)
	}

	private := s.router.Group("/api/v1")
	private.Use(middleware.Auth())
	{
		private.GET("/users/search", userH.Search)
		private.GET("/users/:id", userH.GetByID)
		private.PUT("/users/:id", userH.Update)
		private.DELETE("/users/:id", userH.Delete)
	}
}
