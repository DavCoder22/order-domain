package main

import (
	"context"
	"log"
	"os"

	"order-domain/order-history/src/config"
	"order-domain/order-history/src/handlers"
	"order-domain/order-history/src/middleware"
	"order-domain/order-history/src/repository"
	service "order-domain/order-history/src/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Load configuration
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("SUPABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Initialize services
	repo := repository.NewOrderRepository(conn)
	orderService := service.NewOrderService(repo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Configure Gin server
	r := gin.Default()

	// Use temporary authentication middleware
	r.Use(middleware.TemporaryAuthMiddleware())

	// Define routes
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Start server
	log.Printf("Server started on port %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
