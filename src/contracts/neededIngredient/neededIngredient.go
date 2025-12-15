package neededingredient

import (
	"context"
	"e-restaurant/models/enums/measure"
	neededingredient "e-restaurant/models/neededIngredient"
)

type Repository interface {
	Create(ctx context.Context, nedIngredient *neededingredient.NeededIngredient) (*neededingredient.NeededIngredient, error)

	UpdateQuantity(ctx context.Context, quantity float64, dishId int, ingredientId int) error

	UpdateUnit(ctx context.Context, unit measure.Measure, dishId int, ingredientId int) error

	ToggleUnit(ctx context.Context, unit measure.Measure, dishId int, ingredientId int) error

	Delete(ctx context.Context, dishId int, ingredientId int) error

	GetAll(ctx context.Context) ([]*neededingredient.NeededIngredient, error)

	GetById(ctx context.Context, dishId int, ingredientId int) ([]*neededingredient.NeededIngredient, error)

	GetByIngredient(ctx context.Context, ingredientId int) ([]*neededingredient.NeededIngredient, error)

	GetByDish(ctx context.Context, dishId int) ([]*neededingredient.NeededIngredient, error)
}
