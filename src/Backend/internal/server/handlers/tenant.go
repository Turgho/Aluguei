package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tenant handlers
func GetTenantsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get tenants endpoint"})
}

func CreateTenantHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create tenant endpoint"})
}

func GetTenantHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get tenant endpoint"})
}

func UpdateTenantHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update tenant endpoint"})
}

func DeleteTenantHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete tenant endpoint"})
}

func GetTenantsByOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get tenants by owner endpoint"})
}
