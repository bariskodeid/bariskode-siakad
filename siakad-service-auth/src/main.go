package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/routes"
)

func main() {
    config.LoadConfig()

    r := gin.Default()
    routes.RegisterRoutes(r)

    log.Printf("Auth service running on port %s", config.AppConfig.AppPort)
    if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
        log.Fatal("failed to run server: ", err)
    }
}
