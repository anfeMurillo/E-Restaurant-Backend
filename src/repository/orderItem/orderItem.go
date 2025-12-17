package orderitem

import (
	"context"
	"database/sql"
	model "e-restaurant/models/orderItem"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Create(ctx context.Context, oItem *model.OrderItem) (*model.OrderItem, error) {
	query := `INSERT INTO order_items (order_id, dish_id, quantity, price_at_moment)
			  VALUES ($1, $2, $3, $4)
			  RETURNING order_item_id, order_id, dish_id, quantity, price_at_moment`
	row := r.db.QueryRowContext(ctx, query, oItem.OrderId, oItem.DishId, oItem.Quantity, oItem.PriceAtMoment)
	var res model.OrderItem
	if err := row.Scan(&res.OrderItemId, &res.OrderId, &res.DishId, &res.Quantity, &res.PriceAtMoment); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *OrderItemRepository) GetById(ctx context.Context, oItemId int) (*model.OrderItem, error) {
	query := `SELECT order_item_id, order_id, dish_id, quantity, price_at_moment FROM order_items WHERE order_item_id = $1`
	row := r.db.QueryRowContext(ctx, query, oItemId)
	var res model.OrderItem
	if err := row.Scan(&res.OrderItemId, &res.OrderId, &res.DishId, &res.Quantity, &res.PriceAtMoment); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *OrderItemRepository) GetAllByOrder(ctx context.Context, orderId int) ([]model.OrderItem, error) {
	query := `SELECT order_item_id, order_id, dish_id, quantity, price_at_moment FROM order_items WHERE order_id = $1`
	rows, err := r.db.QueryContext(ctx, query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []model.OrderItem
	for rows.Next() {
		var oi model.OrderItem
		if err := rows.Scan(&oi.OrderItemId, &oi.OrderId, &oi.DishId, &oi.Quantity, &oi.PriceAtMoment); err != nil {
			return nil, err
		}
		result = append(result, oi)
	}
	return result, nil
}

func (r *OrderItemRepository) Delete(ctx context.Context, oItemId int) error {
	query := `DELETE FROM order_items WHERE order_item_id = $1`
	_, err := r.db.ExecContext(ctx, query, oItemId)
	return err
}

func (r *OrderItemRepository) UpdateQuantity(ctx context.Context, oItemId int, quantity int) error {
	query := `UPDATE order_items SET quantity = $1 WHERE order_item_id = $2`
	_, err := r.db.ExecContext(ctx, query, quantity, oItemId)
	return err
}
