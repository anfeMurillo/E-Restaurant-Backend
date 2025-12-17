package inventory

import (
	"context"
	"e-restaurant/models/enums/measure"
	"e-restaurant/models/inventory"
	"time"
)

type Repository interface {
	Create(ctx context.Context, inventory *inventory.Inventory) (*inventory.Inventory, error)

	GetAll(ctx context.Context) ([]*inventory.Inventory, error)

	GetById(ctx context.Context, inventoryId int) (*inventory.Inventory, error)

	GetByRestaurant(ctx context.Context, restaurantId int) ([]*inventory.Inventory, error)

	AddStock(ctx context.Context, inventoryId int, quantity float64, expitationDate time.Time) error

	RemoveStock(ctx context.Context, inventoryId int, quantity float64, unit measure.Measure) error

	UpdateUnit(ctx context.Context, inventoryId int, unit measure.Measure) error

	UpdateExpitationDate(ctx context.Context, inventoryId int, expirationDate time.Time) error

	Delete(ctx context.Context, inventoryId int) error
}
