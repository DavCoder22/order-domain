package routers

import (
	"order-tracking/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	tracking := router.Group("/tracking")
	{
		tracking.GET("/:id", handlers.TrackOrder)
	}
}
