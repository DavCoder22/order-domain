package handlers

import (
	"net/http"
	"order-tracking/services"

	"github.com/gin-gonic/gin"
)

func TrackOrder(c *gin.Context) {
	id := c.Param("id")
	status, err := services.TrackOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}
