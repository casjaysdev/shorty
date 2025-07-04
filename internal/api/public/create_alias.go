// File: internal/api/public/create_alias.go
// Purpose: Implements the /create alias route for stable, public paste creation across versions.

package public

import (
	"net/http"

	"shorty/internal/api/shared"
	"shorty/internal/core"
	"shorty/internal/models"
	"shorty/internal/utils"

	"github.com/gin-gonic/gin"
)

// RegisterCreateAlias adds the /create public alias for the versioned create endpoint.
func RegisterCreateAlias(router *gin.RouterGroup) {
	router.POST("/create", createHandler)
}

// createHandler handles incoming POST requests to /create.
func createHandler(c *gin.Context) {
	var req models.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request payload")
		return
	}

	resp, err := core.CreateLink(shared.RequestContext(c), req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, http.StatusCreated, resp)
}
