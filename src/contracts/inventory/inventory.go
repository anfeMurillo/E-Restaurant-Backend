package inventory

import (
	"context"
	"e-restaurant/models/enums/measure"
	"e-restaurant/models/inventory"
	"time"
)

type Repository interface {
	Create(ctx context.Context, inventory *inventory.Inventory) (inventory.Inventory, error)

	AddStock(ctx context.Context, quantity float64, unit measure.Measure, expitationDate time.Time) error

	RemoveStock(ctx context.Context, quantity float64, unit measure.Measure) (string, error)

	UpdateUnit(ctx context.Context, InventoryId int, unit measure.Measure) error

	UpdateExpitationDate(ctx context.Context, InventoryId int, expirationDate time.Time) error
}
