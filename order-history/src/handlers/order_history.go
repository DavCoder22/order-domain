package handlers

import (
	"net/http"
	"order-history/services"

	"github.com/gin-gonic/gin"
)

func GetOrderHistory(c *gin.Context) {
	userID := c.Param("user_id")
	history, err := services.GetOrderHistory(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}
