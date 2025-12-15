package ingredient

import (
	"context"
	"database/sql"
	"e-restaurant/models/ingredient"
)

type IngredientRepository struct {
	db *sql.DB
}

func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (i *IngredientRepository) Create(ctx context.Context, ingredient *ingredient.Ingredient) (*ingredient.Ingredient, error) {
	query := `
	INSERT INTO ingredients (ingredient_id,ingredient_name,price)
	VALUE ($1,$2,$3)
	RETURNING ingredient_id,ingredient_name,price
	`

	err := i.db.QueryRowContext(ctx, query, ingredient.IngredientId, ingredient.IngredientName, ingredient.Price).Scan(&ingredient.IngredientId, &ingredient.IngredientName, &ingredient.Price)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (i *IngredientRepository) GetById(ctx context.Context, ingredientID string) (*ingredient.Ingredient, error) {
	query := `
	SELECT * FROM ingredients
	WHERE ingredient_id = $1
	`
	ingredient := &ingredient.Ingredient{}
	err := i.db.QueryRowContext(ctx, query, ingredientID).Scan(&ingredient.IngredientId, &ingredient.IngredientName, &ingredient.Price)
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (i *IngredientRepository) GetByDish(ctx context.Context, dishID int) ([]*ingredient.Ingredient, error) {
	query := `
	SELECT

	i.ingredient_id,
	i.ingredient_name,
	i.price

	FROM ingredients i
	INNER JOIN needed_ingredients USING (ingredient_id)
	WHERE dish_id = $1
	`

	ingredientList := []*ingredient.Ingredient{}

	rows, err := i.db.QueryContext(ctx, query, dishID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		ingredient := &ingredient.Ingredient{}

		err := rows.Scan(&ingredient.IngredientId, &ingredient.IngredientName, &ingredient.Price)
		if err != nil {
			return nil, err
		}
		ingredientList = append(ingredientList, ingredient)
	}

	return ingredientList, nil

}

func (i *IngredientRepository) GetAll(ctx context.Context) ([]*ingredient.Ingredient, error) {
	query := `
	SELECT * FROM ingredients
	`
	rows, err := i.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows.Close()

	ingredientList := []*ingredient.Ingredient{}

	for rows.Next() {
		ingredient := &ingredient.Ingredient{}

		err := rows.Scan(&ingredient.IngredientId, &ingredient.IngredientName, &ingredient.Price)
		if err != nil {
			return nil, err
		}

		ingredientList = append(ingredientList, ingredient)
	}

	return ingredientList, nil
}

func (i *IngredientRepository) UpdateName(ctx context.Context, ingredientId int, new string) error {
	query := `
	UPDATE ingredients SET ingredient_name = $1
	WHERE ingredient_id = $2
	`

	_, err := i.db.ExecContext(ctx, query, new, ingredientId)

	return err

}

func (i *IngredientRepository) UpdatePrice(ctx context.Context, ingredientId int, new float64) error {
	query := `
	UPDATE ingredients SET ingredient_price = $1
	WHERE ingredient_id = $2
	`

	_, err := i.db.ExecContext(ctx, query, new, ingredientId)

	return err

}

func (i *IngredientRepository) Delete(ctx context.Context, ingredientId int) error {
	query := `
	DELETE FROM ingredients
	WHERE ingredient_id = $1
	`

	_, err := i.db.ExecContext(ctx, query, ingredientId)

	return err

}
