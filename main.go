package main

import (
	"day-46/config"
	"day-46/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup Gin Router
	router := gin.Default()

	// Load routes
	route.SetupRoutes(router)

	// Start server
	log.Printf("Server running on %s\n", config.ServerPort)
	router.Run(config.ServerPort)
}
