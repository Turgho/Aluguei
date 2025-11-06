package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Turgho/Aluguei/internal/config"
	"github.com/Turgho/Aluguei/internal/middlewares"
	"github.com/Turgho/Aluguei/internal/server/handlers"
	"github.com/Turgho/Aluguei/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
	server *http.Server
}

func New(cfg *config.Config, db *gorm.DB) *Server {
	// Configurar modo do Gin baseado no ambiente
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// Middlewares globais
	router.Use(gin.Recovery())
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.CORSMiddleware())

	s := &Server{
		config: cfg,
		db:     db,
		router: router,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.healthHandler)
	s.router.GET("/ready", s.readyHandler)

	// Swagger com YAML customizado
	s.router.GET("/swagger.yaml", s.serveSwaggerYAML)
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger.yaml")))

	// API v1
	api := s.router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.registerHandler)
			auth.POST("/login", s.loginHandler)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			// Users
			protected.GET("/users/me", s.getCurrentUserHandler)

			// Owners
			owners := protected.Group("/owners")
			{
				owners.GET("", handlers.GetOwnersHandler)
				owners.POST("", handlers.CreateOwnerHandler)
				owners.GET("/:id", handlers.GetOwnerHandler)
				owners.PUT("/:id", handlers.UpdateOwnerHandler)
				owners.DELETE("/:id", handlers.DeleteOwnerHandler)
				owners.GET("/email/:email", handlers.GetOwnerByEmailHandler)
			}

			// Properties
			properties := protected.Group("/properties")
			{
				properties.GET("", handlers.GetPropertiesHandler)
				properties.POST("", handlers.CreatePropertyHandler)
				properties.GET("/:id", handlers.GetPropertyHandler)
				properties.PUT("/:id", handlers.UpdatePropertyHandler)
				properties.DELETE("/:id", handlers.DeletePropertyHandler)
				properties.GET("/owner/:ownerId", handlers.GetPropertiesByOwnerHandler)
			}

			// Tenants
			tenants := protected.Group("/tenants")
			{
				tenants.GET("", handlers.GetTenantsHandler)
				tenants.POST("", handlers.CreateTenantHandler)
				tenants.GET("/:id", handlers.GetTenantHandler)
				tenants.PUT("/:id", handlers.UpdateTenantHandler)
				tenants.DELETE("/:id", handlers.DeleteTenantHandler)
				tenants.GET("/owner/:ownerId", handlers.GetTenantsByOwnerHandler)
			}

			// Contracts
			contracts := protected.Group("/contracts")
			{
				contracts.GET("", handlers.GetContractsHandler)
				contracts.POST("", handlers.CreateContractHandler)
				contracts.GET("/:id", handlers.GetContractHandler)
				contracts.PUT("/:id", handlers.UpdateContractHandler)
				contracts.DELETE("/:id", handlers.DeleteContractHandler)
				contracts.GET("/property/:propertyId", handlers.GetContractsByPropertyHandler)
				contracts.GET("/tenant/:tenantId", handlers.GetContractsByTenantHandler)
				contracts.GET("/property/:propertyId/active", handlers.GetActiveContractByPropertyHandler)
			}

			// Payments
			payments := protected.Group("/payments")
			{
				payments.GET("", handlers.GetPaymentsHandler)
				payments.POST("", handlers.CreatePaymentHandler)
				payments.GET("/:id", handlers.GetPaymentHandler)
				payments.PUT("/:id", handlers.UpdatePaymentHandler)
				payments.DELETE("/:id", handlers.DeletePaymentHandler)
				payments.GET("/contract/:contractId", handlers.GetPaymentsByContractHandler)
				payments.GET("/overdue", handlers.GetOverduePaymentsHandler)
				payments.GET("/period", handlers.GetPaymentsByPeriodHandler)
			}
		}
	}
}

// Handler para servir o YAML diretamente
func (s *Server) serveSwaggerYAML(c *gin.Context) {
	yamlFile, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		logger.Error("Failed to read swagger.yaml", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Swagger documentation not available",
		})
		return
	}

	c.Header("Content-Type", "application/yaml")
	c.String(http.StatusOK, string(yamlFile))
}

// Start inicia o servidor HTTP
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:         ":" + s.config.Port,
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Server starting", zap.String("port", s.config.Port))
	return s.server.ListenAndServe()
}

// Shutdown para o servidor gracefulmente
func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("Shutting down server...")
	return s.server.Shutdown(ctx)
}

// Handlers básicos
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "Aluguei API",
	})
}

func (s *Server) readyHandler(c *gin.Context) {
	// Verificar conectividade com o banco
	if err := s.db.Exec("SELECT 1").Error; err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"error":   "database not available",
			"service": "Aluguei API",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "ready",
		"timestamp": time.Now().UTC(),
		"service":   "Aluguei API",
	})
}

// Auth handlers (mantidos no arquivo principal)
func (s *Server) registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register endpoint"})
}

func (s *Server) loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}

func (s *Server) getCurrentUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get current user endpoint"})
}
