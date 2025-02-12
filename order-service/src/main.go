package main

import (
	"log"

	"order-domain/order-service/src/config"
	"order-domain/order-service/src/handlers"
	"order-domain/order-service/src/repository"
	service "order-domain/order-service/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	if err := config.LoadConfig("."); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database connection
	config.InitDB()

	// Use the global connection initialized in config.InitDB()
	db := config.DB
	defer db.Close()

	// Initialize services
	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Configure Gin server
	r := gin.Default()

	// Define routes without authentication middleware
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Start server
	log.Printf("Server started on port %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
