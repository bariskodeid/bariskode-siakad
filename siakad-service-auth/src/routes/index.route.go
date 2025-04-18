package routes

import (
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/controllers"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/repositories"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, repo *repositories.UserRepository) {
    r.POST("/login", controllers.NewAuthHandler(repo).Login)
    r.POST("/register", controllers.NewAuthHandler(repo).Register)
    r.POST("/refresh-token", controllers.NewAuthHandler(repo).RefreshToken)
    r.POST("/logout", controllers.NewAuthHandler(repo).Logout)
}

func RegisterProfileRoutes(r *gin.RouterGroup) {
    profile := r.Group("/profile")
    {
        profile.GET("/:user_id", controllers.GetProfile)
        // profile.PUT("/:user_id", controllers.UpdateProfile)
    }
}
