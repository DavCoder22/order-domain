package handlers

import (
	"net/http"
	"order-domain/order-tracking/src/services"

	"github.com/gin-gonic/gin"
)

type OrderTrackingHandler struct {
	service *services.OrderTrackingService
}

func NewOrderTrackingHandler(service *services.OrderTrackingService) *OrderTrackingHandler {
	return &OrderTrackingHandler{service: service}
}

func (h *OrderTrackingHandler) GetOrderTracking(c *gin.Context) {
	orderID := c.Param("order_id")
	tracking, err := h.service.GetOrderTracking(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tracking not found"})
		return
	}

	c.JSON(http.StatusOK, tracking)
}
