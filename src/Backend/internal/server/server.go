package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Turgho/Aluguei/internal/config"
	"github.com/Turgho/Aluguei/internal/middlewares"
	"github.com/Turgho/Aluguei/pkg/logger"
	"github.com/gin-gonic/gin"
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

			// Properties
			properties := protected.Group("/properties")
			{
				properties.GET("", s.getPropertiesHandler)
				properties.POST("", s.createPropertyHandler)
				properties.GET("/:id", s.getPropertyHandler)
				properties.PUT("/:id", s.updatePropertyHandler)
				properties.DELETE("/:id", s.deletePropertyHandler)
			}
		}
	}
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

// Placeholder handlers - serão implementados depois
func (s *Server) registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register endpoint"})
}

func (s *Server) loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}

func (s *Server) getCurrentUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get current user endpoint"})
}

func (s *Server) getPropertiesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get properties endpoint"})
}

func (s *Server) createPropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create property endpoint"})
}

func (s *Server) getPropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get property endpoint"})
}

func (s *Server) updatePropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update property endpoint"})
}

func (s *Server) deletePropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete property endpoint"})
}
