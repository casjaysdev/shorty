// File: internal/api/public/routes/health_routes.go
// Purpose: Sets up public health endpoints: /healthz, /livez, /readyz for system observability.

package routes

import (
	"shorty/internal/api/public/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/")
	{
		health.GET("/healthz", handlers.HealthzHandler)
		health.GET("/livez", handlers.LivezHandler)
		health.GET("/readyz", handlers.ReadyzHandler)
	}
}
