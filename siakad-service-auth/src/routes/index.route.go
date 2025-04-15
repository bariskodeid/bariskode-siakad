package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
    r.POST("/register", controllers.Register)
}

func RegisterProfileRoutes(r *gin.RouterGroup) {
    profile := r.Group("/profile")
    {
        profile.GET("/:user_id", controllers.GetProfile)
        // profile.PUT("/:user_id", controllers.UpdateProfile)
    }
}
