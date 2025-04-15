package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/handlers"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/middlewares"
)

func RegisterProfileRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware()) 
	
	profile := api.Group("/profile")
	{
		profile.GET("/", handlers.GetProfileHandler)
		profile.PUT("/", handlers.UpdateProfileHandler)
		profile.POST("/change-password", handlers.ChangePasswordHandler)
		profile.POST("/upload-avatar", handlers.UploadAvatarHandler)
	}
}