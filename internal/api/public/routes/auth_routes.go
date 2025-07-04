// File: internal/api/public/routes/auth_routes.go
// Purpose: Defines routes for public authentication endpoints (login, register, reset, verify, etc.)

package routes

import (
	"shorty/internal/api/public/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/logout", handlers.LogoutHandler)
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/forgot", handlers.ForgotPasswordHandler)
		auth.POST("/reset", handlers.ResetPasswordHandler)
		auth.GET("/verify/:token", handlers.VerifyEmailHandler)
	}
}
