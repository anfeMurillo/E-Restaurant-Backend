package restaurantstaff

import (
	"context"
	restaurantstaff "e-restaurant/models/restaurantStaff"
	userapp "e-restaurant/models/userApp"
)

type Repository interface {
	Create(ctx context.Context, staff *restaurantstaff.RestaurantStaff) (*restaurantstaff.RestaurantStaff, error)

	Delete(ctx context.Context, restaurantId int, userId int) error

	GetByRestaurant(ctx context.Context, restaurantId int, onlyActive bool) ([]*userapp.User, error)
}
