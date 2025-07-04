// File: internal/api/public/handlers/healthz.go
// Purpose: Implements logic for /healthz, /readyz, and /livez public endpoints.

package handlers

import (
	"net/http"
	"time"

	"shorty/internal/core"
	"shorty/internal/lib/uptime"

	"github.com/gin-gonic/gin"
)

var bootTime = time.Now()

func HealthzHandler(c *gin.Context) {
	status, ready := core.CheckSystemStatus()

	components := core.GetComponentStatuses()

	c.JSON(getStatusCode(status, ready), gin.H{
		"status":        status,
		"ready":         boolToYesNo(ready),
		"uptime":        uptime.Since(bootTime),
		"version":       core.Version(),
		"build_commit":  core.BuildCommit(),
		"build_date":    core.BuildDate(),
		"checked_at":    time.Now().UTC().Format(time.RFC3339),
		"public_server": boolToYesNo(core.IsPublicServer()),
		"public_user":   boolToYesNo(core.IsPublicUserEnabled()),
		"public_orgs":   boolToYesNo(core.IsPublicOrgsEnabled()),
		"components":    components,
	})
}

func ReadyzHandler(c *gin.Context) {
	ready, parts := core.CheckReadiness()

	if ready {
		c.String(http.StatusOK, "yes")
		return
	}

	c.JSON(http.StatusServiceUnavailable, gin.H{
		"ready":   "no",
		"details": parts,
	})
}

func LivezHandler(c *gin.Context) {
	c.String(http.StatusOK, "yes")
}

// Helpers

func getStatusCode(status string, ready bool) int {
	if status != "ok" || !ready {
		return http.StatusServiceUnavailable
	}
	return http.StatusOK
}

func boolToYesNo(val bool) string {
	if val {
		return "yes"
	}
	return "no"
}
