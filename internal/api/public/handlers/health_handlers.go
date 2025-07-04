// File: internal/api/public/handlers/health_handlers.go
// Purpose: Handles system health endpoints /healthz, /livez, /readyz

package handlers

import (
	"net/http"

	"shorty/internal/core/system"

	"github.com/gin-gonic/gin"
)

func HealthzHandler(c *gin.Context) {
	status := system.GenerateHealthReport()
	if status.Status != "ok" {
		c.JSON(http.StatusServiceUnavailable, status)
	} else {
		c.JSON(http.StatusOK, status)
	}
}

func LivezHandler(c *gin.Context) {
	c.String(http.StatusOK, "yes")
}

func ReadyzHandler(c *gin.Context) {
	if system.IsReady() {
		c.String(http.StatusOK, "yes")
	} else {
		status := gin.H{
			"ready":   "no",
			"database": system.DBReady(),
			"redis":    system.RedisReady(),
		}
		c.JSON(http.StatusServiceUnavailable, status)
	}
}
