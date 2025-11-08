package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDashboardHandler_GetDashboard_InvalidOwnerID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	handler := handlers.NewDashboardHandler(nil, nil, nil)
	
	req := httptest.NewRequest("GET", "/dashboard/invalid-uuid", nil)
	w := httptest.NewRecorder()
	
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "ownerId", Value: "invalid-uuid"}}
	
	handler.GetDashboard(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
}