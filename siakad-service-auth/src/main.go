package main

import (
	"log"

	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/models"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/repositories"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()

    if config.AppConfig.AppDebugMode == "true" {
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    err := config.InitDB()
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }
    config.DB.AutoMigrate(&models.User{})

    userRepo := repositories.NewUserRepository(config.DB)

    r := gin.Default()
    r.Use(gin.Recovery())

    routes.RegisterRoutes(r, userRepo)

    log.Printf("Auth service running on port %s", config.AppConfig.AppPort)
    if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
        log.Fatal("failed to run server: ", err)
    }
}
