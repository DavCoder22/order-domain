package repository

import (
	"context"
	"order-domain/order-history/src/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderHistoryRepository struct {
	db *pgxpool.Pool
}

func NewOrderHistoryRepository(db *pgxpool.Pool) *OrderHistoryRepository {
	return &OrderHistoryRepository{db: db}
}

func (r *OrderHistoryRepository) GetOrderHistory(ctx context.Context, userID string) (*models.OrderHistory, error) {
	var orderHistory models.OrderHistory
	rows, err := r.db.Query(ctx,
		`SELECT user_id, order_id, product_id, quantity, status, created_at, updated_at FROM order_history WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&orderHistory.UserID, &order.ID, &order.ProductID, &order.Quantity, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orderHistory.Orders = append(orderHistory.Orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &orderHistory, nil
}
