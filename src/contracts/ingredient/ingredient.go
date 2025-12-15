package ingredient

import (
	"context"
	"e-restaurant/models/ingredient"
)

type Repository interface {
	Create(ctx context.Context, ingredient *ingredient.Ingredient) (*ingredient.Ingredient, error)

	GetById(ctx context.Context, ingredientId int) (*ingredient.Ingredient, error)

	GetByDish(ctx context.Context, dishID int) ([]*ingredient.Ingredient, error)

	GetAll(ctx context.Context) ([]*ingredient.Ingredient, error)

	UpdateName(ctx context.Context, ingredientId int, new string) error

	UpdatePrice(ctx context.Context, ingredientId int, new float64) error

	Delete(ctx context.Context, ingredientId int) error
}
