package lostinventory

import (
	lostinventory "e-restaurant/models/lostInventory"
	"time"
)

type Repository interface {
	Create(lInvenotory *lostinventory.LostInventory) (*lostinventory.LostInventory, error)

	GetByRestaurant(restaurantId int, limit int, since time.Time) ([]*lostinventory.LostInventory, error)
}
