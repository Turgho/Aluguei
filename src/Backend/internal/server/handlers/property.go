package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Property handlers
func GetPropertiesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get properties endpoint"})
}

func CreatePropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create property endpoint"})
}

func GetPropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get property endpoint"})
}

func UpdatePropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update property endpoint"})
}

func DeletePropertyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete property endpoint"})
}
