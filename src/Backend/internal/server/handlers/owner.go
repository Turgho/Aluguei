package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Owner handlers
func GetOwnersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get owners endpoint"})
}

func CreateOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create owner endpoint"})
}

func GetOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get owner endpoint"})
}

func UpdateOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update owner endpoint"})
}

func DeleteOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete owner endpoint"})
}

func GetOwnerByEmailHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get owner by email endpoint"})
}

func GetPropertiesByOwnerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get properties by owner endpoint"})
}
