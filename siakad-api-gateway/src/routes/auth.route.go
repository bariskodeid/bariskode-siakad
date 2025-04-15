package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/logout", handlers.LogoutHandler)
		auth.POST("/register", handlers.RegisterHandler)
	}
}