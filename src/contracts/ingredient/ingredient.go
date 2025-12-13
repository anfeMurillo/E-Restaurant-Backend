package ingredient

import (
	"context"
	"e-restaurant/models/ingredient"
)

type Repository interface {
	Create(ctx context.Context, ingredient *ingredient.Ingredient) (*ingredient.Ingredient, error)

	UpdateId(ctx context.Context, ingredientId string, new string) error

	UpdateName(ctx context.Context, ingredientId string, new string) error

	UpdatePrice(ctx context.Context, ingredientId string, new float64) error

	Delete(ctx context.Context, ingredientId string) error
}
