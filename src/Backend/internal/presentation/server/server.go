package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Turgho/Aluguei/internal/application/usecases"
	"github.com/Turgho/Aluguei/internal/infrastructure/middleware"
	"github.com/Turgho/Aluguei/internal/infrastructure/persistence"
	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

func New(db *gorm.DB) *Server {
	// Initialize repositories
	ownerRepo := persistence.NewOwnerRepository(db)
	tenantRepo := persistence.NewTenantRepository(db)
	propertyRepo := persistence.NewPropertyRepository(db)
	contractRepo := persistence.NewContractRepository(db)
	paymentRepo := persistence.NewPaymentRepository(db)

	// Initialize use cases
	ownerUseCase := usecases.NewOwnerUseCase(ownerRepo)
	tenantUseCase := usecases.NewTenantUseCase(tenantRepo)
	propertyUseCase := usecases.NewPropertyUseCase(propertyRepo)
	contractUseCase := usecases.NewContractUseCase(contractRepo)
	paymentUseCase := usecases.NewPaymentUseCase(paymentRepo)

	// Initialize handlers
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key"
	}
	authHandler := handlers.NewAuthHandler(ownerUseCase, jwtSecret)
	ownerHandler := handlers.NewOwnerHandler(ownerUseCase)
	tenantHandler := handlers.NewTenantHandler(tenantUseCase)
	propertyHandler := handlers.NewPropertyHandler(propertyUseCase)
	contractHandler := handlers.NewContractHandler(contractUseCase)
	paymentHandler := handlers.NewPaymentHandler(paymentUseCase)
	dashboardHandler := handlers.NewDashboardHandler(propertyUseCase, contractUseCase, paymentUseCase)
	healthHandler := handlers.NewHealthHandler(db)
	swaggerHandler := handlers.NewSwaggerHandler()

	// Setup router
	router := gin.Default()
	
	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Health routes
	router.GET("/health", healthHandler.Health)
	router.GET("/ready", healthHandler.Ready)

	// Swagger routes
	router.GET("/swagger", swaggerHandler.ServeSwaggerUI)
	router.GET("/swagger/swagger.yaml", swaggerHandler.ServeSwaggerYAML)

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// Owner routes
		owners := api.Group("/owners")
		{
			owners.POST("", ownerHandler.CreateOwner) // Public - registration
			
			// Protected routes
			protected := owners.Use(middleware.AuthMiddleware(jwtSecret))
			protected.GET("", ownerHandler.GetAllOwners)
			protected.GET("/:id", ownerHandler.GetOwner)
			protected.PUT("/:id", ownerHandler.UpdateOwner)
			protected.DELETE("/:id", ownerHandler.DeleteOwner)
			protected.GET("/email/:email", ownerHandler.GetOwnerByEmail)
		}

		// Tenant routes (all protected)
		tenants := api.Group("/tenants").Use(middleware.AuthMiddleware(jwtSecret))
		{
			tenants.POST("", tenantHandler.CreateTenant)
			tenants.GET("", tenantHandler.GetAllTenants)
			tenants.GET("/:id", tenantHandler.GetTenant)
			tenants.DELETE("/:id", tenantHandler.DeleteTenant)
			tenants.GET("/owner/:ownerId", tenantHandler.GetTenantsByOwner)
		}

		// Property routes (all protected)
		properties := api.Group("/properties").Use(middleware.AuthMiddleware(jwtSecret))
		{
			properties.POST("", propertyHandler.CreateProperty)
			properties.GET("", propertyHandler.GetAllProperties)
			properties.GET("/:id", propertyHandler.GetProperty)
			properties.PUT("/:id", propertyHandler.UpdateProperty)
			properties.DELETE("/:id", propertyHandler.DeleteProperty)
			properties.GET("/owner/:ownerId", propertyHandler.GetPropertiesByOwner)
		}

		// Contract routes (all protected)
		contracts := api.Group("/contracts").Use(middleware.AuthMiddleware(jwtSecret))
		{
			contracts.POST("", contractHandler.CreateContract)
			contracts.GET("", contractHandler.GetContracts)
			contracts.GET("/:id", contractHandler.GetContractByID)
			contracts.PUT("/:id", contractHandler.UpdateContract)
			contracts.DELETE("/:id", contractHandler.DeleteContract)
			contracts.GET("/property/:propertyId", contractHandler.GetContractsByProperty)
			contracts.GET("/tenant/:tenantId", contractHandler.GetContractsByTenant)
			contracts.GET("/property/:propertyId/active", contractHandler.GetActiveContractByProperty)
		}

		// Payment routes (all protected)
		payments := api.Group("/payments").Use(middleware.AuthMiddleware(jwtSecret))
		{
			payments.POST("", paymentHandler.CreatePayment)
			payments.GET("", paymentHandler.GetPayments)
			payments.GET("/:id", paymentHandler.GetPaymentByID)
			payments.PUT("/:id", paymentHandler.UpdatePayment)
			payments.DELETE("/:id", paymentHandler.DeletePayment)
			payments.GET("/contract/:contractId", paymentHandler.GetPaymentsByContract)
			payments.GET("/overdue", paymentHandler.GetOverduePayments)
			payments.GET("/period", paymentHandler.GetPaymentsByPeriod)
		}

		// Dashboard routes (all protected)
		dashboard := api.Group("/dashboard").Use(middleware.AuthMiddleware(jwtSecret))
		{
			dashboard.GET("/owner/:ownerId", dashboardHandler.GetDashboard)
		}
	}

	return &Server{
		router: router,
		server: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}