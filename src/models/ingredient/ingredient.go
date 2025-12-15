package ingredient

type Ingredient struct {
	IngredientId   int     `db:"ingredient_id" json:"ingredient_id"`
	IngredientName string  `db:"ingredient_name" json:"ingredient_name"`
	Price          float64 `db:"price" json:"price"`
}
