package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type SwaggerHandler struct{}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (h *SwaggerHandler) ServeSwaggerUI(c *gin.Context) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Aluguei API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@4.15.5/swagger-ui.css" />
    <style>
        html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
        *, *:before, *:after { box-sizing: inherit; }
        body { margin:0; background: #fafafa; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@4.15.5/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@4.15.5/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '/swagger/swagger.yaml',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}

func (h *SwaggerHandler) ServeSwaggerYAML(c *gin.Context) {
	yamlPath := filepath.Join("docs", "swagger.yaml")
	
	yamlData, err := os.ReadFile(yamlPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "swagger.yaml not found"})
		return
	}

	// Parse and re-serialize to ensure valid YAML
	var swaggerDoc interface{}
	if err := yaml.Unmarshal(yamlData, &swaggerDoc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid swagger.yaml"})
		return
	}

	c.Header("Content-Type", "application/x-yaml")
	c.Data(http.StatusOK, "application/x-yaml", yamlData)
}