// File: internal/api/public/routes.go
// Purpose: Defines all public-facing API endpoints, including creation, redirection, health, and Swagger docs.

package public

import (
	"shorty/internal/api/public/handlers"
	"shorty/internal/api/public/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterPublicRoutes sets up all public endpoints.
func RegisterPublicRoutes(r *gin.Engine) {
	public := r.Group("/")
	public.Use(middleware.ThemeMiddleware(), middleware.CLIDetector())

	public.GET("/create", handlers.CreateLink)
	public.POST("/create", handlers.CreateLink)

	public.GET("/:slug", handlers.RedirectLink)

	public.GET("/healthz", handlers.HealthCheck)
	public.GET("/readyz", handlers.ReadyCheck)
	public.GET("/livez", handlers.LiveCheck)

	public.GET("/docs", handlers.SwaggerDocs)
}
