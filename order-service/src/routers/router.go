package router

import (
	"order-domain/order-service/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	return r
}
