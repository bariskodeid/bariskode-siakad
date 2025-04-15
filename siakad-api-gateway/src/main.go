package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/config"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/routes"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Printf("App Port: %s", config.AppConfig.AppPort)
	log.Printf("Auth Service URL: %s", config.AppConfig.AuthServiceURL)

	// Initialize Gin router with debug mode
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	routes.RegisterRoutes(r)
	routes.RegisterProfileRoutes(r)

	// Start the server
	log.Printf("Starting server on port %s...", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}