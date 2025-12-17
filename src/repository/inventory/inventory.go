package inventory

import (
	"context"
	"database/sql"
	"e-restaurant/models/enums/measure"
	"e-restaurant/models/inventory"
	"time"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepository {
	return InventoryRepository{db: db}
}

func (i *InventoryRepository) Create(ctx context.Context, inventory *inventory.Inventory) (*inventory.Inventory, error) {
	query := `
	INSERT INTO inventories (
	restaurant_id,
	ingredient_id,
	stock,
	unit,
	expiration_date			)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING inventory_id,restaurant_id,ingredient_id
	`

	err := i.db.QueryRowContext(
		ctx,
		query,
		inventory.RestaurantId,
		inventory.IngredientId,
		inventory.Stock,
		inventory.Unit,
		inventory.ExpirationDate,
	).Scan(
		&inventory.InventoryId,
		&inventory.RestaurantId,
		&inventory.IngredientId,
	)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (i *InventoryRepository) GetAll(ctx context.Context) ([]*inventory.Inventory, error) {
	query := `
	SELECT * FROM inventories
	`

	rows, err := i.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	inventoryItemList := []*inventory.Inventory{}

	for rows.Next() {
		inventoryItem := &inventory.Inventory{}

		rows.Scan(
			inventoryItem.InventoryId,
			inventoryItem.RestaurantId,
			inventoryItem.IngredientId,
			inventoryItem.Stock,
			inventoryItem.Unit,
			inventoryItem.ExpirationDate,
		)

		inventoryItemList = append(inventoryItemList, inventoryItem)
	}

	return inventoryItemList, nil
}

func (i *InventoryRepository) GetById(ctx context.Context, inventoryId int) (*inventory.Inventory, error) {

	query := `
	SELECT * FROM inventories
	WHERE inventory_id = $1
	`

	inventoryItem := &inventory.Inventory{}
	err := i.db.QueryRowContext(ctx, query, inventoryId).Scan(
		inventoryItem.InventoryId,
		inventoryItem.RestaurantId,
		inventoryItem.IngredientId,
		inventoryItem.Stock,
		inventoryItem.Unit,
		inventoryItem.ExpirationDate,
	)

	if err != nil {
		return nil, err
	}

	return inventoryItem, nil
}

func (i *InventoryRepository) GetByRestaurant(ctx context.Context, restaurantId int) ([]*inventory.Inventory, error) {

	query := `
	SELECT * FROM inventories
	WHERE restaurant_id = $1
	`

	rows, err := i.db.QueryContext(ctx, query, restaurantId)

	if err != nil {
		return nil, err
	}

	inventoryItemList := []*inventory.Inventory{}

	for rows.Next() {
		inventoryItem := &inventory.Inventory{}

		err := rows.Scan(

			inventoryItem.InventoryId,
			inventoryItem.RestaurantId,
			inventoryItem.IngredientId,
			inventoryItem.Stock,
			inventoryItem.Unit,
			inventoryItem.ExpirationDate,
		)

		if err != nil {
			return nil, err
		}

		inventoryItemList = append(inventoryItemList, inventoryItem)
	}

	return inventoryItemList, nil
}

func (i *InventoryRepository) AddStock(ctx context.Context, inventoryId int, quantity float64, expitationDate time.Time) error {
	query := `
	UPDATE inventories SET stock = $2, expiration_date = $3
	WHERE inventori_id = $1
	`
	_, err := i.db.ExecContext(ctx, query, inventoryId, quantity, expitationDate)

	return err

}

func (i *InventoryRepository) RemoveStock(ctx context.Context, inventoryId int, quantity float64, unit measure.Measure) error {
	query := `
	UPDATE inventories SET stock = $2
	WHERE inventori_id = $1
	`
	_, err := i.db.ExecContext(ctx, query, inventoryId, quantity)

	return err
}

func (i *InventoryRepository) UpdateUnit(ctx context.Context, inventoryId int, unit measure.Measure) error {
	query := `
	UPDATE inventories SET unit = $2
	WHERE inventori_id = $1
	`
	_, err := i.db.ExecContext(ctx, query, inventoryId, unit)

	return err
}

func (i *InventoryRepository) UpdateExpitationDate(ctx context.Context, inventoryId int, expirationDate time.Time) error {
	query := `
	UPDATE inventories SET expiration_date = $2
	WHERE inventori_id = $1
	`
	_, err := i.db.ExecContext(ctx, query, inventoryId, expirationDate)

	return err
}

func (i *InventoryRepository) Delete(ctx context.Context, inventoryId int) error {
	query := `
	DELETE FROM inventories
	WHERE inventori_id = $1
	`
	_, err := i.db.ExecContext(ctx, query, inventoryId)

	return err
}
