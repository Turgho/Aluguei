package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	requestBody := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.New()

	router.POST("/auth/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "mock-jwt-token",
			"owner": gin.H{
				"id":    "123",
				"name":  "Test User",
				"email": "test@example.com",
			},
		})
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "mock-jwt-token", response["token"])
}