package main

import (
	"log"

	"order-domain/order-history/src/config"
	"order-domain/order-history/src/routers"
)

func main() {
	// Load configuration from order-service
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database connection
	config.InitDB()

	// Setup and run the router for order-history
	r := routers.SetupRouter()
	log.Printf("Server started on port %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
