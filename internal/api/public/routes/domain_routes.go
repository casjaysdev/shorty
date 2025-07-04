// File: internal/api/public/routes/domain_routes.go
// Purpose: Exposes public routes for listing and checking available domains.

package routes

import (
	"shorty/internal/api/public/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterDomainRoutes(rg *gin.RouterGroup) {
	domains := rg.Group("/domains")
	{
		domains.GET("/", handlers.ListDomainsHandler)
		domains.GET("/check/:name", handlers.CheckDomainAvailabilityHandler)
	}
}
