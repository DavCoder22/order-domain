package routers

import (
	"order-history/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	history := router.Group("/history")
	{
		history.GET("/:user_id", handlers.GetOrderHistory)
	}
}
