package restaurant

import (
	"context"
	"e-restaurant/models/restaurant"
)

type Repository interface {
	Create(ctx context.Context, restaurant *restaurant.Restaurant) (*restaurant.Restaurant, error)

	UpdateName(ctx context.Context, restaurantId int, new string) error

	UpdateAddress(ctx context.Context, restaurantId int, new string) error

	UpdateState(ctx context.Context, restaurantId int, new bool) error

	ToggleOpen(ctx context.Context, restaurantId int) error

	GetById(ctx context.Context, restaurantId int) (*restaurant.Restaurant, error)

	Delete(ctx context.Context, restaurantId int) error

	GetAll(ctx context.Context) ([]*restaurant.Restaurant, error)
}
