package services

import (
	"context"
	"order-domain/order-tracking/src/models"
	"order-domain/order-tracking/src/repository"
)

type OrderTrackingService struct {
	repo *repository.OrderTrackingRepository
}

func NewOrderTrackingService(repo *repository.OrderTrackingRepository) *OrderTrackingService {
	return &OrderTrackingService{repo: repo}
}

func (s *OrderTrackingService) GetOrderTracking(ctx context.Context, orderID string) (*models.OrderTracking, error) {
	return s.repo.GetOrderTracking(ctx, orderID)
}
