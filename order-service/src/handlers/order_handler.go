package handlers

import (
	"net/http"
	"order-domain/order-service/src/models"
	service "order-domain/order-service/src/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler struct {
	service service.OrderServiceInterface
}

func NewOrderHandler(s service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{service: s}
}

// 📌 Crear una orden
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.Order

	// Verificar si el JSON es válido
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Formato inválido",
			"detalle": err.Error(),
		})
		return
	}

	// Generar ID y timestamps
	req.ID = generateOrderID()
	req.Status = models.StatusPending
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	// Guardar la orden en la base de datos
	if err := h.service.CreateOrder(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creando pedido",
			"detalle": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// 📌 Obtener una orden por ID
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("id")

	// Llamar al servicio para obtener la orden
	order, err := h.service.GetOrder(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// 📌 Actualizar estado de una orden
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var statusUpdate models.StatusUpdate

	// Leer el JSON con el nuevo estado
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Formato inválido",
			"detalle": err.Error(),
		})
		return
	}

	// Verificar que el estado no esté vacío
	if statusUpdate.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "El estado no puede estar vacío",
			"detalle": "El campo 'status' es obligatorio",
		})
		return
	}

	// Llamar al servicio para actualizar el estado
	if err := h.service.UpdateOrderStatus(c.Request.Context(), orderID, statusUpdate.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error actualizando estado del pedido",
			"detalle": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado del pedido actualizado"})
}

// 📌 Generar ID para las órdenes
func generateOrderID() string {
	return uuid.New().String()
}
