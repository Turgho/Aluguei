// pkg/response/json.go
package response

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, ErrorResponse{
		Error: message,
		Code:  code,
	})
}

func ValidationError(c *gin.Context, details map[string]string) {
	c.JSON(400, ErrorResponse{
		Error:   "erro de validação",
		Code:    "VALIDATION_ERROR",
		Details: details,
	})
}
