package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Sarvesh-shinde23/microservice_gin_framework/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("Port not found; defaulting to %s\n", port)
	}

	router := gin.Default()    // Correct method to create a default router
	router.Use(cors.Default()) // Use CORS middleware

	// Register routes
	router.POST("/api/register", api.Register)
	router.POST("/api/login", api.Login)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start API server: %v", err)

	}
}
