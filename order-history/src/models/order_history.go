package models

import "time"

// OrderHistory represents the history of orders for a user.
type OrderHistory struct {
	UserID string  `json:"user_id"`
	Orders []Order `json:"orders"`
}

// Order represents an individual order.
type Order struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
