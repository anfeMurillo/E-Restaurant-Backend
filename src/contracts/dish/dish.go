package dish

import (
	"context"
	"e-restaurant/models/dish"
)

type Repository interface {
	Create(ctx context.Context, dish *dish.Dish) (*dish.Dish, error)

	GetById(ctx context.Context, dishId int) (*dish.Dish, error)

	GetByRestaurant(ctx context.Context, restaurantId int) (*dish.Dish, error)

	GetAll(ctx context.Context) ([]*dish.Dish, error)

	Delete(ctx context.Context, dishId int) error

	UpdateName(ctx context.Context, dishId int, new string) error

	UpdatePrice(ctx context.Context, dishId int, new string) error
}
