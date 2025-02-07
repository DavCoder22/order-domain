package repository

import (
	"context"
	"order-domain/order-service/src/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	// Insertar orden
	_, err = tx.Exec(ctx,
		`INSERT INTO orders (id, user_id, status, total, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)`,
		order.ID, order.UserID, order.Status, order.Total, order.CreatedAt, order.UpdatedAt,
	)
	if err != nil {
		return err
	}
	// Insertar items
	for _, item := range order.Items {
		_, err = tx.Exec(ctx,
			`INSERT INTO order_items (order_id, product_id, quantity, price)
            VALUES ($1, $2, $3, $4)`,
			order.ID, item.ProductID, item.Quantity, item.Price,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (r *OrderRepository) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	var order models.Order
	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, status, total, created_at, updated_at FROM orders WHERE id = $1`,
		orderID,
	).Scan(&order.ID, &order.UserID, &order.Status, &order.Total, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Obtener los items de la orden
	rows, err := r.db.Query(ctx,
		`SELECT product_id, quantity, price FROM order_items WHERE order_id = $1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID string, status models.OrderStatus) error {
	_, err := r.db.Exec(ctx,
		`UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`,
		status, orderID,
	)
	return err
}
