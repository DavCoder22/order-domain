package handlers

import (
	"net/http"
	"order-domain/order-history/src/services"

	"github.com/gin-gonic/gin"
)

type OrderHistoryHandler struct {
	service *services.OrderHistoryService
}

func NewOrderHistoryHandler(service *services.OrderHistoryService) *OrderHistoryHandler {
	return &OrderHistoryHandler{service: service}
}

func (h *OrderHistoryHandler) GetOrderHistory(c *gin.Context) {
	userID := c.Param("user_id")
	orderHistory, err := h.service.GetOrderHistory(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order history not found"})
		return
	}

	c.JSON(http.StatusOK, orderHistory)
}
