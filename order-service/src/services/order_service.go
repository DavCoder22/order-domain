package service

import (
	"context"
	"errors"

	"order-domain/order-service/src/models" // O cambia a "order-domain/order-service/src/models" si es necesario
	"order-domain/order-service/src/repository"
)

// OrderServiceInterface define los métodos del servicio de órdenes
type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status models.OrderStatus) error
}

// OrderService es la implementación del servicio de órdenes
type OrderService struct {
	repo *repository.OrderRepository
}

// NewOrderService crea una nueva instancia del servicio
func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

// CreateOrder crea un pedido en la base de datos
func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	if order == nil {
		return errors.New("la orden no puede ser nula")
	}
	return s.repo.CreateOrder(ctx, order)
}

// GetOrder obtiene un pedido por ID
func (s *OrderService) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	if orderID == "" {
		return nil, errors.New("el ID de la orden no puede estar vacío")
	}
	return s.repo.GetOrder(ctx, orderID)
}

// UpdateOrderStatus actualiza el estado de un pedido
func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID string, status models.OrderStatus) error {
	if orderID == "" {
		return errors.New("el ID de la orden no puede estar vacío")
	}
	if status == "" {
		return errors.New("el estado de la orden no puede estar vacío")
	}
	return s.repo.UpdateOrderStatus(ctx, orderID, status)
}
