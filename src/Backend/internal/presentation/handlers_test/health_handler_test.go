package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Turgho/Aluguei/internal/presentation/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type HealthHandlerTestSuite struct {
	suite.Suite
	db      *gorm.DB
	handler *handlers.HealthHandler
	router  *gin.Engine
}

func (suite *HealthHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = db
	suite.handler = handlers.NewHealthHandler(db)
	suite.router = gin.New()
	
	suite.router.GET("/health", suite.handler.Health)
	suite.router.GET("/ready", suite.handler.Ready)
}

func (suite *HealthHandlerTestSuite) TestHealth() {
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "ok", response["status"])
	assert.Equal(suite.T(), "aluguei-api", response["service"])
	assert.Equal(suite.T(), "1.0.0", response["version"])
}

func (suite *HealthHandlerTestSuite) TestReady() {
	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "ready", response["status"])
	assert.Equal(suite.T(), "connected", response["database"])
}

func (suite *HealthHandlerTestSuite) TestReady_DatabaseError() {
	// Close the database to simulate an error
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()

	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusServiceUnavailable, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "not ready", response["status"])
}

func TestHealthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthHandlerTestSuite))
}