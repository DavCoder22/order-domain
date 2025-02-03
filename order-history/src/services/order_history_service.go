package services

import (
	"errors"
	"order-domain/order-history/models"
)

var histories = make(map[string]models.OrderHistory)

func GetOrderHistory(userID string) (models.OrderHistory, error) {
	history, exists := histories[userID]
	if !exists {
		return models.OrderHistory{}, errors.New("history not found")
	}
	return history, nil
}
