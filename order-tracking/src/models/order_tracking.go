package models

import "time"

type TrackingEvent struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}

type OrderTracking struct {
	OrderID        string          `json:"order_id"`
	TrackingEvents []TrackingEvent `json:"tracking_events"`
}
