package handlers

import (
	"net/http"
	"order-billing/src/models"
	service "order-billing/src/services"

	"github.com/gin-gonic/gin"
)

type BillingHandler struct {
	service service.BillingServiceInterface
}

func NewBillingHandler(s service.BillingServiceInterface) *BillingHandler {
	return &BillingHandler{service: s}
}

func (h *BillingHandler) GenerateInvoice(c *gin.Context) {
	var req models.InvoiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format", "detail": err.Error()})
		return
	}

	invoice, err := h.service.GenerateInvoice(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating invoice", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}
