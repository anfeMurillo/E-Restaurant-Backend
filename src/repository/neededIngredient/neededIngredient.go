package neededingredient

import (
	"context"
	"database/sql"
	"e-restaurant/models/enums/measure"
	model "e-restaurant/models/neededIngredient"
)

type NeededIngredientRepository struct {
	db *sql.DB
}

func NewNeededIngredientRepository(db *sql.DB) *NeededIngredientRepository {
	return &NeededIngredientRepository{db: db}
}

func (r *NeededIngredientRepository) Create(ctx context.Context, nedIngredient *model.NeededIngredient) (*model.NeededIngredient, error) {
	query := `INSERT INTO needed_ingredients (
	dish_id, 
	ingredient_id, 
	needed_quantity, 
	unit, 
	is_optional	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING dish_id, ingredient_id, needed_quantity, unit, is_optional`
	row := r.db.QueryRowContext(ctx, query,
		nedIngredient.DishId,
		nedIngredient.IngredientId,
		nedIngredient.NeededQuantity,
		nedIngredient.Unit,
		nedIngredient.IsOptional,
	)
	var ni model.NeededIngredient
	err := row.Scan(&ni.DishId, &ni.IngredientId, &ni.NeededQuantity, &ni.Unit, &ni.IsOptional)
	if err != nil {
		return nil, err
	}
	return &ni, nil
}

func (r *NeededIngredientRepository) UpdateQuantity(ctx context.Context, quantity float64, dishId int, ingredientId int) error {
	query := `UPDATE needed_ingredients SET needed_quantity = $1 WHERE dish_id = $2 AND ingredient_id = $3`
	_, err := r.db.ExecContext(ctx, query, quantity, dishId, ingredientId)
	return err
}

func (r *NeededIngredientRepository) UpdateUnit(ctx context.Context, unit measure.Measure, dishId int, ingredientId int) error {
	query := `UPDATE needed_ingredients SET unit = $1 WHERE dish_id = $2 AND ingredient_id = $3`
	_, err := r.db.ExecContext(ctx, query, unit, dishId, ingredientId)
	return err
}

func (r *NeededIngredientRepository) ToggleUnit(ctx context.Context, unit measure.Measure, dishId int, ingredientId int) error {

	query := `UPDATE needed_ingredients SET unit = $1 WHERE dish_id = $2 AND ingredient_id = $3`
	_, err := r.db.ExecContext(ctx, query, unit, dishId, ingredientId)
	return err
}

func (r *NeededIngredientRepository) Delete(ctx context.Context, dishId int, ingredientId int) error {
	query := `DELETE FROM needed_ingredients WHERE dish_id = $1 AND ingredient_id = $2`
	_, err := r.db.ExecContext(ctx, query, dishId, ingredientId)
	return err
}

func (r *NeededIngredientRepository) GetAll(ctx context.Context) ([]*model.NeededIngredient, error) {
	query := `SELECT dish_id, ingredient_id, needed_quantity, unit, is_optional FROM needed_ingredients`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.NeededIngredient
	for rows.Next() {
		var ni model.NeededIngredient
		if err := rows.Scan(&ni.DishId, &ni.IngredientId, &ni.NeededQuantity, &ni.Unit, &ni.IsOptional); err != nil {
			return nil, err
		}
		result = append(result, &ni)
	}
	return result, nil
}

func (r *NeededIngredientRepository) GetById(ctx context.Context, dishId int, ingredientId int) ([]*model.NeededIngredient, error) {
	query := `SELECT dish_id, ingredient_id, needed_quantity, unit, is_optional FROM needed_ingredients WHERE dish_id = $1 AND ingredient_id = $2`
	rows, err := r.db.QueryContext(ctx, query, dishId, ingredientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.NeededIngredient
	for rows.Next() {
		var ni model.NeededIngredient
		if err := rows.Scan(&ni.DishId, &ni.IngredientId, &ni.NeededQuantity, &ni.Unit, &ni.IsOptional); err != nil {
			return nil, err
		}
		result = append(result, &ni)
	}
	return result, nil
}

func (r *NeededIngredientRepository) GetByIngredient(ctx context.Context, ingredientId int) ([]*model.NeededIngredient, error) {
	query := `SELECT dish_id, ingredient_id, needed_quantity, unit, is_optional FROM needed_ingredients WHERE ingredient_id = $1`
	rows, err := r.db.QueryContext(ctx, query, ingredientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.NeededIngredient
	for rows.Next() {
		var ni model.NeededIngredient
		if err := rows.Scan(&ni.DishId, &ni.IngredientId, &ni.NeededQuantity, &ni.Unit, &ni.IsOptional); err != nil {
			return nil, err
		}
		result = append(result, &ni)
	}
	return result, nil
}

func (r *NeededIngredientRepository) GetByDish(ctx context.Context, dishId int) ([]*model.NeededIngredient, error) {
	query := `SELECT dish_id, ingredient_id, needed_quantity, unit, is_optional FROM needed_ingredients WHERE dish_id = $1`
	rows, err := r.db.QueryContext(ctx, query, dishId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.NeededIngredient
	for rows.Next() {
		var ni model.NeededIngredient
		if err := rows.Scan(&ni.DishId, &ni.IngredientId, &ni.NeededQuantity, &ni.Unit, &ni.IsOptional); err != nil {
			return nil, err
		}
		result = append(result, &ni)
	}
	return result, nil
}
