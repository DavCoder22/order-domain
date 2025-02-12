package routers

import (
	"order-domain/order-service/src/config"
	"order-domain/order-tracking/src/handlers"
	"order-domain/order-tracking/src/repository"
	"order-domain/order-tracking/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize database connection
	db := config.DB

	// Initialize repository and service
	repo := repository.NewOrderTrackingRepository(db)
	service := services.NewOrderTrackingService(repo)
	handler := handlers.NewOrderTrackingHandler(service)

	// Define routes specific to order-tracking
	r.GET("/tracking/:order_id", handler.GetOrderTracking)

	return r
}
