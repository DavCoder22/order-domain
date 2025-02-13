package routers

import (
	"order-domain/order-history/src/config"
	"order-domain/order-history/src/handlers"
	"order-domain/order-history/src/repository"
	"order-domain/order-history/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize database connection
	db := config.DB

	// Initialize repository and service
	repo := repository.NewOrderHistoryRepository(db)
	service := services.NewOrderHistoryService(repo)
	handler := handlers.NewOrderHistoryHandler(service)

	// Define routes specific to order-history
	r.GET("/users/:user_id/history", handler.GetOrderHistory)

	return r
}
