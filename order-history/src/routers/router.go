package routers

import (
	"order-domain/order-history/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	history := router.Group("/history")
	{
		history.GET("/:user_id", handlers.GetOrderHistory)
	}
}
