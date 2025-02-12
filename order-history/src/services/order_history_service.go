package services

import (
	"context"
	"order-domain/order-history/src/models"
	"order-domain/order-history/src/repository"
)

type OrderHistoryService struct {
	repo *repository.OrderHistoryRepository
}

func NewOrderHistoryService(repo *repository.OrderHistoryRepository) *OrderHistoryService {
	return &OrderHistoryService{repo: repo}
}

func (s *OrderHistoryService) GetOrderHistory(ctx context.Context, userID string) (*models.OrderHistory, error) {
	return s.repo.GetOrderHistory(ctx, userID)
}
