package user

import (
	"context" // Esto es util para la DB y poder cerrar
	modeluser "e-restaurant/models/modelUser"
)

type Repository interface {
	Create(ctx context.Context, user *modeluser.User) (*modeluser.User, error)

	GetById(ctx context.Context, userId int) (*modeluser.User, error)

	GetRestaurantStaff(ctx context.Context, role string, restaurantId int) ([]*modeluser.User, error)

	UpdateName(ctx context.Context, userId int, new string) error

	UpdateEmail(ctx context.Context, userId int, new string) error

	UpdateCountryCode(ctx context.Context, userId int, new string) error

	UpdatePassword(ctx context.Context, userId int, newHash string) error

	ToggleActiveState(ctx context.Context, userId int) error

	AssignRole(ctx context.Context, userId int, role string) error
}
