package models

type InvoiceRequest struct {
	OrderID string  `json:"order_id"`
	UserID  string  `json:"user_id"`
	Total   float64 `json:"total"`
	Items   []Item  `json:"items"`
}

type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type InvoiceResponse struct {
	InvoiceID string  `json:"invoice_id"`
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Total     float64 `json:"total"`
	Items     []Item  `json:"items"`
	CreatedAt string  `json:"created_at"`
}
