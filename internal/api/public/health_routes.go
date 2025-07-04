// File: internal/api/public/health_routes.go
// Purpose: Registers the always-public healthcheck endpoints for system observability

package public

import (
	"shorty/internal/api/public/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(router *gin.Engine) {
	router.GET("/healthz", handlers.HealthzHandler)
	router.GET("/livez", handlers.LivezHandler)
	router.GET("/readyz", handlers.ReadyzHandler)
}
