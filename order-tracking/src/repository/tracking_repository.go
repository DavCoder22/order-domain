package repository

import (
	"context"
	"order-domain/order-tracking/src/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderTrackingRepository struct {
	db *pgxpool.Pool
}

func NewOrderTrackingRepository(db *pgxpool.Pool) *OrderTrackingRepository {
	return &OrderTrackingRepository{db: db}
}

func (r *OrderTrackingRepository) GetOrderTracking(ctx context.Context, orderID string) (*models.OrderTracking, error) {
	var tracking models.OrderTracking
	rows, err := r.db.Query(ctx,
		`SELECT order_id, id, status, timestamp, details FROM order_tracking WHERE order_id = $1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.TrackingEvent
		if err := rows.Scan(&tracking.OrderID, &event.ID, &event.Status, &event.Timestamp, &event.Details); err != nil {
			return nil, err
		}
		tracking.TrackingEvents = append(tracking.TrackingEvents, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tracking, nil
}
