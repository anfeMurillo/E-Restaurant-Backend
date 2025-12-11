package order

import (
	"context"
	orderstatus "e-restaurant/models/enums/orderStatus"
	"e-restaurant/models/order"
)

type Repository interface {
	Create(ctx context.Context, order *order.Order) (*order.Order, error)

	GetById(ctx context.Context, orderID int) (*order.Order, error)

	GetByRestaurant(ctx context.Context) ([]*order.Order, error)

	UpdateStatus(ctx context.Context, oStatus orderstatus.OrderStatus) error
}
